## 安装说明

安装rabbitmq有很多中方式。如果看不懂官网文档的开发者，可能在实际安装中会遇到很多的问题，本文将演示一下centos7.x的版本，如何安装。

安装的方式是使用rpm包进行安装。安装所需要的依赖包分别有erlang、socat。因为rabbitmq是erlang开发的，因此需要该环境。socat是一个网络通讯工具，rabbitmq内部的通讯依赖该包，因此也需要安装该依赖库。
> 本文介绍的安装方式，是基于centos7.x全新的环境。环境不同遇到的问题也会不同，本文进攻参考。

## 安装准备

打开rpm安装包下载网站。我这里使用的[packagecloud.io](https://packagecloud.io/rabbitmq/)。会出现如下的界面，我们只需要下载erlang和rabbitmq就可以了。

![Snipaste_2021-12-05_12-14-03](https://gitee.com/bruce_qiq/picture/raw/master/2021-12-5/1638677669159-Snipaste_2021-12-05_12-14-03.png)

找到合适的版本，点击包名称，就会跳转到类似下面的界面。本文安装的版本是:
```shell
erlang版本：erlang-23.3.4.4-1.el7.x86_64.rpm
rabbitmq版本：rabbitmq-server-3.8.26-1.el7.noarch.rpm
```

![Snipaste_2021-12-05_12-15-31](https://gitee.com/bruce_qiq/picture/raw/master/2021-12-5/1638677740749-Snipaste_2021-12-05_12-15-31.png)

看到如上的界面之后，我们直接点击右上角的`download`按钮就可以将rpm包下载到本地，然后你在上传到服务器上就可以了。
> 通过curl的方式，会发现很慢，因此推荐使用本文的方式。

### 安装

首先我们安装erlang，在安装socat，最后安装rabbitmq。如果你不安装前面两个，安装rabbit也会进行提示。类似下面的错误提示信息：
```shell
警告：rabbitmq-server-3.8.26-1.el7.noarch.rpm: 头V4 RSA/SHA512 Signature, 密钥 ID 6026dfca: NOKEY
错误：依赖检测失败：
	socat 被 rabbitmq-server-3.8.26-1.el7.noarch 需要
```
```shell
rpm erlang-23.3.4.4-1.el7.x86_64.rpm
```
```shell
yum install socat
```
可能在安装socat时，会提示下面的信息，此时会报错没有socat包或是找不到socat包。直接执行`yum install -y install epel-release`。如果还是不行，可以直接进行源码安装。
下载socat源代码包：http://www.dest-unreach.org/socat/download/
编译安装  把下载的软件包解压后按照传统的方式编译安装：  
```shell
./configure     #需要gcc
make  
make install  
```
在编译的过程中可能遇到如下错误： `/sbin/sh: fipsld:command not found`   
解决方法有两种：  
第一种是禁用fips，使用如下命令配置：  ./configure --disable-fips  
第二种是安装fips，首先到网站http://www.openssl.org/source/ 下载openssl-fips安装包，然后解压安装：  
```shell
./config
make  
make install  
```
安装完成之后，就可以直接安装rabbitmq-server了。
```shell
rpm -ivh rabbitmq-server-3.8.26-1.el7.noarch.rpm
```
显示100%则表示安装成功。

## 相关命令
```shell
# 启动服务
systemctl start rabbitmq-server.service
# 停止服务
systemctl stop rabbitmq-server.service
# 重启服务
systemctl restart rabbitmq-server.service
# 设置开机启动
chkconfig rabbitmq-server on
```
安装web管理界面。
```shell
rabbitmq-plugins enable rabbitmq_management
systemctl restart rabbitmq-server.service
```

## 创建用户

安装并启动服务之后，就可以通过IP:15672进行访问，就可以正常访问了。不过会出现下面的提示信息：

![Snipaste_2021-12-05_14-36-02](https://gitee.com/bruce_qiq/picture/raw/master/2021-12-5/1638686190182-Snipaste_2021-12-05_14-36-02.png)
> rabbitmq的默认账户和密码是:guest,guest。这里的提示信息就是说，guest账户只能通过localhost进行访问。

为了解决该问题，我们需要创建一个独立的用户，并给其超级管理员的权限。
```shell
# 创建一个用户和设置密码
rabbitmqctl add_user admin 123456
# 设置角色
rabbitmqctl set_user_tags admin administrator
# 设置权限
rabbitmqctl set_permissions  admin ConfP WriteP ReadP
```
> 本文在演示中，创建的账户和密码分别是admin、123456。

配置好之后，我们就可以使用admin账户进行登录了。登录之后，就可以看到如下的界面：
![Snipaste_2021-12-05_14-41-25](https://gitee.com/bruce_qiq/picture/raw/master/2021-12-5/1638711843979-Snipaste_2021-12-05_14-41-25.png)

## 相关配置

按照上面的流程，默认的情况下，我们就能正常访问与使用了。既然rabbitmq作为一个服务，肯定有一些配置文件。在有的版本中，安装之后会有一份rabbit-server.conf.example的文件。但是在我这个版本，还是在个人的安装方式问题。发现是没有的。需要到rabbitmq官方拉取一份。具体的地址。
```shell
wget https://github.com/rabbitmq/rabbitmq-server/blob/master/deps/rabbit/docs/rabbitmq.conf.example
```
将该文件放在`/etc/rabbitmq`目录下面，并且重命名为`rabbitmq.conf`。每次启动，rabbitmq会自动去加载该文件的。创建好之后，我们只需要修改一下几个配置项就可以了。
```shell
# 这几个值，是默认自带的
management.tcp.port = 15672
management.tcp.ip   = 0.0.0.0
# 下面这个目录可以根据自己的需要指定，默认值是没有的
management.http_log_dir = /usr/local/rabbitmq/logs/
```
> 默认情况下，直接把文件前面的注释去掉就可以了。