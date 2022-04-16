[TOC]

## 分享背景

由于一直在虚拟机的状态下开发 PHP,尝试一下 mac 本地搭建环境.mac 本身是自带 Apache+php 的,在低版本的 mac 系统中,mac 中的 php 版本是 5.6 的版本.
本文分享的是在 mac 的 10.13 版本.前提是本地安装了 brew 包管理工具,如果还没安装的可以参考官网文档安装,[传送门](https://brew.sh/).

## 安装步骤

1.停止本地的 Apache 服务

```shell
sudo apachectl stop
```

2.安装 NGINX

```shell
brew install nginx
```

3.启动 NGINX

```shell
// 查看帮助命令
nginx -v
// 重启命令
nginx -s reload
// 每次我们配置了虚拟主机之后，都要使用下面的命令进行重启
sudo nginx -s reload
```

如果出现如下错误信息

```shell
nginx: [error] open() "/usr/local/nginx/logs/nginx.pid" failed (2: No such file or directory)
```

解决思路是(前提是使用的 brew 包管理工具安装的，如果是其他的安装方式根据情况而定)

```shell
mkdir /usr/local/Cellar/nginx/1.15.12/logs
```

4.访问 NGINX,打开浏览器,输入如下网址,正确的情况就可以看到如下的截图.

```shell
127.0.0.1:8080
```

![5](http://qiniucloud.qqdeveloper.com/mweb/5.png)
5.NGINX 项目目录介绍,<font color='red'>通过上面的步骤,就表示 NGINX 已经完成了.这里有几个文件,我们需要关注一下.</font>

```shell
1.nginx配置目录
/usr/local/etc/nginx
2.nginx的项目根目录
/usr/local/var/www
```

6.配置 php(由于 mac 的高版本中已经内置了 PHP7.1 的版本,该文章也是基于这个基础上操作的.后续完善该文章,实现一个多版本的切换.) 1.去掉 nginx.conf 中如下代码中的注释(在去掉之前最好备份一份 `shell cp nginx.conf nginx.conf.bak)`)

```shell
location ~ \.php$ {
    fastcgi_pass   127.0.0.1:9000;
    fastcgi_index  index.php;
    fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
    include        fastcgi_params;
}
```

该段代码的主要意思就是让 nginx 能够解析到 php,不然你去访问 php 的文件,nginx 会默认下载该 php 文件.在 Apache 中是以模块的方式加载的 php,就是去掉

```shell
LoadModule php_module libexec/apache2/libphp7.so
```

2.修改 1 中的部分配置
原配置中的值为

```php
/scripts$fastcgi_script_name
```

需要修改为

```php
$document_root$fastcgi_script_name
```

该代码主要的意思就是让 nginx 能够识别到 nginx 配置文件中的 root 项.不然会报 not find file 错误信息.
<font color='red'>重启 nginx 服务,`php nginx -s reload`</font> 3.配置 php-fpm 文件

```php
sudo cp /private/etc/php-fpm.conf.default /private/etc/php-fpm.conf
```

修改我们新复制的 php-fpm.conf 文件中的如下配置项目.修改为下面的示例

```php
pid=/var/run/php-fpm.pid
error_log=/var/log/php-fpm.log
```

4.启动 php-fpm 服务

```php
sudo php-fpm
```

启动服务的时候可能会遇到下面的问题,反正我是遇到了(下面的路径根据你图片指示的路径来定,可能有的环境路径不一致).解决办法是
![6](http://qiniucloud.qqdeveloper.com/mweb/6.png)

```shell
cp /data/server/php7/etc/php-fpm.d/www.conf.default /data/server/php7/etc/php-fpm.d/www.conf
```

然后在启动一次,即可. 5.编写测试文件,我们找到/usr/local/var/www 目录,创建一个 test.php 的文件.写入如下代码.

```php
phpinfo();
```

6.测试效果
打开浏览器,访问http://localhost/test.php ,即可看到如下phpinfo界面。

7.如何实现 PHP 版本升级
在上面的示例中，我们安装的 PHP 版本为 7.1 的版本.但在一些情况下面，我们可能需要更高的版本.下面升级的示例则是使用 brew 包管理工具进行升级.
查看 brew 可安装的 PHP 包

```shell
brew search php
```

![Screen Shot 2019-07-12 at 16.34.54](http://qiniucloud.qqdeveloper.com/mweb/Screen%20Shot%202019-07-12%20at%2016.34.54.png)
其中带 √ 的是当前安装的版本,接下里我们就可以安装新版本

```shell
brew install php72
```

启动服务

```shell
sudo brew services start php72
```

这样我们就不需要再次去执行`shell sudo php-fpm`,否则会发生端口冲突.这种情况下，系统默认的 PHP 版本还是 7.1,同时我们也便于实现 PHP 多版本管理.
新 PHP 版本安装位置

```shell
/usr/local/Cellar/php@7.2
```
启动新版本PHP服务(启动php7.2版本之前，需要关掉之前PHP的服务，可以选择重启电脑来关闭)
```shell 
// 重启服务
brew services restart php@7.2
启动服务
brew services start php@7.2
已守护进程模式启动
php-fpm
```
新版本PHP配置文件
```shell
使用 find / -name php.ini命令发现，有如下两个配置文件，我们这里以第一个为主，以后的配置可以直接在第一个配置文件进行修改
/usr/local/etc/php/7.2/php.ini
/usr/local/Cellar/php@7.2/7.2.20/.bottle/etc/php/7.2/php.ini
```
新版本PHP扩展文件
```shell
这里我们需要提及到常用的两个文件,一个是phpize,一个是php-config，这两个文件都需要我们在安装php扩展时需要用到，他们的目录均在如下目录
/usr/local/Cellar/php@7.2/7.2.20/bin
```