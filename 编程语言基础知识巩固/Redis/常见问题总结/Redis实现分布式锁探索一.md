## 实战说明

最近在一个项目营销活动中，一位同事用到了Redis来实现商品的库存管理。在压测的过程中，发现存在超卖的情况。这里总结一篇如何正确使用Redis来解决秒杀场景下，超卖的情况。

## 演示步骤

这里不会直接给大家说明，该怎么去实现安全、高效的分布式锁。而是通过循序渐进的方式，通过不同的方式实现锁，并发现每一种锁的缺点以及针对该类型的锁进行如何优化，最终达到实现一个高效、安全的分布锁。

### 第一种场景

该场景是利用Redis来存储商品数量。先获取库存，针对库存判断，如果库存大于0，则减少1，再更新Redis库存数据。大致示意图如下：
![1635168007654-Snipaste_2021-10-25_21-19-53](http://qiniucloudtest.qqdeveloper.com/mweb/1635168007654-Snipaste_2021-10-25_21-19-53.png)
1. 当第一个请求来之后，去判断Redis的库存数量。接着给商品库存减少一，然后去更新库存数量。

2. 当在第一个请求处理库存逻辑之间，第二个请求来了，同样的逻辑，去读Redis库存，判断库存数量接着执行减少库存操作。此时他们操作的商品其实就是同一个商品。

3. 然后这样的逻辑，在秒杀这样大量请求来，就很容易实际商品售卖的数量远远大于商品库存数量。

```php
public function demo1(ResponseInterface $response)
{
    $application = ApplicationContext::getContainer();
    $redisClient = $application->get(Redis::class);
    /** @var int $goodsStock 商品当前库存*/
    $goodsStock = $redisClient->get($this->goodsKey);

    if ($goodsStock > 0) {
        $redisClient->decr($this->goodsKey);
        // TODO 执行额外业务逻辑
        return $response->json(['msg' => '秒杀成功'])->withStatus(200);
    }
    return $response->json(['msg' => '秒杀失败，商品库存不足。'])->withStatus(500);
}
```

问题分析：

1. 该方式使用Redis来管理商品库存，减少对MySQL的压力。

2. 假设此时库存只有1，第一请求在判断库存为大于0，减少库存的过程中。如果存在第二个请求来读取到了数据，发现商品库存是大于0的。两者都会执行秒杀的逻辑，然而库存只有一个，就遇到了超卖的情况。

3. 此时，我们试想一下，如果我们只能让一个请求处理库存，其他的请求只有等待直到上一个请求结束才能去进行获取商品库存，是不是就能实现超卖呢？这就是下面几种场景提到的锁机制实现。

### 第二种场景

使用文件锁，第一请求来了之后，打开文件锁。处理完毕业务之后，释放当前的文件锁，接着处理下一个请求，依次循环。保证当前的所有请求，只有一个请求在处理库存。请求处理完毕之后，则释放锁。
![1635168534444-Snipaste_2021-10-25_21-28-40](http://qiniucloudtest.qqdeveloper.com/mweb/1635168534444-Snipaste_2021-10-25_21-28-40.png)
1. 使用文件锁，来一个请求给一个文件加锁。此时另外的请求就会被阻塞，直到上一个请求成功释放锁文件，下一个请求才会执行。

2. 所有的请求就犹如一个队列一样，前一个先入队列后一个后入队列，一次按照FIFO的顺序进行。
```php
public function demo3(ResponseInterface $response)
{
    $fp = fopen("/tmp/lock.txt", "r+");

    try {
        if (flock($fp, LOCK_EX)) {  // 进行排它型锁定
            $application = ApplicationContext::getContainer();
            $redisClient = $application->get(Redis::class);
            /** @var int $goodsStock 商品当前库存*/
            $goodsStock = $redisClient->get($this->goodsKey);
            if ($goodsStock > 0) {
                $redisClient->decr($this->goodsKey);
                // TODO 处理额外的业务逻辑
                $result = true; // 业务逻辑处理最终结果
                flock($fp, LOCK_UN);    // 释放锁定
                fclose($fp);
                if ($result) {
                    return $response->json(['msg' => '秒杀成功'])->withStatus(200);
                }
                return $response->json(['msg' => '秒杀失败'])->withStatus(200);
            } else {
                flock($fp, LOCK_UN);    // 释放锁定
                fclose($fp);
                return $response->json(['msg' => '库存不足，秒杀失败。'])->withStatus(500);
            }
        } else {
            fclose($fp);
            return $response->json(['msg' => '活动过于火爆，抢购的人过多，请稍后重试。'])->withStatus(500);
        }
    } catch (\Exception $exception) {
        fclose($fp);
        return $response->json(['msg' => '系统异常'])->withStatus(500);
    } finally {
        fclose($fp);
    }
}
```

问题分析：

1. 如果使用文件锁，开启一个锁和释放一个锁都是比较耗时的。在秒杀的业务场景下，大量请求过来，很容易出现大部分用户一直处于请求等待的过程中。

2. 当开启一个文件锁时，都是针对当前服务器。如果我们的项目属于分布式部署，上述的加锁也只能针对当前的服务器进行加锁，而不是针对的请求进行加锁。如下图：容易时刻存在多个多服务器多个锁。
![1635168727744-Snipaste_2021-10-25_21-31-57](http://qiniucloudtest.qqdeveloper.com/mweb/1635168727744-Snipaste_2021-10-25_21-31-57.png)
## 第三种场景

该方案是通过先Redis存储商品库存，来一个请求就针对上面的库存减少1，Redis如果返回的库存小于0则表示当前的秒杀失败。主要是利用到了Redis的单线程写。保证每次对Redis的写都只有一个线程在执行。
![1635168974964-Snipaste_2021-10-25_21-36-05](http://qiniucloudtest.qqdeveloper.com/mweb/1635168974964-Snipaste_2021-10-25_21-36-05.png)
```php
public function demo2(ResponseInterface $response)
{
    $application = ApplicationContext::getContainer();
    $redisClient = $application->get(Redis::class);
    /** @var int $goodsStock Redis减少1后的库存数据 */
    $goodsStock = $redisClient->decr($this->goodsKey);
    if ($goodsStock > 0) {
        // TODO 执行额外业务逻辑
        $result = true;// 业务处理的结果
        if ($result) {
            return $response->json(['msg' => '秒杀成功'])->withStatus(200);
        } else {
            $redisClient->incr($this->goodsKey);// 将减少的库存进行增加1
            return $response->json(['msg' => '秒杀失败'])->withStatus(500);
        }
    }
    return $response->json(['msg' => '秒杀失败，商品库存不足。'])->withStatus(500);
}
```

问题分析：

1. 该方案虽然利用了Redis的单线程模型特点，可以避免超卖的清空。当库存为0时，来一个秒杀请求，就会将库存减少1，最终Redis的缓存数据肯定会为小于0。

2. 该方案存在用户秒杀数量与实际秒杀商品数量不一致。如上述代码，在业务处理结果为FALSE的时候，给Redis增加1，如果增加1的过程中发生异常，没有增加成功就会导致商品数量不一致的情况。

## 第四种场景

通过上面的三种情况分析，我们可以得出文件锁的情况是最好的一种方案。但是文件锁不能解决分布式部署的清空。这时候我们可以利用Redis的setnx，expire来实现分布锁。setnx命令先设置一个锁，expire给锁加一个超时时间。

```php
public function demo4(ResponseInterface $response)
{
    $application = ApplicationContext::getContainer();
    $redisClient = $application->get(Redis::class);

    if ($redisClient->setnx($this->goodsKey, 1)) {
        // 假设该执行下面的操作时服务器宕机
        $redisClient->expire($this->goodsKey, 10);
        // TODO 处理业务逻辑
        $result = true;// 处理业务逻辑的结果
        // 删除锁
        $redisClient->del($this->goodsKey);
        if ($result) {
            return $response->json(['msg' => '秒杀成功。'])->withStatus(200);
        }
        return $response->json(['msg' => '秒杀失败。'])->withStatus(500);
    }
    return $response->json(['msg' => '系统异常，请重试。'])->withStatus(500);
}
```

问题分析：

1. 通过上面的实例代码，我们会感觉到该这种方法似乎没有什么问题。加一个锁，在释放锁。但细想一下，setnx命令在添加锁之后，给锁设置过期时间(expire)时发生了异常导致没有正常给锁加上过期时间。是不是这一把锁就一直在呢？

2. 所以上述的情况来实现Redis分布式锁，是不满足原子性的。

### 第五种场景

在第四种场景中，利用到了Redis实现分布式锁。但是该分布式锁不是原子性的，好在Redis提供该两个命令的结合版，可以实现原子性。就是set(key, value, ['nx', 'ex' => '过期时间'])。

```php
public function demo5(ResponseInterface $response)
{
    $application = ApplicationContext::getContainer();
    $redisClient = $application->get(Redis::class);

    if ($redisClient->set($this->goodsKey, 1, ['nx', 'ex' => 10])) {
        try {
            // TODO 处理秒杀业务
            $result = true;// 处理业务逻辑的结果
            $redisClient->del($this->goodsKey);
            if ($result) {
                return $response->json(['msg' => '秒杀成功。'])->withStatus(200);
            } else {
                return $response->json(['msg' => '秒杀失败。'])->withStatus(200);
            }
        } catch (\Exception $exception) {
            $redisClient->del($this->goodsKey);
        } finally {
            $redisClient->del($this->goodsKey);
        }
    }
    return $response->json(['msg' => '系统异常，请重试。'])->withStatus(500);
}
```

问题分析：

1. 通过一步一步的推进，可能你会觉得第五种场景，Redis来实现分布式应该是天衣无缝了吧。我们仔细去观察打TODO的地方，也是处理业务逻辑的地方。要是业务逻辑超过缓存设置的10秒会怎么样？

2. 如果逻辑处理超过10秒，此时第二个秒杀请求就能正常处理自身的业务请求。恰好，第一个请求的业务逻辑执行完毕，要删除Redis锁了，就会把第二个的请求的Redis锁给删除。第三个请求就会正常执行，按照此逻辑是不是Redis的锁一样是一个无效的锁呢？

3. 此情况就会导致当前的请求在删除Redis锁时，删除的不是自身的锁。如果我们在删除锁时，做一个验证，只能删除自身的锁，看看此方案是否行的通？接下来，我们看看第六种情况。

### 第六种场景

该方案针对上面第五种情况，在删除时添加了一个请求的唯一标识判断。也就是说只有删除自身添加锁时的标识。

```php
public function demo6(ResponseInterface $response)
{
    $application = ApplicationContext::getContainer();
    $redisClient = $application->get(Redis::class);
    /** @var string $client 当前请求的唯一标识*/
    $client = md5((string)mt_rand(100000, 100000000000000000).uniqid());
    if ($redisClient->set($this->goodsKey, $client, ['nx', 'ex' => 10])) {
        try {
            // TODO 处理秒杀业务逻辑
            $result = true;// 处理业务逻辑的结果
            $redisClient->del($this->goodsKey);
            if ($result) {
                return $response->json(['msg' => '秒杀成功'])->withStatus(200);
            }
            return $response->json(['msg' => '秒杀失败'])->withStatus(500);
        } catch (\Exception $exception) {
            if ($redisClient->get($this->goodsKey) == $client) {
                // 此处存在时间差
                $redisClient->del($this->goodsKey);
            }
        } finally {
            if ($redisClient->get($this->goodsKey) == $client) {
                // 此处存在时间差
                $redisClient->del($this->goodsKey);
            }
        }
    }
    return $response->json(['msg' => '请稍后重试'])->withStatus(500);
}
```

## 问题分析

1. 通过上面的情况分析下来，貌似一点问题都没有了。然而，仔细的你可以看看我添加注释的地方"此处存在时间差"。如果Redis在读取到缓存时，并且判断请求的唯一标识是一致的，在执行del删除锁时，发生了一个阻塞、网络波动等情况。在该锁过期之后，才去执行到del命令，此时删除的锁还是当前请求的锁吗？

2. 此时去删除锁，肯定就不是当前请求的锁。而是下一个请求的锁。这种情况，是否也会存在锁无效的情况呢？

## 问题总结

通过上面的几种实例代码演示，发现很大问题是在给Redis释放锁的时候，因为不属于一个原子性操作。结合第六种情况，如果我们能够保证释放锁是一个原子性，添加锁也是一个原子性，这样是不是就能正确保证我们的分布锁没有问题了？

1. 添加锁时，实现原子性操作，我们用Redis原生的命令就可以了。

2. 释放锁时，只删除自身添加的锁，我们在第六种场景中已经得到解决。

3. 接下来，就只需要考虑释放锁的时候，能够实现原子性操作。由于Redis原生没有这样的命令，我们就需要借助lua操作，来实现原子性。

## 具体实现

通过打开官网，可以看到官网提供分布式锁实现的几种客户端，直接使用即可。[官网地址](https://redis.io/topics/distlock)，这里我使用的客户端是[rtckit/reactphp-redlock](https://github.com/ronnylt/redlock-php)
。具体安装方式，直接按照文档操作即可。这里简单的说明一下两种方式的调用。

### 第一种方式

```php
 public function demo7()
{
    /** @var Factory $factory 初始化一个Redis实例*/
    $factory = new \Clue\React\Redis\Factory();
    $client  = $factory->createLazyClient('127.0.0.1');

    /** @var Custodian $custodian 初始化一个锁监听器*/
    $custodian = new \RTCKit\React\Redlock\Custodian($client);
    $custodian->acquire('MyResource', 60, 'r4nd0m_token')
        ->then(function (?Lock $lock) use ($custodian) {
            if (is_null($lock)) {
                // 获取锁失败
            } else {
                // 添加一个10s生命周期的锁
                // TODO 处理业务逻辑
                // 释放锁
                $custodian->release($lock);
            }
        });
}
```
该方式的大致逻辑，和我们在第六种方案中是差不多的，都是使用Redis的`set + nx` 命令实现原子性加锁，然后给当前加的锁设置一个随机的字符串，用来处理释放当前锁时，不能去释放他人的锁。做大的差别就是在使用`release`释放锁时，该方法去调用了一个lua脚本，来删除锁。保证锁的释放是一个原子性的。下面是释放锁的大致截图。
```php
// lua脚本
public const RELEASE_SCRIPT = <<<EOD
if redis.call("get", KEYS[1]) == ARGV[1] then
    return redis.call("del", KEYS[1])
else
    return 0
end
EOD;

public function release(Lock $lock): PromiseInterface
{
    /** @psalm-suppress InvalidScalarArgument */
    return $this->client->eval(self::RELEASE_SCRIPT, 1, $lock->getResource(), $lock->getToken())
        ->then(function (?string $reply): bool {
            return $reply === '1';
        });
}
```

### 第二种方式

第二种方式和第一种差距不大，无非是增加了一个`自旋锁`。一直去获取锁，如果没有获取到则放弃当前的请求。

```php
public function demo8()
{
    /** @var Factory $factory 初始化一个Redis实例*/
    $factory = new \Clue\React\Redis\Factory();
    $client  = $factory->createLazyClient('127.0.0.1');

    /** @var Custodian $custodian 初始化一个锁监听器*/
    $custodian = new \RTCKit\React\Redlock\Custodian($client);
    $custodian->spin(100, 0.5, 'HotResource', 10, 'r4nd0m_token')
        ->then(function (?Lock $lock) use ($custodian) : void {
            if (is_null($lock)) {
                // 将进行100次的场次，每一次间隔0.5秒去获取锁，如果没有获取到锁。则放弃加锁请求。
            } else {
                // 添加一个10s生命周期的锁
                // TODO 处理业务逻辑
                // 释放锁
                $custodian->release($lock);
            }
        });
}
```

## 自旋锁

spinlock又称自旋锁，是实现保护共享资源而提出一种锁机制。自旋锁与互斥锁比较类似，都是为了解决对某项资源的互斥使用。

无论是互斥锁，还是自旋锁，在任何时刻，最多只能有一个保持者，只能有一个执行单元获得锁。但是两者在调度机制上略有不同。对于互斥锁，如果资源已经被占用，资源申请者只能进入睡眠状态。但是自旋锁不会引起调用者睡眠，如果自旋锁已经被别的执行单元保持，调用者就一直循环在那里看是否该自旋锁的保持者已经释放了锁，"自旋"一词就是因此而得名。

## 总结

其实通过上面的几种方案，细心的你，可能还会发现很多问题。

1. 本身并发可以是多一个多线程的处理方式，我们这里添加锁之后，是不是并行处理变成串行处理了。降低了秒杀所谓的高性能。

2. 在Redis主从复制、集群等部署架构方案中，上面的方案还能行得通吗？

3. 很多人都在说zookeeper更适合拿来用分布式锁场景，那zookeeper比Redis耗在哪些地方呢？

带着种种疑问，我们在下一篇文章再见。喜欢的，感兴趣的，欢迎你关注我的文章。文章中存在不足的地方，也欢迎指正。