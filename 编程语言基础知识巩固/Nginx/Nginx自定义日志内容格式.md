### 题外话

该问题还是在公司项目终于到的，边记录下来分享一下。公司一个项目采用的前后端分离，前端在请求后端接口是，后端日志无法获取到请求的参数，进过查看代码日志也未发现请求参数，边考虑通过 nginx 的日志文件来处理。

#### 配置前的记录信息

![1](http://qiniucloud.qqdeveloper.com/mweb/1.jpeg)

#### 设置后的信息

![2](http://qiniucloud.qqdeveloper.com/mweb/2.jpeg)

> 两者对比

通过上面的对比，我们很容易看得出，后者多记录了一些请求参数。例如 content,type 等字段。这些都是 nginx 配置的 log 日志格式。第一张图是 nginx 默认的配置格式，我们完全可以自定义一些格式信息。配置一些我们需要的格式信息，边与我们排除错误，方便查看请求信息。

> 配置步骤

第一步、找到 nginx 的默认配置文件 nginx.conf 文件，在 http 配置段里面配置如下信息

```shell
# log_format是配置日志格式的前缀，固定
# main 是配置名，就好比变成中创建一个变量一样，方便在后面的配置文件中直接引用
log_format  main escape=json '$remote_addr - $remote_user [$time_local] "$request" '
'$status $body_bytes_sent "$http_refer
er" '
'"$http_user_agent" "$http_x_forwarded_for"'
'$upstream_addr $upstream_response_time $request_time $request_body ';
```

第二步、在域名配置文件中引用该配置信息
例如，我配置了一个 www.test.com.conf 的域名配置文件。在配置下面该项配置时，直接在后面加一个 main 即可。

```shell
access_log  /data/logs/nginx/www.test.com.log main;
```

第三步、访问 www.test.com 的域名，查看/data/logs/nginx/www.test.com.log 该文件，即可看到类似与图二的信息。

> nginx 的日志常见配置项

```shell
$args                    #请求中的参数值
$query_string            #同 $args
$arg_NAME                #GET请求中NAME的值
$is_args                 #如果请求中有参数，值为"?"，否则为空字符串
$uri                     #请求中的当前URI(不带请求参数，参数位于$args)，可以不同于浏览器传递的$request_uri的值，它可以通过内部重定向，或者使用index指令进行修改，$uri不包含主机名，如"/foo/bar.html"。
$document_uri            #同 $uri
$document_root           #当前请求的文档根目录或别名
$host                    #优先级：HTTP请求行的主机名>"HOST"请求头字段>符合请求的服务器名.请求中的主机头字段，如果请求中的主机头不可用，则为服务器处理请求的服务器名称
$hostname                #主机名
$https                   #如果开启了SSL安全模式，值为"on"，否则为空字符串。
$binary_remote_addr      #客户端地址的二进制形式，固定长度为4个字节
$body_bytes_sent         #传输给客户端的字节数，响应头不计算在内；这个变量和Apache的mod_log_config模块中的"%B"参数保持兼容
$bytes_sent              #传输给客户端的字节数
$connection              #TCP连接的序列号
$connection_requests     #TCP连接当前的请求数量
$content_length          #"Content-Length" 请求头字段
$content_type            #"Content-Type" 请求头字段
$cookie_name             #cookie名称
$limit_rate              #用于设置响应的速度限制
$msec                    #当前的Unix时间戳
$nginx_version           #nginx版本
$pid                     #工作进程的PID
$pipe                    #如果请求来自管道通信，值为"p"，否则为"."
$proxy_protocol_addr     #获取代理访问服务器的客户端地址，如果是直接访问，该值为空字符串
$realpath_root           #当前请求的文档根目录或别名的真实路径，会将所有符号连接转换为真实路径
$remote_addr             #客户端地址
$remote_port             #客户端端口
$remote_user             #用于HTTP基础认证服务的用户名
$request                 #代表客户端的请求地址
$request_body            #客户端的请求主体：此变量可在location中使用，将请求主体通过proxy_pass，fastcgi_pass，uwsgi_pass和scgi_pass传递给下一级的代理服务器
$request_body_file       #将客户端请求主体保存在临时文件中。文件处理结束后，此文件需删除。如果需要之一开启此功能，需要设置client_body_in_file_only。如果将次文件传 递给后端的代理服务器，需要禁用request body，即设置proxy_pass_request_body off，fastcgi_pass_request_body off，uwsgi_pass_request_body off，or scgi_pass_request_body off
$request_completion      #如果请求成功，值为"OK"，如果请求未完成或者请求不是一个范围请求的最后一部分，则为空
$request_filename        #当前连接请求的文件路径，由root或alias指令与URI请求生成
$request_length          #请求的长度 (包括请求的地址，http请求头和请求主体)
$request_method          #HTTP请求方法，通常为"GET"或"POST"
$request_time            #处理客户端请求使用的时间,单位为秒，精度毫秒； 从读入客户端的第一个字节开始，直到把最后一个字符发送给客户端后进行日志写入为止。
$request_uri             #这个变量等于包含一些客户端请求参数的原始URI，它无法修改，请查看$uri更改或重写URI，不包含主机名，例如："/cnphp/test.php?arg=freemouse"
$scheme                  #请求使用的Web协议，"http" 或 "https"
$server_addr             #服务器端地址，需要注意的是：为了避免访问linux系统内核，应将ip地址提前设置在配置文件中
$server_name             #服务器名
$server_port             #服务器端口
$server_protocol         #服务器的HTTP版本，通常为 "HTTP/1.0" 或 "HTTP/1.1"
$status                  #HTTP响应代码
$time_iso8601            #服务器时间的ISO 8610格式
$time_local              #服务器时间（LOG Format 格式）
$cookie_NAME             #客户端请求Header头中的cookie变量，前缀"$cookie_"加上cookie名称的变量，该变量的值即为cookie名称的值
$http_NAME               #匹配任意请求头字段；变量名中的后半部分NAME可以替换成任意请求头字段，如在配置文件中需要获取http请求头："Accept-Language"，$http_accept_language即可
$http_cookie
$http_host               #请求地址，即浏览器中你输入的地址（IP或域名）
$http_referer            #url跳转来源,用来记录从那个页面链接访问过来的
$http_user_agent         #用户终端浏览器等信息
$http_x_forwarded_for
$sent_http_NAME          #可以设置任意http响应头字段；变量名中的后半部分NAME可以替换成任意响应头字段，如需要设置响应头Content-length，$sent_http_content_length即可
$sent_http_cache_control
$sent_http_connection
$sent_http_content_type
$sent_http_keep_alive
$sent_http_last_modified
$sent_http_location
$sent_http_transfer_encoding
```

> nginx 代理处理

在实际的项目中，可能我们的请求可能会遇到一些代理服务器，这时候我们是无法获取到真实的客服端信息的。这就需要我们配置一些请求头信息。大致的原理就是，代理服务器和我们的服务器之间有一个默认的协议，代理服务器在向我们服务器发送请求的时候，实际会把真实的 ip 地址传递过来，只需要我们配置一些参数即可获取。（代理服务器大致的工作流程就是，客户端在请求我们的服务器时，先是去访问代理服务器，代理服务器在把这个请求转发给我们真实的服务器，真实的服务器在返回给代理服务器，会后客户端在从代理服务器获取到相关的数据信息）。下面实例择抄至网络。

> 正确设置 nginx 中 remote_addr 和 x_forwarded_for 参数

#### 什么是 remote_addr?

remote_addr 代表客户端的 IP，但它的值不是由客户端提供的，而是服务端根据客户端的 ip 指定的，当你的浏览器访问某个网站时，假设中间没有任何代理，那么网站的 web 服务器（Nginx，Apache 等）就会把 remote_addr 设为你的机器 IP，如果你用了某个代理，那么你的浏览器会先访问这个代理，然后再由这个代理转发到网站，这样 web 服务器就会把 remote_addr 设为这台代理机器的 IP

#### 什么是 x_forwarded_for?

正如上面所述，当你使用了代理时，web 服务器就不知道你的真实 IP 了，为了避免这个情况，代理服务器通常会增加一个叫做 x_forwarded_for 的头信息，把连接它的客户端 IP（即你的上网机器 IP）加到这个头信息里，这样就能保证网站的 web 服务器能获取到真实 IP

#### 使用 HAProxy 做反向代理时

通常网站为了支撑更大的访问量，会增加很多 web 服务器，并在这些服务器前面增加一个反向代理（如 HAProxy），它可以把负载均匀的分布到这些机器上。你的浏览器访问的首先是这台反向代理，它再把你的请求转发到后面的 web 服务器，这就使得 web 服务器会把 remote_addr 设为这台反向代理的 IP，为了能让你的程序获取到真实的客户端 IP，你需要给 HAProxy 增加以下配置

```shell
option forwardfor
```

它的作用就像上面说的，增加一个 x_forwarded_for 的头信息，把客户端的 ip 添加进去，否则的话经测试为空值
如上面的日志格式所示：$http_x_forwarded_for 是客户端真实的IP地址，$remote_addr 是前端 Haproxy 的 IP 地址
或者：
当 Nginx 处在 HAProxy 后面时，就会把 remote_addr 设为 HAProxy 的 IP，这个值其实是毫无意义的，你可以通过 nginx 的 realip 模块，让它使用 x_forwarded_for 里的值。使用这个模块需要重新编译 Nginx，增加--with-http_realip_module 参数

```shell
./configure  --user=www --group=www --prefix=/usr/local/nginx --with-http_stub_status_module --with-http_ssl_module \       --with-http_realip_module --http-log-path=/data/logs/nginx/access.log --error-log-path=/data/logs/nginx/error.log
```

```shell
 set_real_ip_from 10.1.10.0/24;
real_ip_header X-Forwarded-For;
```

上面的两行配置就是把从 10.1.10 这一网段过来的请求全部使用 X-Forwarded-For 里的头信息作为 remote_addr，这样此时 remote_addr 就是客户端真实的 IP 地址。

> X-Forwarded-For 和 X-Real-IP 获取客户端的 ip 的区别：

一般来说，X-Forwarded-For 是用于记录代理信息的，每经过一级代理(匿名代理除外)，代理服务器都会把这次请求的来源 IP 追加在 X-Forwarded-For 中 来自 4.4.4.4 的一个请求，header 包含这样一行 X-Forwarded-For: 1.1.1.1, 2.2.2.2, 3.3.3.3 代表 请求由 1.1.1.1 发出，经过三层代理，第一层是 2.2.2.2，第二层是 3.3.3.3，而本次请求的来源 IP 4.4.4.4 是第三层代理 而 X-Real-IP，一般只记录真实发出请求的客户端 IP，上面的例子，如果配置了 X-Read-IP，将会是 X-Real-IP: 1.1.1.1 所以 ，如果只有一层代理，这两个头的值就是一样的。
[更多文章](http://nginx.org/en/docs/varindex.html)