[TOC]

## 什么是CGI

讲Fastcgi之前需要先讲CGI，CGI是为了保证web server传递过来的数据是标准格式的，它是一个协议。每种动态语言（ PHP,Python 等）的代码文件需要通过对应的解析器才能被服务器识别，而 CGI 协议就是用来使解释器与服务器可以互相通信。PHP 文件在服务器上的解析需要用到 PHP 解释器，再加上对应的 CGI 协议，从而使服务器可以解析到 PHP 文件。

![](http://qiniucloud.qqdeveloper.com/iShot%20%20%20%20%202019-12-04%20AM12.11.26.png)
1.用户通过客户端(浏览器)输入一个网址，例如www.baidu.com。

2.浏览器经过一些列的处理(这里省略其中的流程)，请求到对应服务器的。

3.服务器网卡根据监听到的端口，将请求发送给对应的软件服务。

4.web服务器(Nginx/Apache等web软件)接收请求后，通过fast-cgi或者cgi协议，将请求数据转发给php-fpm进程管理器。php-fpm将任务下发给下面空闲的work进程，此时work进程中的php解释器开始处理文件。

5.php解释器处理好，通过fast-cgi或者cgi协议，再将转换后的数据返给web服务器，web服务器再响应给客户端。

## 什么是FastCG

Fastcgi是CGI的更高级的一种方式，是用来提高CGI程序性能的。<font color='red'>由于 CGI 的机制是每处理一个请求需要 fork 一个 CGI 进程，请求结束再kill掉这个进程</font>  ，在实际应用上比较浪费资源，于是就出现了CGI 的改良版本 FastCGI，<font color='red'>FastCGI 在请求处理完后，不会 kill 掉进程，而是继续处理多个请求</font>，这样就大大提高了效率。
![](http://qiniucloud.qqdeveloper.com/fastcgi.png)

## 什么是php-fpm

PHP-FPM是PHP实现的FastCGI Process Manager(FastCGI进程管理器), 用于替换PHP FastCGI的大部分附加功能，适用于高负载网站，并提供了进程管理的功能，可以有效控制内存和进程，平滑重载PHP配置。

1.平滑停止/启动的高级进程管理功

2.慢日志记录脚本

3.动态/静态子进程产生

4.基于php.ini的配置文件

进程包含 master 进程和 worker 进程两种；master 进程只有一个，负责监听端口，接收来自服务器的请求，而 worker 进程则一般有多个（[具体数量根据实际需要进行配置](http://www.qqdeveloper.com/docs/php/php-1ch9t4fejf6a9)），每个进程内部都会嵌入一个 PHP 解释器，是代码真正执行的地方。
在没有php-fpm之前，每当我们修改了php.ini的配置信息，都会面临着下面的几个问题:
1.需要重启php-cgi程序,才能使配置文件生效,同时php-cgi不支持平滑重启。
2.kill掉php-cgi程序时，必须重新启动，否则PHP就不能正常工作。
![](http://qiniucloud.qqdeveloper.com/iShot%20%20%20%20%202019-12-04%20AM12.22.46.png)
因此就可以把php-fpm理解为,是一个实现了Fastcgi协议的程序,用来管理Fastcgi启动的进程的,即能够调度php-cgi进程的程序。现已在PHP内核中就集成了PHP-FPM，使用--enalbe-fpm这个编译参数即可。另外，修改了php.ini配置文件后，没办法平滑重启，需要重启php-fpm才可。此时新fork的worker会用新的配置，已经存在的worker继续处理完手上的活。

### Web服务器与程序解析器运行流程（Nginx与php-fpm通信机制(通信流程)）

web server（如nginx）只是内容的分发者。比如，如果请求/index.html，那么web server会去文件系统中找到这个文件，发送给浏览器，这里分发的是静态资源。

如果现在请求的是/index.php，根据配置文件，nginx知道这个不是静态文件，需要去找PHP解析器来处理，那么他会把这个请求简单处理后交给PHP解析器。此时CGI便是规定了要传什么数据／以什么格式传输给php解析器的协议。当web server收到/index.php这个请求后，会启动对应的CGI程序，这里就是PHP的解析器。
接下来PHP解析器会解析php.ini文件，初始化执行环境，然后处理请求，再以CGI规定的格式返回处理后的结果，退出进程。

### CGI与FastCGI相比较
两者的主要差距在于性能瓶颈。前者接受到一个请求就会fork一个进程，后者则是事先启动一部分的worker进程。

### CGI工作原理

1.CGI针对每个http请求都是fork一个新进程来进行处理，接着读取php.ini文件配置信息，初始化执行环境等。

2.然后这个进程会把处理完的数据返回给web服务器，最后web服务器把内容发送给用户。

3.刚才fork的进程也随之退出。 

4.如果下次用户还请求动态资源，那么web服务器又再次fork一个新进程，周而复始的进行。

### FastCGI工作原理

1.Fastcgi则会先fork一个master，解析配置文件，初始化执行环境,然后再fork多个worker。

2.当请求过来时，master会传递给一个worker，然后立即可以接受下一个请求。这样就避免了重复的劳动，效率自然是高。

3.而且当worker不够用时，master可以根据配置预先启动几个worker等着；当然空闲worker太多时，也会停掉一些，这样就提高了性能，也节约了资源。这就是Fastcgi的对进程的管理。大多数Fastcgi实现都会维护一个进程池。注：swoole作为httpserver，实际上也是类似这样的工作方式。

## Nginx与php-fpm通信分析

Nginx与php-fpm通信有两种方式，一种是通过tcp socket和 unix socket。前者是通过ip:端口方式进行通信，后者是通过php启动生成的socket文件进行通信。因此tcp socket的方式可以将两者分布再不同的机器上，只要Nginx能够连接到php服务器的端口即可。后者的方式，是统一主机上进行通讯的方式，因此两者只能再同一台机器上面。

## tcp socket和 unix socket两者的优缺点 

由于<font color='red'> Unix socket 不需要经过网络协议栈，不需要打包拆包、计算校验和、维护序号和应答等，只是将应用层数据从一个进程拷贝到另一个进程</font>。所以其效率比 tcp socket 的方式要高，可减少不必要的 tcp 开销。不过，unix socket 高并发时不稳定，连接数爆发时，会产生大量的长时缓存，在没有面向连接协议的支撑下，大数据包可能会直接出错不返回异常。而 tcp 这样的面向连接的协议，可以更好的保证通信的正确性和完整性。

## 如何选择tcp socket与unix socket

1.由于tcp方式相对unix的方式，并发量更高，因此针对并发量高的项目，建议采用tcp方式，现在Nginx配置示例文件默认的也是tcp方式。

2.使用unix方式，可以优化的点，就是将socket文件放在/dev/shm目录下面，至于为什么放在这个目录可以参考.https://www.linuxidc.com/Linux/2014-05/101818.htm。大致的意思，就是该目录下面的文件是不是存储再硬盘中的，而是存储再内存中的。至于硬盘读取和内存读取，谁快谁慢，肯定是内存最快了。

3.使用unix方式可以使用backlog，backlog的介绍，可以参考该文章。https://blog.csdn.net/raintungli/article/details/37913765。具体的配置可以参考下面。
nginx 配置:
```shell
server {
  listen 80 default backlog = 100;
  }
```
php-fpm 配置:
```shell
listen.backlog = 1000
```

## 配置示例
Nginx与PHP通信的方式，取决于PHP启动的方式，下面直接演示两种方式的配置文件，后面记录PHP启动方式的配置
### tcp socket
```shell
server {
        listen       80;
        server_name  laravel_demo.com;
        charset utf-8;
        access_log  logs/laravel_demo.com.access.log;
        root   /usr/local/var/www/laravel64_demo/public/;
        index  index.php index.html index.htm;

        error_page   500 502 503 504  /50x.html;

        location / {
            if (!-e $request_filename) {
                rewrite ^(.*)$  /index.php?s=$1 last;
                break;
            }
        }
        ### 此处就是Nginx与tcp socket通信配置，这里部署的同一台及其，因此IP就是127.0.0.1
        location ~ \.php$ {
                fastcgi_pass   127.0.0.1:9000;
                fastcgi_index  index.php;
                fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
                include        fastcgi_params;
            }
    }
```
### unix socket
```shell
server {
        listen       80;
        server_name  laravel_demo.com;
        charset utf-8;
        access_log  logs/laravel_demo.com.access.log;
        root   /usr/local/var/www/laravel64_demo/public/;
        index  index.php index.html index.htm;

        error_page   500 502 503 504  /50x.html;

        location / {
            if (!-e $request_filename) {
                rewrite ^(.*)$  /index.php?s=$1 last;
                break;
            }
        }
        ### 此处就是Nginx与unix socket通信配置,我的socket文件在/usr/run/目录下
        location ~ \.php$ {
                fasrcgi_pass /usr/tmp/php-fpm.sock;
                fastcgi_index  index.php;
                fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
                include        fastcgi_params;
            }
    }
```
上面演示了具体的配置示例，下面演示的是PHP启动方式的配置，这里需要和上面Nginx的配置一致，如果是tcp方式就采用tcp方式启动，如果是unix方式采用socket文件的方式启动。
```shell
# tcp方式启动
listen = 127.0.0.1:9000
# unix方式启动
listen = /usr/tmp/php-fpm.sock;
```
采用上面的两种方式中的任意一种，直接使用php-fpm方式启动php即可。
1.tcp方式启动的效果图
![](http://qiniucloud.qqdeveloper.com/iShot%20%20%20%20%202019-12-03%20PM11.32.29.png)

2.unix方式启动的效果图
![](http://qiniucloud.qqdeveloper.com/iShot%20%20%20%20%202019-12-03%20PM11.35.04.png)
### 注意
在演示的过程中遇到一个问题就是提示Nginx无法读取php生成的unix socket文件。这中情况是因为权限组导致的。因此再php-fpm的配置配置文件中要设置权限组，同时Nginx也需要设置权限组，再很多的集成开发环境中已经配置好了，因此可以减少此步骤。下面就是示例
Nginx配置文件：
```shell
#配置php-fpm的用户以及所属用户组
user  www www;

worker_processes auto;

error_log  /home/wwwlogs/nginx_error.log  crit;

pid        /usr/local/nginx/logs/nginx.pid;

#Specifies the value for maximum file descriptors that can be opened by this process.
worker_rlimit_nofile 51200;

```
php-fpm配置文件(我们平常可能更多的是配置php.ini的文件，这里需要区分两者之间的区别，php.ini是针对php的配置文件，可以简单的理解为php再编译源码时会用到这里的配置，而关于php这个应用程序执行的情况就会用到php-fpm的配置文件):
```shell
; 配置php-fpm的用户以及所属用户组
user = www
group = www

; The address on which to accept FastCGI requests.
; Valid syntaxes are:
;   'ip.add.re.ss:port'    - to listen on a TCP socket to a specific IPv4 address on
;                            a specific port;
;   '[ip:6:addr:ess]:port' - to listen on a TCP socket to a specific IPv6 address on
;                            a specific port;
;   'port'                 - to listen on a TCP socket to all addresses
;                            (IPv6 and IPv4-mapped) on a specific port;
;   '/path/to/unix/socket' - to listen on a unix socket.
; Note: This value is mandatory.
;listen = 127.0.0.1:9000
listen = /usr/tmp/php-fpm.sock;
```
## 总结

CGI: CGI是 web 服务器与编程语言(PHP、Python)解析器的一种交互协议。
FASTCGI: FASTCGI是 CGI 的升级版本。
PHP-FPM: 是 PHP 的FASTCGI的协议管理器，负责与web 服务器进行通信，通过交互协议将数据传递给 work 进程。