[TOC]

## php-fpm是什么

php-fpm是PHP的一个进程管理器。php下面的众多work进程皆有php-fpm进程管理器管理。[参考文章](http://mindoc.qqdeveloper.com/docs/php/php-1ch9talu8n9ms)

## php-fpm的工作原理

php-fpm全名是PHP FastCGI进程管理器。php-fpm启动后会先读php.ini，然后再读相应的conf配置文件，conf配置可以覆盖php.ini的配置。
启动php-fpm之后，会创建一个master进程，监听9000端口（可配置），master进程又会根据fpm.conf/www.conf ，去创建若干子进程，子进程用于处理实际的业务。
当有客户端（比如nginx）来连接9000端口时，空闲子进程会自己去accept，如果子进程全部处于忙碌状态，新进的待accept的连接会被master放进队列里，等待fpm子进程空闲；这个存放待accept的半连接的队列有多长，由listen.backlog 配置。

## 如何查看php-fpm进程与子进程

### 查看php-fpm相关的所有进程。
![](https://oscimg.oschina.net/oscnet/up-9dc7562a3651acae1ccc1e69255da301ec4.png)
这里<kbd>pool www</kbd>皆是php-fpm的子进程，也就是我们常说的work进程。

### 查看php-fpm下面的子进程

通过上面的命令，其实我们能够看出php-fpm相关的进程了，如果我们需要更加直观的查看php-fpm的master进程和work进程，可以通过下面的方式进程查看。
这里的5370则是php-fpm的master进程号。通过上面的命令已经很能直观的得出。
![](https://oscimg.oschina.net/oscnet/up-7035dd3b75139a6ba6f722db9c4154ffa57.png)
通过上面的命令，可以看出php-fpm作为master进程，下面有15个子进程。这里的子进程数都是可以进程自定义配置。通过如下几个参数进程配置:
```shell
pm = dynamic # 动态创建子进程
pm.max_children = 20 # 最大子进程数
pm.start_servers = 15 # 初始化php-fpm进程时，默认的子进程数
```

## php-fpm参数配置说明

### php-fpm全局配置参数

```shell
#php-fpm的运行权限。
#以什么用户什么组的权限来运行池fpm。
user = www
group = www

#php-fpm的运行方式，可以使端口，也可以使socke文件。
#如果是端口则是走tcp，如果是socket则直接读socket文件，这样性能相对更好。
listen = 127.0.0.1:9000 

#拥有socket权限的用户，需要和上面的user、group配置相结合。
#如果采用的端口的方式，则不配置。
listen.owner = www
listen.group = www
listen.mode = 0660

#这是php-fpm端口连接的地址。多个用","隔开。默认任意地址都可以连接。
#例如Nginx和php-fpm不在同一台服务器上，这里的值就是Nginx服务的ip地址。
#当Nginx和php-fpm配置在同一台服务器上，则直接写127.0.0.1即可。
listen.allowed_clients = 127.0.0.1 

#pid进程文件存放的位置，当我们启用一个php服务，
#则会自动创建一个该pid文件，其实我们可以直接把该文件理解理解php-fpm的进程号文件，
#两则是等价的。默认为none。
pid = /opt/remi/php72/root/var/run/php-fpm/php-fpm.pid 

#错误日志位置，默认：安装路径 #INSTALL_PREFIX#/log/php-fpm.log。
#如果设置为syslog，log就会发送给syslogd服务而不会写进文件里。
error_log = /opt/remi/php72/root/var/log/php-fpm/error.log 

#PHP限制的文件扩展名
security.limit_extensions = .php .php3 .php4 .php5 .php7

#系统日志标示，如果跑了多个fpm进程，需要用这个来区分日志是谁的。
syslog.ident = php-fpm 

#日记登记，可选：alert, error, warning, notice, debug。
log_level = notice 

#紧急重启阈值，需要与下面emergency_restart_interval参数一起配置。
emergency_restart_threshold = 60 

# 紧急重启阈值的时间范围。在此参数设置的时间内，
# 出现SIGSEGV或SIGBUS的子进程数超过emergency_restart_threshold参数设置的值。
# 那么fpm就会优雅的重启，值是0表示off这个功能，可用的单位有：s秒,m分,h时,d天。
emergency_restart_interval = 60s 

#设置子进程接受主进程复用信号的超时时间。
process_control_timeout = 0 

#当动态管理子进程时，fpm最多能fork多少个进程，0表示无限制，
# 这是所有进程池能启动子进程的总和，谨慎使用。
process.max = 128 

#设置子进程的优先级，在master进程以root用户启动时有效；
#如果没有设置，子进程会继承master进程的优先级，值范围-19(最高)到20(最低)，默认不设置。
process.priority = -19 

#设置成no用于调试bug，默认为yes。
daemonize = yes 

#master进程最多能打开的文件数量。默认采用系统设置的值。
rlimit_files = 1024 

#master进程核心rlimit限制值；可选unlimited或>=0的整数，默认为系统的值。
rlimit_core = 0

#事件处理机制，默认自动检测，可选值：select，poll，
#epoll(linux>=2.5.44)，kqueue，/dev/poll，port
events.mechanism = epoll 

#fpm想系统发送状态的频率。单位有s,m,h。
#前提是fpm被设置会系统服务。
systemd_interval = 10s 
```
### php-fpm的进程进程池配置

```php
#php-fpm的队列长度。
listen.backlog = 65535 

#php进程池权限，同样要master进程是root用户才有效，
#和上面的全局设置一样，不设置的话会继承master进程的优先级。
process.priority = -19 

#子进程管理方式
#static(静态配置，在启动php-fpm时根据该值创建固定的子进程数量)；
#dynamic(动态配置，在启动php-fpm时根据pm.start_servers的值初始化对应的子进程数，至少一个子进程)；
#ondemand(按需配置，在启动php-fpm时不创建子进程，而是根据请求动态fork子进程)；
pm = dynamic 

#最大子进程数量
pm.max_children = 5 

#初始化子进程数量，与上面的pm = dynamic配置使用。
pm.start_servers = 2 

#服务器闲置时最少保持2个子进程，不够这个数就会创建，只适用动态dynamic管理方式
pm.min_spare_servers = 2 

#服务器闲置时最多要有几个，多了会kill，只适用动态dynamic管理方式
pm.max_spare_servers = 3 

#子进程闲置时间，也就是说子进程没有可处理的任务时，在该之间使就会被killed。
pm.process_idle_timeout = 10s

#每个子进程最大的处理请求数量。在一定程度上可以防止内存泄漏。
pm.max_requests = 500 

#php-fpm状态监控的uri
pm.status_path string

#php-fpm监控页面的 ping 网址。
#如果没有设置，则无法访问 ping 页面。
#该页面用于外部检测php-fpm是否存活并且可以响应请求。请注意必须以斜线开头（/）。
ping.path string

#用于定义ping请求的返回响应。返回为 HTTP 200 的 text/plain 格式文本。默认值：pong。
ping.response string

#设置worker的nice(2)优先级（如果设置了的话）。
#该值从 -19（最高优先级） 到 20（更低优先级）。 
#默认值：不设置
process.priority int

#检测路径时使用的前缀
prefix string

#访问文件日志，没啥用处，比如yii2每次都记录访问index.php，只是记录真实的PHP文件。
access.log = var/log/$pool.access.log 

#php的慢日志
slowlog = var/log/$pool.log.slow 

#慢日志时间阈值
request_slowlog_timeout = 2s 

#单个请求的超时时间，当php.ini设置的最大执行时间未生效，则交由它来处理。
request_terminate_timeout = 3s 

#最大打开句柄数，默认为系统值。
rlimit_files = 1024 

#最多的核心使用数，默认为系统分配。
rlimit_core = 0 
```

## 部分配置演示

### php-fpm的backlog大小设置

php-fpm的backlog大小设置与php-fpm的处理能力有关，而不是越大越好。

当该值设置过大，导致php-fpm处理不过来，nginx那边等待超时，断开连接，报504 gateway timeout错。同时php-fpm处理完准备write 数据给nginx时，发现TCP连接断开了，报“Broken pipe”。

当该值设置过小，nginx之类的client请求，根本进入不了php-fpm的accept queue，报“502 Bad Gateway”错。所以，这还得去根据php-fpm的QPS来决定backlog的大小。计算方式最好为QPS=backlog。
具体可以参考:https://www.jianshu.com/p/3ecc99ebf566

### php-fpm启动模式

php-fpm以socket启动或者端口启动，这两种的方式根据实际情况进行配置。

nginx和php-fpm在同一台服务器上，这时可以直接用unix socket进程间通信，不走tcp端口通信，可以节约创建连接的时间，从而提高性能。sock文件随便创建到哪里都可以，只要fpm有权限在那个目录里写文件，nginx有权限去读就可以。tcp连接会更稳定，因为有tcp协议保证数据的正确性，但是sock有更少的数据拷贝和上下文切换，更少的资源占用。不过只能在nginx和fpm在同一台机器上才能用socket。

如何选择socket启动还是端口启动。

由于tcp方式相对unix的方式，并发量更高，因此针对并发量高的项目，建议采用tcp方式，现在Nginx配置示例文件默认的也是tcp方式。

使用unix方式，可以优化的点，就是将socket文件放在/dev/shm目录下面，至于为什么放在这个目录可以参考.https://www.linuxidc.com/Linux/2014-05/101818.htm  。大致的意思，就是该目录下面的文件是不是存储再硬盘中的，而是存储再内存中的。至于硬盘读取和内存读取，谁快谁慢，肯定是内存最快了。

socket方式启动如何查看socke文件。

socket文件是根据上面提到的pid配置项而定的。我们可以直接使用cat命令，查看进程号。
![](https://oscimg.oschina.net/oscnet/up-fb007baa540e8cf546ce463f1a3358ec804.png)

子进程默认启动数量，通过上面的pm = dynamic 配置，我们知道这种方式是动态配置子进程大小的，同时我们也可以设置默认的子进程数。

```shell
pm = dynamic
pm.max_children = 20
### 默认15个子进程，演示的效果就是上面的shell命令的结果图。
pm.start_servers = 15 
```
当我们尝试设置为3时，显示如下错误信息。
![](https://oscimg.oschina.net/oscnet/up-945676ec3066c3625fa09f42bb55af05325.png)
说明，这里的start_servers配置项和min_spare_servers配置是有一定的关系的。我们设置为最小10，结果就能正常启动php-fpm了。