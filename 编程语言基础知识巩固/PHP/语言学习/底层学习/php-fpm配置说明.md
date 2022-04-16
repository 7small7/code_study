[TOC]

## php-fpm全局配置参数

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
## php-fpm的进程进程池配置

```php
# php-fpm的队列长度。
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