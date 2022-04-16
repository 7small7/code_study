## xdebug介绍

Xdebug是PHP的一个扩展，方便我们调试PHP应用程序的执行流程信息。使用过JavaScript中的debug，应该就能很好的理解xdebug。总结，大致有如下的功能：

1. 它包含一个用于IDE的调试器。
2. 它升级了PHP的var_dump()函数。
3. 它为通知，警告，错误和异常添加了堆栈跟踪。
4. 它具有记录每个函数调用和磁盘变量赋值的功能。
5. 它包含一个分析器。
6. 它提供了与PHPUnit一起使用的代码覆盖功能。

本文变针对Mac上进行xdebug的安装与简单调试。

## 环境

环境都是在Mac上运行的，使用到了Apache、PHP。

PHP版本：7.4.20。xdebug版本：3.1.1。

## 安装流程

### PHP安装

默认的Mac是自带PHP环境的，由于版本的需要，因此不会使用该版本。我们使用brew进行安装。我们使用brew搜索有哪些PHP版本。
```shell
 kert@192  /usr/local/Cellar/php@7.4/7.4.20  brew search php
==> Formulae
brew-php-switcher     php-code-sniffer      php-cs-fixer@2        php@7.3               phpbrew               phpmd                 phpstan
php                   php-cs-fixer          php@7.2               php@7.4 ✔             phplint               phpmyadmin            phpunit
==> Casks
eclipse-php
```
我们可以看到与PHP相关的包，直接使用`brew install php@7.4`安装即可。所有的操作都是一步完成。安装完成之后，相关的配置文件，如下：
```shell
kert@192   cd /usr/local/etc/php/7.4
kert@192  /usr/local/etc/php/7.4  ll
total 168
drwxr-xr-x  7 kert  admin    224 11 14 02:26 .
drwxr-xr-x  3 kert  admin     96 11 13 23:42 ..
drwxr-xr-x  3 kert  admin     96 11 13 23:42 conf.d
-rw-r--r--  1 kert  admin   1361 11 13 23:42 pear.conf
-rw-r--r--  1 kert  admin   5407 11 13 23:42 php-fpm.conf
drwxr-xr-x  3 kert  admin     96 11 13 23:42 php-fpm.d
-rw-r--r--  1 kert  admin  73231 11 14 02:31 php.ini
```
我们可以看到，php.ini、pear.conf等相关的配置文件。以后修改PHP的配置文件，以及php-fpm的进程文件都是在这里进行配置即可。下面是PHP相关的cli工具。
```shell
cd /usr/local/Cellar/php@7.4/7.4.20/bin
kert@192  /usr/local/Cellar/php@7.4/7.4.20/bin  ll
total 78792
drwxr-xr-x  12 kert  admin       384 11 13 23:42 .
drwxr-xr-x  16 kert  admin       512 11 14 00:19 ..
-r-xr-xr-x   1 kert  admin       960 11 13 23:42 pear
-r-xr-xr-x   1 kert  admin       981 11 13 23:42 peardev
-r-xr-xr-x   1 kert  admin       894 11 13 23:42 pecl
lrwxr-xr-x   1 kert  admin         9  6  1 23:42 phar -> phar.phar
-rwxr-xr-x   1 kert  admin     14910 11 13 23:42 phar.phar
-r-xr-xr-x   1 kert  admin  13398608 11 13 23:42 php
-r-xr-xr-x   1 kert  admin  13346744 11 13 23:42 php-cgi
-r-xr-xr-x   1 kert  admin      7095 11 13 23:42 php-config
-r-xr-xr-x   1 kert  admin  13542824 11 13 23:42 phpdbg
-r-xr-xr-x   1 kert  admin      4575 11 13 23:42 phpize
```
对PHP的配置操作之后，记住一定要重启PHP服务，有时候如果没有生效的情况，最好是重启一下Apache(下面有写)服务。
```shell
# 重启服务
brew services restart php@7.4
# 启动服务
brew services start php@7.4
# 停止服务
brew services stop php@7.4
```
### Apache安装

同样，Mac也是自带Apache环境的。你可以直接使用，也可以自定安装。但是推荐不要用Mac自带的Apache。在使用中有下面几个原因：
1. 自带的Apache要解析PHP，直接在/etc/apache2/httpd.conf配置文件中将php.so的注释放开就行了。他默认解析的是Mac自带的PHP，后面想改成上面自定义的PHP版本会遇到很多问题。

2. 自带的Apache相关配置，很多文件会遇到权限问题，而且是read-only权限。

因此，直接使用brew安装Apache即可。首先使用`sudo apachectl -k stop`命令，将自带的Apache给禁用。操作前之后，我们就可以安装Apache了。直接使用下面的命令安装即可。
```shell
brew install nghttp2
```
安装完之后，相关的配置文件都会放在下面的目录。后期需要配置，更改下面的配置文件即可。
```shell
cd /usr/local/etc/httpd
kert@192  /usr/local/etc/httpd  ll
total 200
drwxr-xr-x   8 kert  admin    256 11 14 02:06 .
drwxrwxr-x  17 kert  admin    544 11 14 01:35 ..
drwxr-xr-x  14 kert  admin    448 11 14 01:39 extra
-rw-r--r--   1 kert  admin  21475 11 14 02:06 httpd.conf
drwxr-xr-x   6 kert  admin    192 11 14 02:14 logs
-rw-r--r--   1 kert  admin  13064 11 14 01:35 magic
-rw-r--r--   1 kert  admin  60847 11 14 01:35 mime.types
drwxr-xr-x   4 kert  admin    128 11 14 01:35 original
```
默认情况是使用的8080，由于我本地8080的端口已经被占用，因此需要修改默认的端口地址。直接对httpd.conf文件接口。
```shell
# 打开该文件，直接搜索listen，将8080改为8088即可。
Listen 8088
# 如果你需要修改程序的根目录，可以修改如下的配置。如果不需要，可以使用默认的配置，后面将你的PHP代码放到默认配置的目录也可以。
DocumentRoot "/Users/kert/code/php_dnmp/www"
<Directory "/Users/kert/code/php_dnmp/www">
    Options Indexes FollowSymLinks
    AllowOverride None
    Require all granted
</Directory>
# 修改权限，否则访问服务的时候会出现403的错误信息。
<Directory />
    AllowOverride none
    Require all granted（将之前的deny改成granted）
</Directory>
```
开启域名配置文件。后面我们会创建不同的项目，每一个项目都会配置一个域名。直接在下面的文件中进行配置就可以了，配置示例如下：
```shell
# 首先我们要修改一下httpd.conf，加载域名配置文件。将下面的配置的注释打开就可以了。
# Virtual hosts
Include /usr/local/etc/httpd/extra/httpd-vhosts.conf
# 接下来的所有域名配置就可以在httpd-vhosts.conf中进行操作。
<VirtualHost *:8088>
    ServerAdmin 18228937997@163.com
    DocumentRoot "/Users/kert/code/php_dnmp/www/xdebug"
    ServerName xdebug_dev.com
    ServerAlias xdebug_dev.com
    ErrorLog "/usr/local/etc/httpd/logs/xdebug_dev_logs-error_log.log"
    CustomLog "/usr/local/etc/httpd/logs/xdebug_dev_logs-access_log.log" common
</VirtualHost>
```
Apache的基本配置就操作完成了。接下来，重启一下Apache服务并且在/Users/kert/code/php_dnmp/www/xdebug下面创建一个index.php文件，写入如下代码:
```php
<?php
    phpinfo();
```
```shell
# 重启Apache
brew services restart httpd
# 停止Apache
brew services stop httpd
# 启动Apache
brew services start httpd
```
访问`http://127.0.0.1:8088/index.php`，出现下面的界面，表示我们安装成功了。
![Snipaste_2021-11-14_15-49-59](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636876209119-Snipaste_2021-11-14_15-49-59.png)

### xdebug安装
使用xdebug一定要注意PHP的版本，否则无法使用。xdebug官方是提供了一个检测工具，帮助我们如何选择xdebug的版本。我们可以在终端使用`php -i`将输出的内容，填充到[网站](https://xdebug.org/wizard)的文本框内，检测之后，会自动给出安装的版本，以及安装的流程。

![Snipaste_2021-11-14_15-53-39](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636876467085-Snipaste_2021-11-14_15-53-39.png)

![Snipaste_2021-11-14_15-54-15](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636876477338-Snipaste_2021-11-14_15-54-15.png)

至于xdebug的如何安装，这里就直接省略了，和常规的PHP扩展安装没有什么区别。安装完成之后，需要在php.ini中做如下配置：
```php
[xdebug]
zend_extension=xdebug.so

;启用代码自动跟踪
xdebug.mode = develop,debug,profile,trace
xdebug.profiler_append = 0
xdebug.profiler_output_name = cachegrind.out.%p
xdebug.start_with_request = default|yes|no|trigger
xdebug.trigger_value = StartProfileForMe

;指定性能分析文件的存放目录
xdebug.output_dir ="/usr/local/Cellar/php@7.4/7.4.20/log/xdebug2"
xdebug.show_local_vars=0

;配置端口和监听的域名
xdebug.client_host=9003
xdebug.clent_host="localhost"
```
这里的配置格式可能和你在网络上看到的不太一样，是因为xdebug在高版本中做了一些配置的变更，部分命令的名称做了一些改变。具体可以参考[官方介绍](https://xdebug.org/docs/upgrade_guide#changed-xdebug)。
这里需要注意一下，9003端口号。你可以设置为其他的端口，只需要在后面提及到的PHPstorm中保持一致即可。

## PHPstorm配置

PHPstorm我使用的是2021版本，因此在界面可能有一些不太一样。大家根据自己的情况，找到对应的配置即可。大致的界面如下：

![Snipaste_2021-11-14_16-02-38](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636877053902-Snipaste_2021-11-14_16-02-38.png)

![Snipaste_2021-11-14_16-06-46](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636877220137-Snipaste_2021-11-14_16-06-46.png)

![Snipaste_2021-11-14_16-07-30](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636877288007-Snipaste_2021-11-14_16-07-30.png)


![Snipaste_2021-11-14_16-07-54](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-14/1636877299918-Snipaste_2021-11-14_16-07-54.png)


