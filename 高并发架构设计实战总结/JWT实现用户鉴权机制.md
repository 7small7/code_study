[TOC]

## JWT是什么

JSON Web Token（缩写 JWT）是目前最流行的<font color='red'>跨域</font>认证解决方案。它是有三部分组成,示例如下，具体的讲解如下(jwt是不会有空行的，下面只是为了显示，便使用了换行看着比较方便)。

```shell
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjMfQ.
SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```
它是由一个"."号隔开、三部分组成。

第一部分是header信息，

```php
{

  "alg": "HS256",// 加密的算法

  "typ": "JWT"// 加密的方式，填写JWT

}

```

第二部分是Payload，有固定的六个部分和自定义数据组成,自定义数据看自己的情况需要来定义，是可以省去的。

```php
'iss' => 'https://www.qqdeveloper.com',// 签发人
'exp' => time() + 86400,// 过期时间(这里的有效期时间为1天)
'sub' => '主题内容',// 主题
'aud' => '受众内容',// 受众
'nbf' => $time,// 生效时间
'iat' => $time,// 签发时间
'jti' => 123,// 编号

```

第三部分是Signature(是对前两部分加密得来的)。由于前两部分是公开透明的数据，因此防止数据的篡改和泄露，我们需要加密处理。首先，需要指定一个密钥（secret）。这个密钥只有服务器才知道，不能泄露给用户。然后，使用 Header 里面指定的签名算法（默认是 HMAC SHA256），按照下面的公式产生签名。

```php

第一部分的加密方式(
base64UrlEncode(header) + "." +
base64UrlEncode(payload),
secret)
```

最终生成的就是上面很长的一段字符串了。

## 为什么会使用JWT

这就需要从我们传统的认证模式来说了，传统的认证模式是基于session和cookie来实现用户的认证和鉴权。具体的流程模式如下图。

![](http://qiniucloud.qqdeveloper.com/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-09-09%2021.00.14.png)

<center>(图一)Session与Cookie认证与鉴权</center>

1.客户端向服务端发送一个http请求。

2.服务端在收到客户端的请求时，生成一个唯一的sessionid，<font color='red'>这里需要将该生成的session存储在服务端</font>,这个sessionid存储具体的session内容，默认的是文件存储，当然我们可以修改具体的存储方式，例如数据库存储。

3.客户端在接受到这个sessionid时，存在cookie里面，每次请求时携带该sessionid。

4.服务端在接收到客户端的请求之后，根据客户端发送的sessionid来进行认证与授权。

这里也推荐一下自己之前分享的一篇有关session于cookie的知识点。[session与cookie详解](https://www.qqdeveloper.com/2019/08/18/PHP-session%E4%B8%8Ecookie%E8%AF%A6%E8%A7%A3/)

![](http://qiniucloud.qqdeveloper.com/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-09-09%2021.08.19.png)

<center>(图二)传统的token授权</center>

1.客户端向服务端发送一个http请求。

2.服务端在收到客户端的请求之后，生成一个唯一token，<font color='red'>这里需要将该生成的token存储在服务端</font>，至于怎么存，可以和上面session与cookie的方式一致。也可以存在缓存数据库中，如redis，memcached。

3.服务端将该token返回给客户端，客户端存在本地，可以存请求头header中，也可以存在cookie中，同时也可以存在localstorage中。

4.向服务端发送请求时，携带该token，服务端进行认证或者授权。

![](http://qiniucloud.qqdeveloper.com/%E5%B1%8F%E5%B9%95%E5%BF%AB%E7%85%A7%202019-09-09%2021.13.30.png)

<center>(图三)JWT认证模式</center>

1.客户端向服务端发送一个http请求。

2.服务端根据jwt的生成规则，生成一个token，并返回给客户端，<font color='red'>这里服务端是不需要存储的</font>。

3.客户端在接受到该token时，存在客户端。

4.客户端向服务端发送请求时，服务端对请求的token进行解析，如果发现解析出来的数据和生成的数据是一致的代表是一个合法的token，则进行相应的操作。

## 基于session和cookie的认证和鉴权模式有什么好与不好的地方呢？

通过上面几张图，我们也大致可以看得出来，基于session都是需要服务端存储的，而JWT是不需要服务端来存储的。针对以上几点，总结如下：

一、缺点

1.容易遇到跨域问题。不同域名下是无法通过session直接来做到认证和鉴权的。

2.分布式部署的系统，需要使用共享session机制

3.容易出现csrf问题。

二、优点

1.方便灵活，服务器端直接创建一个sessionid，下发给客户端，客户端请求携带sessionid即可。

2.session存储在服务端，更加安全。

3.便于服务端清除session，让用户重新授权一次。

## JWT与session有什么区别呢？

JWT是基于客户端存储的一种认证方式，然而session是基于服务端存储的一种认证方式。JWT虽然不用服务端存储了，也可以避免跨域、csrf等情况。但也存在如下几个不太好的地方。

1.无法清除认证token。由于JWT生成的token都是存储在客户端的，不能有服务端去主动清除，只有直到失效时间到了才能清除。除非服务端的逻辑做了改变。

2.存储在客户端，相对服务端，安全性更低一些。当JWT生成的token被破解，我们不便于清除该token。

## 如何使用JWT

这里推荐使用GitHub上面人家封装好的包，这里我使用的是firebase/php-jwt，在项目中直接使用即可安装成功。

```php 
composer require firebase/php-jwt
```

接下来创建一个控制器,我这里使用的ThinkPHP5.1的框架
```php

use think\Controller;
use Firebase\JWT\JWT;

class Test extends Controller

{
    private $key = 'jwtKey';
    // 生成JWT
    public function createJwt()
    {
        $time  = time();
        $key   = $this->key;
        $token = [
            'iss' => 'https://www.qqdeveloper.com',// 签发人
            'exp' => $time + 86400,// 过期时间(这里的有效期时间为1天)
            'sub' => '主题内容',// 主题
            'aud' => '受众内容',// 受众
            'nbf' => $time,// 生效时间
            'iat' => $time,// 签发时间
            'jti' => 123,// 编号
            // 额外自定义的数据
            'data' => [
                'userName' => '编程浪子走四方'
            ]];
        // 调用生成加密方法('Payloadn内容','加密的键',['加密算法'],['加密的可以'],['JWT的header头'])
        $jwt = JWT::encode($token, $key);
        return json(['data' => $jwt]);
    }
    // 解析JWT
    public function analysisJwt()
    {
        try {
            $key = $this->key;
            $jwt = 'eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJodHRwOlwvXC9leGFtcGxlLm9yZyIsImV4cCI6MTU2ODA5NjE4MCwic3ViIjoiXHU0ZTNiXHU5ODk4XHU1MTg1XHU1YmI5IiwiYXVkIjoiXHU1M2Q3XHU0ZjE3XHU1MTg1XHU1YmI5IiwibmJmIjoxNTY4MDA5NzgwLCJpYXQiOjE1NjgwMDk3ODAsImp0aSI6MTIzLCJkYXRhIjp7InVzZXJOYW1lIjoiXHU3ZjE2XHU3YTBiXHU2ZDZhXHU1YjUwXHU4ZDcwXHU1NmRiXHU2NWI5In19.kHb_9Np0zjE25YE9czUEGvmFPYtqMJT9tuZzJTuMZl0';
            // 调用解密方法('JWT内容','解密的键,和加密时的加密键一直','加密算法')
            $decoded = JWT::decode($jwt, $key, array('HS256'));
            return json(['message' => $decoded]);
        } catch (\Exception $exception) {
            return json(['message' => $exception->getMessage()]);
        }
    }
}
```

通过访问第一个方法，可以生成下图一段字符串
![](http://qiniucloud.qqdeveloper.com/create.jpg)
我们将上图中的字符串复制到第二图中的$jwt变量,访问第二个方法即可解析出具体的数据。
![](http://qiniucloud.qqdeveloper.com/%E8%A7%A3%E5%AF%86.jpg)