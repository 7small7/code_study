[TOC]
## 文章简介

一直想开发或者找一个开源的软件，功能就类似看云一样，用来搭建属于一套自己的文档管理系统，将自己平常的东西集中化管理，形成一个手册。于是找到了mindoc这样一款不错的文档管理系统软件。本文大致介绍一下如何安装，基本的功能介绍。

## 功能介绍

主要功能包括文档管理、导出、团队管理、文章管理等功能。特别适合团队或者个人做一个知识库文档管理系统。

## 搭建环境

mindoc支持Linux和Windows系统环境，我这里使用的是Linux系统。系统的配置信息如下:
Centos7.4；
MySQL5.7;
MySQL属于独立搭建的，如果不会的可以参考一下个人之前分享的一篇文章[Linux搭建MySQL环境](http://www.qqdeveloper.com/2020/03/12/install-mysql-for-linux/) ,其他的东西也没了，属于一个纯净的Linux系统。

## 搭建步骤

### 拉取源码

[源码地址](https://gitee.com/longfei6671/godoc/releases)

### 解压并配置环境

```shell
tar -zxvf mindoc_linux_amd64.zip
```
找到config目录下面的app.config文件，如果不存在该文件，应该有一个app.config.example文件，将该文件复制一分，命名为app.conf即可.
```shell
cp app.conf.example app.conf
```
```shell
#支持MySQL和sqlite3两种数据库，如果是sqlite3 则 db_database 标识数据库的物理目录
db_adapter="${MINDOC_DB_ADAPTER||mysql}"
db_host="${MINDOC_DB_HOST||127.0.0.1}"
db_port="${MINDOC_DB_PORT||3306}"
db_database="${MINDOC_DB_DATABASE||mindoc}"
db_username="${MINDOC_DB_USERNAME||root}"
db_password="${MINDOC_DB_PASSWORD||}"
```
> 我这里使用的MySQL，因此将adapter改为mysql即可。下面的一些信息改成MySQL实际的配置信息即可。其他的配置信息就根据自己实际需要来做修改即可。

### 配置MySQL信息
```shell
# 创建mysql数据库
create database mindoc;
# 创建mysql用户
CREATE USER 'username'@'host' IDENTIFIED BY 'password';
# 授权给新建的mysql用户
grant all ON databasename.tablename TO 'username'@'host';
# 刷新全新啊
flush privileges;
```
> 注意事项

1. username：创建的MySQL用户名称。
2. host: MySQL用户授权地址，如果新建的MySQL用户只能本地登录，则使用127.0.0.1即可。如果需要允许该用户远程登录，则使用使用%。
3. password：MySQL用户的密码
4. databasename：新建MySQL用户授权对应的数据库。这里直接写新建的数据库mindoc即可。如果是授权所有数据库，则使用*表示。
5. tablename：授权数据库对应的数据表，如果只是授权新建MySQL操作部分表，直接写表名，一般都是授权所有表，直接写*即可。
6. grant all：这里指的给新建的MySQL用户，授予所有的权限。如果只是部分授权，例如增删改查，则使用insert,delete,update,select 代替，每一个权限之间用","隔开。

## 安装启动

初始化数据
进入mindoc解压的根目录，会发现有一个mindoc_linux_amd64文件，该文件为启动文件。执行下面的命令:
```
./mindoc_linux_amd64 install
```
> 该命令的作用时初始化一些数据到MySQL中。就类似PHP很多软件，通过界面来安装应用。

启动服务
```shell
./mindoc_linux_amd64
```

问题出现，在启动的过程中可能会出现如下的情况:
```
OperationalError: (_mysql_exceptions.OperationalError) (1055, "Expression #1 of
SELECT list is not in GROUP BY clause and contains nonaggregated column 
'db.table.create_time' which is not functionally dependent on columns in 
GROUP BY clause; this is incompatible with sql_mode=only_full_group_by") [SQL: 
u'SELECT table.create_time AS table_create_time, table.count AS 
table_count \nFROM table \nWHERE table.cluster_id = %s AND 
table.create_time > %s GROUP BY table.count'] [parameters: (1, 
datetime.datetime(2000, 1, 1, 0, 0))]
```
> 在实际的过程中错误的语言不是这样的，在安装时忘记记录了，不过大致的错误信息是这样的。只要关注下面这一段错误码，就行了。1055, "Expression #1 of
SELECT list is not in GROUP BY clause and contains nonaggregated column 
'db.table.create_time' which is not functionally dependent on columns in 
GROUP BY clause; this is incompatible with sql_mode=only_full_group_by"

解决问题。
```shell
mysql> show variables like 'sql_mode';
+---------------+-------------------------------------------------------------------------------------------------------------------------------------------+
| Variable_name | Value                                                                                                                                     |
+---------------+-------------------------------------------------------------------------------------------------------------------------------------------+
| sql_mode      | ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION |
+---------------+-------------------------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```
执行下面的命令，将该sql_mode的配置信息改外其他的模式：
```shell
set @@sql_mode=STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
```
此时，再来查询该环境变量的值，就是如下的值了:
```shell
mysql> show variables like 'sql_mode';
+---------------+------------------------------------------------------------------------------------------------------------------------+
| Variable_name | Value                                                                                                                  |
+---------------+------------------------------------------------------------------------------------------------------------------------+
| sql_mode      | STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION |
+---------------+------------------------------------------------------------------------------------------------------------------------+
1 row in set (0.00 sec)
```
> 记得修改配置之后，一定的重启MySQL服务，否则是不会生效的。

重启mindoc服务。
```shell
./mindoc_linux_amd64
```
> 如果需要将该服务以守护进程的方式启动，则在后面加一个 &即可。

## 使用mindoc

上面正常将服务启动并安装，接下来使用下面的链接即可访问:
```shell
# ip为你服务器ip地址，如果是本地则使用localhost，或者127.0.0.1
ip:8181
```
下面是几张页面的截图，由于功能很好上手使用，这里就不单独介绍如何使用功能了。
![](https://oscimg.oschina.net/oscnet/up-d6c4ecaa39663a308718dea4daefbdf1757.png)
![](https://oscimg.oschina.net/oscnet/up-bf23e6a3a1009c0b11e6a0b2a843be95aaf.png)
![](https://oscimg.oschina.net/oscnet/up-bd03140cffa58e93bed881dbc0c43fc5538.png)