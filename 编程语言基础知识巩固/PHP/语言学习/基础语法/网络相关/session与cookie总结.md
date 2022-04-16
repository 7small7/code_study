[TOC]

## PHP session 与 cookie 区别

### session 与 cookie 是什么?

session 与 cookie 属于一种会话控制技术.常用在身份识别，登录验证，数据传输等.举个例子，就像我们去超市买东西结账的时候，我们要拿出我们的会员卡才会获取优惠.这时候，我们怎么识别这个会员卡真实有效的呢?当我们将会员号给到收银员，收银员根据我们提供的会员号，输入到系统中，系统根据这个会员号去查询，如果查询到了就证明这个会员号是真实存在的.这里的会员号就好比 cookie 与 session.会员系统就好比服务器端，收银员就好比客。

### 为什么会用到 session 与 cookie 呢?

根据上述的例子，我们知道 session 与 cookie 是可以干什么的了，那为什么必须用这个来实现呢？这里就有必要了解一下 http 应用传输协议的特点了。由于 http 协议是无状态的，即浏览器去请求了一个网页，这时候就是一个 http 请求，当服务端接收到请求之后，返回客户端需要的数据，在这过程中浏览器与服务器是建立了一个连接的。但是当服务端返回数据，客户端收到数据之后，他们的这种连接关系就断开了。下次浏览器再去发送请求的时候，又是重新建立一个连接，这两个链接没有任何关系。试想一下，当我们登录一个商场系统的时候，进入首页做了登录操作，但是我们下单或者加入购物车的时候，还需要登录，每访问一个页面就要登录，是不是很繁琐同时也是很不科学的，万一我们加入购物车的商品，我们点击下单了，下单页面要登录而且还无法正确的反馈出你下单时的那些商。

## Http 特点

1.http 协议支持客户端/服务端模式，也是一种请求/响应模式的协议。

2.无连接。所谓的无连接就是服务器收到了客户端的请求之后，响应完成并收到客户端的应答之后，即断开连接。限制每次的连接只处理一次请求。从而节省传输时间。 

3.无状态。http 协议对事务的处理没有记忆能力。也就意味着如果需要前面的信息，只能重传，这无形之中增加数据的传输量。这种方式某种方面上讲解放了服务器，但是却不利于客户端与服务器的连接。为了弥补这种不足，产生了两项记录 http 状态的技术，一个叫做 Cookie,一个叫做 Session，后面我们再细讲它们。 

4.简单快捷：所谓的简单快捷是指客户端向服务器请求服务时，一般来说只需要传输请求方法和路径，就能进行访问 

5.灵活：这里主要指的是客户端可以通过 http 协议传输任意类型的数据。比如传输.jpg 文件、.ppt 文件等等，只需要设定 content-type 就可以进行传输。

## Cookie

### cookie 的基本概念

cookie 是远程浏览器存储数据以此追踪用户和识别用户的的机制，从实现来说，cookie 是存储在客户端上的一个数据片段。

### cookie 的运行原理与存储机制

#### 运行原理
 1.客户端向服务端发起一个 http 请求。
 
 2.服务端设置一个创建 cookie 的指令，响应给客户端。
 
 3.客户端收到服务端响应的指令，根据指令在客户端创建一个 cookie。
 
 4.挡下一次请求时，客户端携带这个 cookie 向服务端发送请求。
 
#### 存储机制
总的来说，cookie 在客户端存储的形式有三种,不同的浏览器的存储机制不同，存的 cookie 也不同。

1.文件存储.浏览器会针对不同的域，在磁盘的对应目录创建一个单独的文件，来存储该域下面的 cookie 值。

2.内存存储.当浏览器关闭时，该 cookie 随之消失.根据下面的创建语法，当我们未设置过期时间时则会出现这种情况。

3.flash 存储.这种存储方式是永久存储在磁盘中，即使通过浏览器删除一些数据都是无法删除该方式存储的 cookie，如果需要删除，可能通过磁盘的方。

### cookie 的设置

```php
Bool setcookie(string $name[, string $values, $expire=0[,string $path[,string $domain[, bool $secure = false[, bool $httpOnly = false]]]]] );
```
1.$name:cookie存储的名称，必填选项。

2.$values:cookie存储的值。这里需要注意的是，当把该值设置为false时，客户端会尝试删除这个cookie值，因此在要将值这是为true或者false的时候，我们用另外的值来代替，例如true用1代替，false用0来代替。

3.$expire:cookie的过期时间，秒为单位，当该值被设置时，定时删除；当该值没有设置时，该值是永久有效的.该值设置为小于当前时间时，会出发浏览器的删除机制，会自动删除cookie。

4.$path:cookie有效的目录，默认的目录是"/"，即表示当前的整个个域名都生效.

5.$domain:cookie的作用域名，默认的是当前域名有效，如果需要设置直接填写生效的域名即可.需要注意的是IE浏览器有长度限制，当只有大于5的时候才会生效。

6.$secure:cookie的加密处理，当设置为true的时候，需要使用HTTPS协议，才会生效。

7.$httpOnly:决定cookie是否只使用http协议，当设置为1或者true，其他非http协议是无法操作cookie的。例如我们未设置的时候，我们JavaScript是可以对cookie进行设置的.这样一定程度上保证了安全性.这种情况需考虑浏览器是否支持该配置项。该方式的原理是，服务端在发送创建cookie的时候，会告知浏览器添加一个特殊的参数，来防止JavaScript进行读取。

. 设置 cookie 的函数还有 setrawcookie()函数，只不过该函数不会对值 进行 urlencode 序列号
.<font color="red">有时候，我们可能遇到这种情况，我们在这个页面设置了 cookie，但是去刷新页面获取 cookie，按理说是会获取到 cookie 的，但实际情况是无法获取到，这是由于 cookie 运行机制导致，PHP 创建了 cookie 这个指令，告诉浏览器，你需要执行这个指令了，这时候浏览器才会去执行这个指令，因此是无法获取到 cookie 的.</font>
. 在设置 cookie 之前，不能有任何输出.

```
// 实现方式一
setcookie($cookie,"hello,world!", 3600);
// 实现方式二
header("header("Set-Cookie: testcookie=中文; path=/; domain=.sunphp.org; expires=".gmstrftime("%A, %d-%b-%Y %H:%M:%S GMT",time()+9600));");
```
> 两则的作用是一样的，setcookie是PHP内置函数，是对http协议的操作封装。

### cookie 的获取

```php
$_COOKIE['$cookeName'];
```

### cookie 的应用场景

1.用户身份识别

2.数据传输

3.登录控制(是否登录、SSO单点登录)

### cookie 跨域设置

我们都知道，在前端开发中时常会遇到 ajax 跨域问题，我们解决的方式有很多种，可以参考这篇文章[传送门 1](http://www.qqdeveloper.com/web/detail/2)，[传送门 2](http://www.qqdeveloper.com/web/detail/8)，cookie 跨域我们可以参考 p3p 传输协议[传送门](https://www.cnblogs.com/lmy01/p/6369159.html)

### cookie存储的形式

cookie存在客户端的主要形式分为如下3种:

1.硬盘:对于cookie存储，是放在本地的文件中，该文件通过加密处理。

2.内存:当生成cookie时，给cookie的expire(过期时间)设置为空，客户端在创建cookie时，存储在浏览器的内存中，当浏览器关闭则内存释放，故cookie则也被释放掉了。

3.flash:这类的cookie是存在flash上，也能存储在硬盘上。这类的cookie不受浏览器的控制。

### cookie 使用的注意事项

1.数量限制，客户端对每一个 domian 下的 cookie 是有数量限制的，不是创建任意数量就行。

2.安全性，根据上面的创建语法，我们可以得知，当我们未设置$httpOnly值得时候，非http协议是可以操作cookie的值的，例如JavaScript通过cookie($cookieName).而且一些抓包工具也是可以抓取到 cookie 的,还有就是 cookie 存储在客户端的文件中，如果获取到这个 cookie，也是可以对 cookie 做一些操作的.为了防止别人可以拷贝 cookie 文件，进行恶意操作，可以对 cookie 进行加密处理。

3.数据传输:当 cookie 数量很多，数据很大的时候，其实对于带宽是有消耗的.比较 http 传输都需要带宽，当 http 传输的数据量大了，带了的带宽消耗就大。

## Session

### 运行原理与存储机制

. 运行原理 
1.客户端向服务端发起请求，建立通信 

2.服务端根据设置的 session 创建指令，在服务端创建一个编号为 sessionid 的文件，里面的值就是 session 具体的值(组成部分 变量名 | 类型 :长度:值). 

3.服务端将创建好的 sessionid 编号响应给客户端，客户则将该编号存在 cookie 中(一般我们在浏览器存储的调试栏中会发现 cookie 中有一个 PHPSESSID 的键，这就是 sessionid，当然这个名称，我可以通过设置服务端是可以改变的).

.当下一次请求时，客户端将这个 sessionid 携带在请求中，发送给服务端，服务端根据这个 sessionid 来做一些业务判断.

.存储机制 
1.存储方式.session 默认是文件存储的.我们可以通过 php.ini 的配置来设置存储驱动[传送门](http://php.net/manual/zh/session.configuration.php#ini.session.serialize-handler)

2.生命周期.当我们未设置 session 的生命周期时，当浏览器关闭之后存储在客户端的 phpsessid 自动消失，因为它是存在内存，下次建立连接的时候会重新创建一个 phpsessid.之前的 session，PHP 会自动的根据垃圾回收机制自动删除.这里我们可以根据[session_set_cookie_params(\$expire)函数](http://php.net/manual/zh/function.session-set-cookie-params.php)来设置一个生命周期;

### session 的设置

```php
session_start();
$_SESSION = $values;
```

. session_start()设置之前，不能有任何输出

### session 的获取

```php
$_SESSION['values'];
```

### session 的删除

```php
// 只是单纯的给重新赋了一个空的值
$_SESSION['values'] = '';
// 该函数是清空所有的session，慎用!
session_destroy();
// 连values这个session键都会删除
unset($_SESSION['values']);
```

### session 的使用场景

. 用户身份识别
. 数据传输
. 登录控制(是否登录、单点登录)

### session 的注意事项

.安全性,sessionid 是按照一定的算法生成，要保证 session 的值唯一性和随机性.
.客户端禁用 cookie，根据上面 session 的运行原理可以得出，session 的存储于传送还是依赖于客户端，因此当客户端禁用 cookie 时，客户端是无法保存 PHPSESSID 的，这时候可以通过 url 重写或者表单来实现 session 的传输.
.存储优化,按照上面的 session 创建，所有的 session 都会创建在一个目录下面，同时有的无效 session 在垃圾回收机制时间内还不会删除，当一台服务器配置的站点较多时，这时候会生成很多的 session 文件，导致我们读取速度变慢，我们可以设置 session 的存储目录级别,[save_path 函数](http://php.net/manual/zh/session.configuration.php#ini.session.save-path).一般大型的项目(如分布式的项目),可以使用其他的存储方式，如数据存储，内存存储.

## session 与 cookie 的区别

1.session存储在服务端，cookie存储在客户端。
2.cookie的创建指令由服务端生成，服务端只是发送cookie创建的指令，客户端在接收到服务端的指令时，在再客户端进行创建。
3.session 的 sessionid 需要客户端存储。

### cookie 与 session 的几个误区

1.客户端禁止 cookie，session 无法使用？

a.使用url重写或者表单提交可以实现。

2.session 和 cookie 的安全性比较，session 存在客户端安全更高?

a.由于cookie是存在客户端的，相对来说安全性是要低一些，不过在创建的时候可以设置$httpOnly值。
b.由于cookie与session是相互关联的，获取到cookie一定程度上获取到了session，同样可以操作session。

3.cookie 与 session 是不是在浏览器关闭的时候会消失?

a.这需要查看存储机制了。cookie可以存文件，内存，flash.存内存当然浏览器关闭则消失了；
b.session由于垃圾回收机制，当在垃圾回收机制内是不会删除的，除非你代码中显示的做了删除操作。

4.cookie 是存储在客户端中，如何增加其安全性?

a.我们可以在设置cookie的时候，增加一些特殊参数，如客户端信息ip、浏览器信息等。

5.当 cookie 存在客户端的文件中，是不是每个浏览器获取到这个文件都可以进行操作?

a.要看浏览器之间对cookie的管理机制是不是一样。

6.如果浏览器禁用了cookie，session是否还能继续使用?

默认情况下，cookie如果被禁用了，则无法在客户端创建sessionid了。但是session也可以通过url或者表单等方式进行传输。