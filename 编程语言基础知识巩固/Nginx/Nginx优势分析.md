## NGINX介绍
Nginx是一款轻量级的 Web 服务器/反向代理服务器及电子邮件(IMAP/POP3)代理服务器，并在一个BSD-like 协议下发行。其特点是占有内存少，并发能力强，事实上NGINX的并发能力确实在同类型的网页服务器中表现较好，中国大陆使用nginx网站用户有:百度、京东、新浪、网易、腾讯、淘宝等。

## NGINX 优缺点分析

### 优点分析
1.高并发量：根据官方给出的数据，能够支持高达 50,000 个并发连接数的响应。

2.内存消耗少：处理静态文件，同样起web 服务，比apache 占用更少的内存及资源，所以它是轻量级的(至于区别在哪？后面内容进行分析)。

3.简单稳定：一般在使用 Nginx 中，针对单个域名配置文件配置接口，学习成本很低。相比 Apache 配置简单很多。

4.模块化程度高：Nginx是高度模块化的设计，编写模块相对简单，包括 gzipping, byte ranges, chunked responses,以及 SSI-filter 等 filter，支持 SSL 和 TLSSNI。

5.支持Rwrite重写规则：能够根据域名、URL的不同， 将HTTP请求分发到不同的后端服务器群组。

6.低成本：Nginx可以做负载均衡，且Nginx是开源免费的，如果使用商业软件F5等硬件来做负载均衡，硬件成本比较高。

7.支持多系统：Nginx是由C语言开发，适用于各个平台。

### 缺点分析

1.动态处理能力较差：Nginx处理静态文件好,耗费内存少，但是处理动态页面则很鸡肋。
这一点怎么来说呢？个人觉得这一点不算是 Nginx 的弱点，但是从理论分析，好像有那么回事一样。

2.rewrite弱：虽然Nginx支持rewrite功能，但是相比于Apache来说，Apache比Nginx 的rewrite 强大。Apache 配置 rewrite 是通过项目下面的.htaccess 文件做配置还有就是打开 Apache 配置中的 rewrite 模块。 而 Nginx 则需要在做域名配置时，对 rewrite 做配置。
```shell
// Apache 需要将该注释给去掉
#LoadModule rewrite_module libexec/apache2/mod_rewrite.so
// Nginx 配置 ThinkPHP
 location / {
      if (!-e $request_filename) {
          rewrite  ^(.*)$  /index.php?s=/$1  last;
          break;
      }
}
```

### Nginx动态处理能力分析

####  1.Nginx 工作原理分析
经常在网上会看到一些文章，提及到 Nginx 的动态处理能力差，暂且咱们不说这是不是正确的一个观点，我们先看看 Nginx 在执行过程中的一个流程。(文章都以 PHP 的配置作为示例配置。)
[![Dp34Vf.md.png](http://qiniucloud.qqdeveloper.com/iShot%20%20%20%20%202019-12-04%20AM12.11.26.png)](http://qiniucloud.qqdeveloper.com/iShot%20%20%20%20%202019-12-04%20AM12.11.26.png)
这里说明一下 FastCGI 和 CGI 是什么？参考链接:http://mindoc.qqdeveloper.com/docs/php/php-1ch9talu8n9ms
```shell
配置 FastCGI
location ~ \.php$ {
      # 根据你 PHP 启动的方式决定是 socket 还是 tcp 方式
      #  fasrcgi_pass /usr/tmp/php-fpm.sock;
      fastcgi_pass   php73:9000;
      include        fastcgi-php.conf;
      include        fastcgi_params;
}
```
#### 2.Apache 如何与 PHP 通行的？
通过上面的参考链接，我们可以知道Nginx 要与 PHP 相关的东西通信是通过 FastCGI 这种协议进行通信，然后 php-fpm 进程管理器接收到信息之后，在提交给 php-fpm 下的 work 进程进行执行。然后 Apache 是直接将 PHP 作为一个模块加载到自身配置中，这样处理的速度更快。
```shell
开启 PHP 模块支持
LoadModule php7_module libexec/apache2/libphp7.so
```
#### 3.总结
通过上面的分析，从理论上说 Apache 确实较好一些。其实无论是mod_php、还是FastCGI，都有其自己的优势。以前在FastCGI技术还不成熟的时候，自然是mod_php稳定、处理速度更快一些，可是社会是不断在向前进步的，现如今FastCGI技术已经非常成熟了，网上也有很多人做了相关的测试，说是FastCGI比mod_php更稳定、速度更快。我个人认为，如果是单机部署的话，可考虑使用mod_php方式，因为毕竟多启一个进程对系统而言就多了一些资源消耗；如果分开部署的话，可考虑使用FastCGI，现在越来越多的人使用nginx+php架构了。

### 功能对比
一提及到 Nginx，可能都会想到这几个词，反向代理、负债均衡、高性能的 web 服务器。今天简单看了一下 Apache 其实也是支持反向代理功能的。具体没配置过，看着倒是没 Nginx 的简单。
```shell
#LoadModule proxy_module libexec/apache2/mod_proxy.so
#LoadModule proxy_connect_module libexec/apache2/mod_proxy_connect.so
#LoadModule proxy_ftp_module libexec/apache2/mod_proxy_ftp.so
#LoadModule proxy_http_module libexec/apache2/mod_proxy_http.so
#LoadModule proxy_fcgi_module libexec/apache2/mod_proxy_fcgi.so
#LoadModule proxy_scgi_module libexec/apache2/mod_proxy_scgi.so
#LoadModule proxy_uwsgi_module libexec/apache2/mod_proxy_uwsgi.so
#LoadModule proxy_fdpass_module libexec/apache2/mod_proxy_fdpass.so
#LoadModule proxy_wstunnel_module libexec/apache2/mod_proxy_wstunnel.so
```