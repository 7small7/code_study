[TOC]

## Mac安装

使用Mac系统安装，可以使用brew包管理工具也可以使用源码编译安装。
### brew包管理工具安装
```shell
# 搜索Redis
kert@192 ~ % brew search redis
==> Formulae
hiredis
redis
redis-leveldb
redis@3.2
redis@4.0
==> Casks
another-redis-desktop-manager     redis    redisinsight
```
```shell
#安装Redis(默认安装最新版本)
kert@192 ~ % brew install redis
Updating Homebrew...
==> Auto-updated Homebrew!
Updated 1 tap (homebrew/services).
No changes to formulae.

Warning: Treating redis as a formula. For the cask, use homebrew/cask/redis
==> Downloading https://mirrors.ustc.edu.cn/homebrew-bottles/bottles/redis-6.0.9.catalina.bottle.tar.gz
######################################################################## 100.0%
==> Pouring redis-6.0.9.catalina.bottle.tar.gz
==> Caveats
To have launchd start redis now and restart at login:
  brew services start redis
Or, if you don't want/need a background service you can just run:
  redis-server /usr/local/etc/redis.conf
==> Summary
🍺  /usr/local/Cellar/redis/6.0.9: 13 files, 3.9MB
```
管理服务
```shell
# 启动服务
 brew services start redis
 # 重启服务
 brew services restart redis
 # 停止服务
 brew services stop redis
```
### 源码安装

1. 直接到官网下载Redis源码，下载地址：https://redis.io/download

2. 解压并安装
```shell
tar xzf redis-6.0.8.tar.gz
cd redis-6.0.8
make && make install
```

## Linux安装

### yum包管理工具安装

```shell
[root@VM-51-113-centos ~]# yum install redis
Loaded plugins: fastestmirror, langpacks
Loading mirror speeds from cached hostfile
................................
Resolving Dependencies
--> Running transaction check
---> Package redis.x86_64 0:3.2.12-2.el7 will be installed
--> Processing Dependency: libjemalloc.so.1()(64bit) for package: redis-3.2.12-2.el7.x86_64
--> Running transaction check
---> Package jemalloc.x86_64 0:3.6.0-1.el7 will be installed
--> Finished Dependency Resolution
.................................
Total download size: 648 k
Installed size: 1.7 M
...........................
Installed:
redis.x86_64 0:3.2.12-2.el7
Dependency Installed:
jemalloc.x86_64 0:3.6.0-1.el7
```

### 源码安装

由于Mac和Linux都是基于Unix，安装方式一致。依次执行下面的每一行命令即可。
```shell
wget http://download.redis.io/releases/redis-6.0.8.tar.gz
tar xzf redis-6.0.8.tar.gz
cd redis-6.0.8
make && make install
```

## 安装优化

在文章中提到的源码安装，我们都是采用的直接使用源码安装，这里有个小技巧，可以对日后的升级有所帮助。就是对源码文件创建一个软连接。以后升级直接对源码目录替换即可。下面的命令就是将下载的Redis源码创建一个软件列到 /opt/redis/6.0.8目录。
```shell
ln -s redis-6.0.8/ /opt/redis/6.0.8/ 
cd /opt/redis/6.0.8/ && make && make install
```
> Linux创建软连接 ln -s 源文件 目标文件