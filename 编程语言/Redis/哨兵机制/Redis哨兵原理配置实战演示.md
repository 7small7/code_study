## 文章简介

本文将通过理论+实践的方式从头到尾总结Redis中的哨兵机制。文章内容**从主从复制的弊端**、**如何解决弊端**、**什么是哨兵**、**哨兵监控的图形结构**、**哨兵监控的原理**、**如何配置哨兵**、**哨兵与主从复制的关系**等方面来演示。

文中相关资料下载地址:[链接:](https://pan.baidu.com/s/1cDV9eXuUwuA0QFELDMkIHQ)  密码: mv86

## 主从复制弊端
![1618556721550-Snipaste_2021-04-16_15-04-45](http://qiniucloudtest.qqdeveloper.com/mweb/1618556721550-Snipaste_2021-04-16_15-04-45.png)
上面的图形结构，大致的可以理解为Redis的主从复制拓扑图。

1. 其中1个主节点负责应用系统的写入数据，另外的4个从节点负责应用系统的读数据。

2. 同时4个从节点向其中的1个一个主节点发起复制请求操作。

在Redis服务运行正常的情况下，该拓扑结结构不会出现什么问题。试想一下这样的一个场景。如果主节点服务发生了异常，不能正常处理服务(如写入数据、主从复制操作)。**这时候，Redis服务能正常响应应用系统的读操作，但是没法进行写操作。** 出现该情况就会严重影响到系统的业务数据。那该如何解决呢？

可以大致想到下面的几种情况来解决。

1. 当主节点发生异常情况时，手动的从部分从节点中选择一个节点作为主节点。然后改变其他从节点的主从复制关系。

2. 我们也可以写一套自动处理该情况的服务，避免依赖于人为的操作。

上面的方案在一定程度上是能帮助我们解决问题。但是过多的人为干预。例如第1点，我们需要考虑人工处理的实时性和正确性。第2点，自动化处理是能够很好的解决第1点中的问题，但是自动处理存在如何选择新主节点的问题，因此这也是一个不好的地方。

通过上面大致的分析，我们不难得出Redis的哨兵机制就是针对种种问题出现的。

## 什么是哨兵

可以把Redis的哨兵理解为一种**Redis分布式架构**。 该架构中主要存在两种角色，一种是哨兵，另外一种是数据节点(主从复制节点)。
![1618558716540-Snipaste_2021-04-16_15-38-08](http://qiniucloudtest.qqdeveloper.com/mweb/1618558716540-Snipaste_2021-04-16_15-38-08.png)
哨兵主要负责的任务是：

1. 每一个哨兵都会监控数据节点以及其他的哨兵节点。

2. 当其中的一个哨兵监控到节点不可达是，会给对应的节点做下线标识。如果下线的节点为主节点。这时候会通知其他的哨兵节点。

3. 哨兵节点通过“协商”推举出从节点中的某一个节点为主节点。

4. 接着将其他的从节点断开与旧主节点的复制关系，将推举出来的新主节点作为从节点的主节点。

5. 将切换的结果通知给应用系统。
![1618558995604-Snipaste_2021-04-16_15-43-04](http://qiniucloudtest.qqdeveloper.com/mweb/1618558995604-Snipaste_2021-04-16_15-43-04.png)
## 配置哨兵

在演示环境中，配置了三台数据节点(1主2从)，三台哨兵节点。演示中用到的Redis为6.0.8版本。

| 角色 | IP | 端口号 | 
| :---: | :---: | :---: |
| (数据节点)master | 127.0.0.1 | 8002 |
| (数据节点)slave | 127.0.0.1 | 8003 |
| (数据节点)slave | 127.0.0.1 | 8004 |
| 哨兵节点 | 127.0.0.1 | 8005 |
| 哨兵节点 | 127.0.0.1 | 8006 |
| 哨兵节点 | 127.0.0.1 | 8007 |

1. (数据节点)master配置。

```redis
# 服务配置
daemonize yes

# 端口号
port 8002

# 数据目录
dir "/Users/kert/config/redis/8002"

# 日志文件名称
logfile "8002.log"

# 设置密码
bind 0.0.0.0
# requirepass 8002

# 多线程
# 1.开启线程数。
io-threads 2
# 2.开启读线程。
io-threads-do-reads yes

# 持久化存储(RDB)
# 1.每多少秒至少有多少个key发生变化，则执行save命令。
save 10 1
save 20 1
save 30 1
# 2.当bgsave命令发生错误时，停止写入操作。
stop-writes-on-bgsave-error yes
# 3.是否开启rbd文件压缩
rdbcompression yes
```
2. (数据节点)slave配置。

```redis
# 服务配置
daemonize yes

port 8004

dir "/Users/kert/config/redis/8004"

logfile "8004.log"

# 多线程
# 1.开启线程数。
io-threads 2
# 2.开启读线程。
io-threads-do-reads yes

# 持久化存储(RDB)
# 1.每多少秒至少有多少个key发生变化，则执行save命令。
save 10 1
save 20 1
save 30 1
# 2.当bgsave命令发生错误时，停止写入操作。
stop-writes-on-bgsave-error yes
# 3.是否开启rbd文件压缩
rdbcompression yes

# 配置主节点信息
replicaof 127.0.0.1 8002
```

3. 哨兵节点配置。

```redis
# 端口号
port 8006
# 运行模式
daemonize yes
# 数据目录
dir "/Users/kert/config/redis/sentinel/8006"
# 日志文件
logfile "8006.log"
# 监听数据节点
sentinel monitor mymaster 127.0.0.1 8002 2(判定主节点下线状态的票数)
# 设置主节点连接权限信息
sentinel auth-pass mymaster 8002
# 判断数据节点和sentinel节点多少毫秒数内没有响应ping，则处理为下线状态
sentinel down-after-milliseconds mymaster 30000
# 主节点下线后，从节点向新的主节点发起复制的个数限制(指的一次同时允许几个从节点)。
sentinel parallel-syncs mymaster 1
# 故障转移超时时间
sentinel failover-timeout mymaster 180000
```
> 所有的哨兵节点直接将port、dir和logfile修改为对应的具体哨兵信息即可。

接着启动对应的服务Redis服务。
```shell
// 启动master节点
kert@kertdeMacBook-Pro-2  ~/config/redis/8002  redis-server ./redis.conf

// 启动slave节点
kert@kertdeMacBook-Pro-2  ~/config/redis/8003  redis-server ./redis.conf
kert@kertdeMacBook-Pro-2  ~/config/redis/8004  redis-server ./redis.conf

// 启动哨兵节点
kert@kertdeMacBook-Pro-2  ~/config/redis/sentinel  redis-sentinel 8007.conf
kert@kertdeMacBook-Pro-2  ~/config/redis/sentinel  redis-sentinel 8006.conf
kert@kertdeMacBook-Pro-2  ~/config/redis/sentinel  redis-sentinel 8005.conf
``` 
> 哨兵启动，需要用到Redis安装完之后自带的 redis-sentinel命令。

查看Redis服务运行状态。
```shell
 kert@kertdeMacBook-Pro-2  ~/config/redis/sentinel  ps -ef | grep redis
  501 99742     1   0  3:53PM ??         0:00.47 redis-server 0.0.0.0:8002
  501 99776     1   0  3:53PM ??         0:00.36 redis-server 0.0.0.0:8003
  501 99799     1   0  3:53PM ??         0:00.10 redis-server *:8004
  501 99849     1   0  3:53PM ??         0:00.06 redis-sentinel *:8007 [sentinel]
  501 99858     1   0  3:53PM ??         0:00.04 redis-sentinel *:8006 [sentinel]
  501 99867     1   0  3:53PM ??         0:00.03 redis-sentinel *:8005 [sentinel]
```
> 看到上面的结果，则表示我们的Redis服务已经正常启动。

## 演示故障切换

我们先打开三个终端，分配时master节点和两个slave节点。检测是否能够正常进行主从复制。
![1618560150817-Snipaste_2021-04-16_16-01-40](http://qiniucloudtest.qqdeveloper.com/mweb/1618560150817-Snipaste_2021-04-16_16-01-40.png)
我们在主节点任意写入一些数据，然后在从节点进行查询数据。为了方便，后面将master称作1号终端，两个slave分配叫做2号和3号终端。
1. 我们在1号终端写入数据。

```shell
127.0.0.1:8002> set name tony
OK
127.0.0.1:8002> set age 1
OK
127.0.0.1:8002> set socre 1
OK
127.0.0.1:8002>
```
2. 接着在2号和3号终端下面执行如下的查询操作。

```shell
127.0.0.1:8003> get name
"tony"
127.0.0.1:8003> get age
"1"
127.0.0.1:8003> get socre
"1"
```
> 事实证明我们的主从复制是成功的，接下来我们就停掉master节点的服务。

我们实现查看一下哨兵节点的一个状态信息。
1. 查看哨兵端口为8005的节点。

```shell
kert@kertdeMacBook-Pro-2  ~  redis-cli -p 8005 info
# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=mymaster,status=ok,address=127.0.0.1:8002,slaves=2,sentinels=3
```
2. 查看哨兵端口为8006的节点。

```shell
kert@kertdeMacBook-Pro-2  ~  redis-cli -p 8006 info
# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=mymaster,status=ok,address=127.0.0.1:8002,slaves=2,sentinels=3
```
3. 查看哨兵端口为8007的节点。

```shell
kert@kertdeMacBook-Pro-2  ~  redis-cli -p 8007 info
# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=mymaster,status=ok,address=127.0.0.1:8002,slaves=2,sentinels=3
```
> 通过上面的几个状态信息，我们可以看到哨兵检测的主节点信息，主节点下面有几个从节点，同时哨兵节点有几个。

我们杀掉master的进程。可以看到1号端口自动断开了连接。
![1618560942345-Snipaste_2021-04-16_16-15-12](http://qiniucloudtest.qqdeveloper.com/mweb/1618560942345-Snipaste_2021-04-16_16-15-12.png)
接着我们通过哨兵机制查看一下数据节点状态信息。
```shell
kert@kertdeMacBook-Pro-2  ~  redis-cli -p 8005 info
# Sentinel
sentinel_masters:1
sentinel_tilt:0
sentinel_running_scripts:0
sentinel_scripts_queue_length:0
sentinel_simulate_failure_flags:0
master0:name=mymaster,status=ok,address=127.0.0.1:8004,slaves=2,sentinels=3
``` 
> 通过上面的查询结果，我们可以看到address的值编程了8004端口了，其他的信息没有发生改变，说明哨兵已经完成切换工作。

接下来我们在新的主节点执行操作命令，查看在从节点是否能够完成主从复制。
1. 在3号端口(新的master)执行一个del命令。

```shell
127.0.0.1:8004> del age
(integer) 1
127.0.0.1:8004> keys *
1) "name"
2) "socre"
```
2. 在2号端口执行读命令。

```shell
127.0.0.1:8003> keys *
1) "socre"
2) "name"
```
> 此时可以发现我们的主从复制也是正常的。
3. 启动旧的master，并执行读命令。

```shell
 kert@kertdeMacBook-Pro-2  ~/config/redis/8002  redis-server ./redis.conf
 kert@kertdeMacBook-Pro-2  ~/config/redis/8002  redis-cli -p 8002
127.0.0.1:8002> keys *
1) "name"
2) "socre"
```
> 此时你也会发现，原来的master节点变成了slave节点，并且能够正常复制新master节点的数据。

## 配置文件对比

在我们启动了哨兵模式之后，我们的哨兵配置文件和数据节点配置文件的内容都会自动的生成一个特定的内容。
1. 数据节点(master距离)。

变化前
```shell
# 服务配置
daemonize yes

# 端口号
port 8002

# 数据目录
dir "/Users/kert/config/redis/8002"

# 日志文件名称
logfile "8002.log"

# 设置密码
bind 0.0.0.0
# requirepass 8002

# 多线程
# 1.开启线程数。
io-threads 2
# 2.开启读线程。
io-threads-do-reads yes
# 持久化存储(RDB)
# 1.每多少秒至少有多少个key发生变化，则执行save命令。
save 10 1
save 20 1
save 30 1
# 2.当bgsave命令发生错误时，停止写入操作。
stop-writes-on-bgsave-error yes
# 3.是否开启rbd文件压缩
rdbcompression yes
```
变化后
```shell
# 服务配置
daemonize yes

# 端口号
port 8002

# 数据目录
dir "/Users/kert/config/redis/8002"

# 日志文件名称
logfile "8002.log"

# 设置密码
bind 0.0.0.0
# requirepass 8002

# 多线程
# 1.开启线程数。
io-threads 2
# 2.开启读线程。
io-threads-do-reads yes

# 持久化存储(RDB)
# 1.每多少秒至少有多少个key发生变化，则执行save命令。
save 10 1
save 20 1
save 30 1
# 2.当bgsave命令发生错误时，停止写入操作。
stop-writes-on-bgsave-error yes
# 3.是否开启rbd文件压缩
rdbcompression yes
# Generated by CONFIG REWRITE
pidfile "/var/run/redis.pid"
user default on nopass ~* +@all

replicaof 127.0.0.1 8004
```

2. 哨兵节点

变化前
```shell
# 端口号
port 8006
# 运行模式
daemonize yes
# 数据目录
dir "/Users/kert/config/redis/sentinel/8006"
# 日志文件
logfile "8006.log"
# 监听数据节点
sentinel monitor mymaster 127.0.0.1 8002 2(判定主节点下线状态的票数)
# 设置主节点连接权限信息
sentinel auth-pass mymaster 8002
# 判断数据节点和sentinel节点多少毫秒数内没有响应ping，则处理为下线状态
sentinel down-after-milliseconds mymaster 30000
# 主节点下线后，从节点向新的主节点发起复制的个数限制(指的一次同时允许几个从节点)。
sentinel parallel-syncs mymaster 1
# 故障转移超时时间
sentinel failover-timeout mymaster 180000
```

变化后
```shell
# 端口号
port 8005
# 运行模式
daemonize yes
# 数据目录
dir "/Users/kert/config/redis/sentinel/8005"
# 日志文件
logfile "8005.log"
# 监听数据节点
sentinel myid 5724fd60af87e728e6f8f03ded693960c983e156
# 判断数据节点和sentinel节点多少毫秒数内没有响应ping，则处理为下线状态
sentinel deny-scripts-reconfig yes
# 主节点下线后，从节点向新的主节点发起复制的个数限制(指的一次同时允许几个从节点)。
sentinel monitor mymaster 127.0.0.1 8004 2
# 故障转移超时时间
sentinel config-epoch mymaster 3
# Generated by CONFIG REWRITE
protected-mode no
user default on nopass ~* +@all
sentinel leader-epoch mymaster 3
sentinel known-replica mymaster 127.0.0.1 8002
sentinel known-replica mymaster 127.0.0.1 8003
sentinel known-sentinel mymaster 127.0.0.1 8006 8fbd2cce642c881f752775afee9b3591e0d90dc6
sentinel known-sentinel mymaster 127.0.0.1 8007 69530c74791e5f32db1c2a006c826a6463bc6496
sentinel current-epoch 3
pidfile "/var/run/redis.pid"
```

## 实战代码

这里我们使用PHP原生类操作Redis哨兵，首先我们创建一个Redis操作类，类中代码如下：
```php
class OperationRedis
{
 private $redis;

    private $requestParams;

    private $redisHandler;

    /**
     * 本机IP地址
     * @var string
     */
    private $redisHost = '192.168.2.102';

    public function __construct()
    {
        $this->requestParams = Request::instance()->all();
        $this->redisHandler  = new \Redis();
        $this->redis         = $this->redisHandler->connect($this->redisHost, 8005);
    }

    /**
     * 获取Redis哨兵信息
     * @author kert
     */
    public function getRedisNode()
    {
        /** @var array $masterLists 通过sentinel节点获取所有主节点 */
        $masterLists = $this->redisHandler->rawCommand('SENTINEL', 'masters');
        dump('master列表配置信息', $masterLists);
        foreach ($masterLists as $value) {
            /** @var array $masterInfo 主节点信息 */
            $masterInfo = $this->redisHandler->rawCommand('SENTINEL', 'master', $value[1]);
            dump('master节点信息', $masterInfo);

            // 向主节点插入数据
            $insertNumber = $this->insertInfoByMaster((string)$this->redisHost, (int)$value[5]);
            dump('Redis队列数量', $insertNumber);

            /** @var array $slaveLists 主节点下面的从节点信息 */
            $slaveLists = $this->redisHandler->rawCommand('SENTINEL', 'slaves', $value[1]);
            dump('master下的slave节点信息', $slaveLists);

            foreach ($slaveLists as $v) {
                $this->getInfoBySlave((string)$this->redisHost, (int)$v[5]);
            }
        }
    }

    /**
     * 向主节点插入数据
     * @param string $masterIp
     * @param int $port
     * @return int
     * @author kert
     */
    private function insertInfoByMaster(string $masterIp, int $port): int
    {
        $masterRedis = new \Redis();
        $masterRedis->connect($masterIp, $port);

        return $masterRedis->lPush('sentinel', time());
    }

    /**
     * 向所有的从节点获取数据
     * @param string $slaveIp
     * @param int $port
     * @author kert
     */
    private function getInfoBySlave(string $slaveIp, int $port)
    {
        $slaveRedis = new \Redis();
        $slaveRedis->connect($slaveIp, $port);
        $array = $slaveRedis->lRange('sentinel', 0, 1000);
        echo "从节点{$slaveIp},端口号{$port}获取到的对应数据为:" . PHP_EOL;
        dump($array);
    }
}
```

通过访问该代码，得到如下结果：
![1618648165832-Snipaste_2021-04-17_16-28-59](http://qiniucloudtest.qqdeveloper.com/mweb/1618648165832-Snipaste_2021-04-17_16-28-59.png)
> 改代码在实际的生产中，肯定使用时不对的，这里只是为了演示代码如何操作哨兵。

其中的操作逻辑大致如下图：
![1618650744497-Snipaste_2021-04-17_17-11-46](http://qiniucloudtest.qqdeveloper.com/mweb/1618650744497-Snipaste_2021-04-17_17-11-46.png)
## Laravel框架配置哨兵

Laravel框架自带Redis操作类。我们只需要简单配置即可。找到config/database.php文件。设置如下配置信息即可：
```php
'redis' => [
'client' => env('REDIS_CLIENT', 'predis'),
'options' => [
    'cluster' => env('REDIS_CLUSTER', 'predis'),
    'prefix'  => env('REDIS_PREFIX', Str::slug(env('APP_NAME', 'laravel'), '_') . '_database_'),
],

'default' => [
    'tcp://192.168.2.102:8005',
    'tcp://192.168.2.102:8006',
    'tcp://192.168.2.102:8007',    //这3个都是sentinel节点的地址
    'options' => [
        'replication' => 'sentinel',
        'service'     => env('REDIS_SENTINEL_SERVICE', 'mymaster'),    //sentinel
        'parameters'  => [
            'password' => env('REDIS_PASSWORD', null),    //redis的密码,没有时写null
            'database' => 0,
        ],
    ],           
    'database' => env('REDIS_DB', 0),
  ],
],
```
接下来就可以直接操作Redis数据了。
```php
public function laravelRedis()
{
    var_dump(Redis::connection()->set(time(), time()));
}
```
```php
// output
object(Predis\Response\Status)#237 (1) {
  ["payload":"Predis\Response\Status":private]=>
  string(2) "OK"
}
```