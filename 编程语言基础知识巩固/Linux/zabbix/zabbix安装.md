## 安装说明

本文将基于Centos7和lnmp环境搭建zabbix。zabbix安装的版本是5.4的版本，PHP的版本是7.3(这里需要注意一下，zabbix5.x的版本要求PHP的版本必须是>=7.2)，MySQL的版本是5.7。

很多文章和官网都提及到了安装MySQL和PHP等环境。但是实际情况，很多时候服务端其实已经安装好了PHP和MySQL的环境，不需要安装PHP和MySQL环境了。所以本文也是基于这种情况下来搭建zabbix环境。

## 整体流程

安装lnmp环境->安装zabbix服务端->安装zabbix客户端->配置zabbix前端->配置主机信息。

## 安装流程

### 安装lnmp环境。

lnmp属于PHP集成环境按照包，通过该包可以一键PHP相关的服务环境。安装的操作也很简单，直接根据官网的介绍安装即可。[lnmp安装包与安装说明](https://lnmp.org/download.html)，这里就不演示具体的安装流程了。

### 安装zabbix服务端

我这里将服务端和客户端都安装在一台机器上，在安装zabbix服务端时，自动将客户端安装完成。打开zabbix的官网，[下载源码](https://www.zabbix.com/cn/download_sources)。
![Snipaste_2021-10-06_15-57-41](https://gitee.com/bruce_qiq/picture/raw/master/2021-10-6/1633507094200-Snipaste_2021-10-06_15-57-41.png)
```shell
# 将源码包下载到/home/wwwroot/目录下面
[root@centos wwwroot]# wget https://cdn.zabbix.com/zabbix/sources/stable/5.4/zabbix-5.4.5.tar.gz 
# 解压zabbix源码包
[root@centos wwwroot]# tar -zxvf zabbix-5.4.5.tar.gz
# 查看当前源码包
[root@centos wwwroot]# ll
总用量 23764
drwxr-xr-x. 13 mysql mysql     4096 10月  6 03:08 zabbix-5.4.5
-rw-r--r--.  1 root  root  24324118 9月  29 15:46 zabbix-5.4.5.tar.gz
# 查看源码包内容
[root@centos zabbix-5.4.5]# ls
aclocal.m4  build      conf          config.status  configure.ac  depcomp  install-sh  Makefile.am  misc     README  ui
AUTHORS     ChangeLog  config.guess  config.sub     COPYING       include  m4          Makefile.in  missing  sass
bin         compile    config.log    configure      database      INSTALL  Makefile    man          NEWS     src
```
准备安装源码之后，就可以开始安装zabbix。
```shell
# 安装相关依赖包
yum install net-snmp-devel libxml2-devel libcurl-devel libevent-devel  mysql-devel
```
```shell
# 编译安装并指定相关依赖扩展
./configure --enable-server --enable-agent --with-mysql 
--with-net-snmp --with-libcurl --with-libxml2
make && make install
```
创建一个zabbix用户组和用户。
```shell
groupadd zabbix
useradd -g zabbix zabbix
```
> 为什么要给zabbix创建一个单独的用户，很简单，是因为安全问题。为了安全考虑zabbix只使用普通用户运行，假如你当前用户叫ttlsa，那么你运行他，他便使用ttlsa身份运行。但是如果你在root环境下运行zabbix，那么zabbix将会主动使用zabbix用户来运行。但是如果你的系统没有名叫zabbix的用户，你需要创建一个用户。

安装完成之后，就可以查看响应的配置文件和对应的服务。
```shell
# 服务命令文件
[root@centos wwwroot]# cd /usr/local/sbin/
[root@centos sbin]# ll
总用量 13788
-rwxr-xr-x. 1 root root  1879944 10月  6 03:08 zabbix_agentd
-rwxr-xr-x. 1 root root 12236688 10月  6 03:08 zabbix_server
```
下面是配置文件。zabbix_agentd.conf是客户端配置文件，zabbix_server.conf是服务端的配置文件。xxx.d则是配置文件的目录，后期增加对应的配置，直接在该目录下面去创建，会被自动加载进来。
```shell
[root@centos sbin]# cd /usr/local/etc/
[root@centos etc]# ll
总用量 48
-rw-r--r--. 1 root root 15862 10月  6 14:21 zabbix_agentd.conf
drwxr-xr-x. 2 root root  4096 10月  6 03:08 zabbix_agentd.conf.d
-rw-r--r--. 1 root root 24469 10月  6 03:50 zabbix_server.conf
drwxr-xr-x. 2 root root  4096 10月  6 03:08 zabbix_server.conf.d
```
> 安装完成服务端之后，我们暂且不启动服务。
### 配置数据库

在下载的源码包中，我们可以看到有一个database的目录，里面有如下的文件信息。这些信息就是我们对应的数据库信息。由于我们使用的是MySQL数据库，因此只需要关注MySQL目录下面的文件即可。
```shell
[root@centos database]# ll
总用量 76
drwxr-xr-x. 2 mysql mysql  4096 9月  27 16:08 elasticsearch
-rw-r--r--. 1 root  root  21769 10月  6 03:08 Makefile
-rw-r--r--. 1 mysql mysql   144 9月  27 16:08 Makefile.am
-rw-r--r--. 1 mysql mysql 21953 9月  29 15:34 Makefile.in
drwxr-xr-x. 2 mysql mysql  4096 10月  6 03:08 mysql
drwxr-xr-x. 2 mysql mysql  4096 10月  6 03:08 oracle
drwxr-xr-x. 2 mysql mysql  4096 10月  6 03:08 postgresql
drwxr-xr-x. 2 mysql mysql  4096 10月  6 03:08 sqlite3
```
配置MySQL用户信息，可以使用默认的root用户，也可以根据自己的需要创建一个用户。这里我就默认使用root用户。
```shell
# 登录MySQL
[root@centos database]# mysql -uroot -p
Enter password: 
# 创建Mzabbix数据库
CREATE DATABASE `zabbix` CHARACTER SET utf8_bin COLLATE utf8_bin;
# 导入默认的数据库信息
source /home/wwwroot/zabbix-5.4.5/database/schema.sql
source /home/wwwroot/zabbix-5.4.5/database/data.sql
source /home/wwwroot/zabbix-5.4.5/database/images.sql
```
### 启动服务端
```shell
# 启动并指定服务配置文件
zabbix_server -c /usr/local/etc/zabbix_server.conf
# 查看服务状态
[root@centos database]# netstat -tnlp|grep zabbix
 tcp        0      0 0.0.0.0:10051           0.0.0.0:*               LISTEN      2286/zabbix_server
```
### 安装zabbix客户端

通过刚才安装zabbix服务端，我们把zabbix的客户端也安装完成。安装的可执行文件和配置文件都和server端一样。
```shell
# 查看配置文件
[root@centos etc]#  cd /usr/local/etc/ && ll
总用量 48
-rw-r--r--. 1 root root 15862 10月  6 14:21 zabbix_agentd.conf
drwxr-xr-x. 2 root root  4096 10月  6 03:08 zabbix_agentd.conf.d
-rw-r--r--. 1 root root 24469 10月  6 03:50 zabbix_server.conf
drwxr-xr-x. 2 root root  4096 10月  6 03:08 zabbix_server.conf.d
# 查看可执行文件
[root@centos sbin]# cd /usr/local/sbin/ && ll
总用量 13788
-rwxr-xr-x. 1 root root  1879944 10月  6 03:08 zabbix_agentd
-rwxr-xr-x. 1 root root 12236688 10月  6 03:08 zabbix_server
```

### 修改客户端配置
```shell
vim /usr/local/etc/zabbix_agentd.conf
# zabbix服务端ip地址
Server=127.0.0.1
ServerActive=127.0.0.1
# zabbix客户端主机名
Hostname=centos
```
### 启动客户端和服务端
```shell
# 启动服务端
zabbix_server -c /usr/local/etc/zabbix_server.conf
# 启动客户端
zabbix_agentd -c /usr/local/etc/zabbix_agentd.conf
# 查看服务情况
[root@centos sbin]# netstat -tnlp|grep zabbix
tcp        0      0 0.0.0.0:10050           0.0.0.0:*               LISTEN      1038/zabbix_agentd  
tcp        0      0 0.0.0.0:10051           0.0.0.0:*               LISTEN      2286/zabbix_server  
tcp6       0      0 :::10050                :::*                    LISTEN      1038/zabbix_agentd  
```
### 配置zabbix管理端

安装完成之后，我们可以使用IP访问或者域名指向zabbix的管理端界面。我这里使用IP的方式访问。
```shell
[root@centos sbin]# vim /usr/local/nginx/conf/nginx.conf
server
    {
        listen 80 default_server reuseport;
        server_name _;
        index index.html index.htm index.php;
        # 管理端程序的根目录
        root /home/wwwroot/default/zabbix-5.4.5/ui;
    }
```
接下来就可以通过IP直接访问zabbix了。根据界面进行相关配置信息即可。zabbix默认的账号和密码是Admin，zabbix。

