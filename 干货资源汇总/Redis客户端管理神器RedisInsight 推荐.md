[TOC]

> 专注于PHP、MySQL、Linux和前端开发，感兴趣的感谢点个关注哟！！！文章整理在[GitHub](https://github.com/bruceqiq/code_study),主要包含的技术有PHP、Redis、MySQL、JavaScript、HTML&CSS、Linux、Java、Golang、Linux和工具资源等相关理论知识、面试题和实战内容。

Redis作为一个`高性能`、`内存性`的`nosql数据库`，已经成为日常开发必不可少的技术。对于一个开发人员来讲，不仅仅会使用Redis那么简单，也要学会如何管理、监控Redis服务。

对于Redis的管理，我们可以使用使用命令工具来进行管理，如果只是单机服务还比较好，如果存在集群或者多台服务的管理，通过命令行来操作就显得尤为繁琐。这时候我们就可以借助于一些可视化的界面端工具来进行管理。针对这样的工具也非常多。

![Snipaste_2021-07-25_23-51-48](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-25/1627228321982-Snipaste_2021-07-25_23-51-48.png)

上面的这些工具，有的是收费的，有只能查询、修改一些key的功能，功能比较单一。一般我都是使用`redisdesktopmanager`进行日常的Redis管理工具。它是一个跨平台的Redis管理工具，支持的功能也比较多。

官网地址：https://rdm.dev/

![Snipaste_2021-07-25_23-52-53](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-25/1627228525477-Snipaste_2021-07-25_23-52-53.png)

最近在使用Redis的布隆过滤器的时候，无意中发现一款Redis神器。真的是超级好用，功能也非常强大。对比了一下，应该是世界上最好用的Redis管理神器。`redis-insight`

官网地址：https://redislabs.com/redis-enterprise/redis-insight/#insight-form
![](https://img.soogif.com/XnNqohnplGxKXlor3enFe5JZlow7hxQN.gif?scope=mdnice)

> 这里就不介绍上面推荐的几种工具怎么使用了，反正是傻瓜式的操作。一用就会的那种。

## redislabs介绍

先说说redislabs这个网站是干啥吧。Redis Labs是一家`云数据库服务供应商`，致力于为Redis、 Memcached等流行的NoSQL 开源数据库提供云托管服务，推出产品全面管理的Redis Cloud服务和创建、管理Redis 数据库的Redis Labs Enterprise Cluster等。 公司地址：北美洲-美国 公司规模：10人以下 融资阶段：D轮 官网：https://redislabs.com/

登录到redislabs官网，可以看到支持很多的Redis插件。例如RedisJson、Redisgraph、Redisfliter等等。具体的介绍：https://redislabs.com/redis-enterprise/modules/

## 步入正题

上面感觉说了一堆的杂七杂八的介绍，下面就开始不如正题，引出今天的主人翁。
![](https://img.soogif.com/P1PAMcL3paRQBTcipjtgestB9RuLMu1Q.gif?scope=mdnice)
![](https://img.soogif.com/nfBw9bSV8WIheWwuojWQrwp6jQTvYUDB.gif?scope=mdnice)

RedisInsight旨在于简化Redis应用程序开发。

![Snipaste_2021-07-26_00-09-19](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229369724-Snipaste_2021-07-26_00-09-19.png)

## 功能介绍

### 可视化并与Redis数据库交互

扫描现有密钥，添加新密钥，执行CRUD或批量操作。以漂亮的打印JSON对象格式显示对象，并支持友好的键盘导航。

![redis-insight-hero-screenshot3](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229435570-redis-insight-hero-screenshot3.png)

### 对Redis模块的内置支持

查询、可视化和交互式操作图形、流和时间序列数据。使用多行查询编辑器生成查询、浏览结果、优化和快速迭代。支持RedisJSON、RediSearch、RedisGraph、Streams、RedisTimeSeries和RedisGears。

![redisinsights-redisgraph](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229528825-redisinsights-redisgraph.png)

### Redis的内存分析

离线分析内存使用情况，通过密钥模式、密钥过期和高级搜索来确定内存问题，而不影响Redis的性能。利用建议减少内存使用。

![redisinsights-analyze-overview](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229629957-redisinsights-analyze-overview.png)

### Trace Redis命令

识别顶键、键模式和命令。按群集所有节点上的客户端IP地址、密钥或命令进行筛选。有效地调试Lua脚本。

![keyspace_summary](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229683562-keyspace_summary.png)

### 直观的CLI

当一个GUI还不够时，我们的命令行界面利用redis cli提供语法高亮显示和自动完成，并使用集成的帮助来提供直观的即时帮助。

![redisinsights-CLI](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229736302-redisinsights-CLI.png)

### 管理Redis

深入了解实时性能指标，检查慢命令，并通过接口直接管理Redis配置。

![redisinsights-overview](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229779737-redisinsights-overview.png)

## 安装与使用

安装的方式也支持多样，支持Mac端、Linux端、Windows端，以及docker搭建。后面演示我也适用docker安装的，同时也推荐大家使用docker来进行操作。

在使用之前需要你进行填写一个个人的基本信息。

![Snipaste_2021-07-25_23-38-43](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627229903982-Snipaste_2021-07-25_23-38-43.png)

具体的操作文档：https://docs.redislabs.com/latest/ri/

1. docker安装

首先在我们本地新建一个目录，我这里的目录为`redisinsight`。主要是将docker内部的数据给和本地磁盘做一个挂载。
```shell
// 创建挂载目录
mkdir redisinsight
// 拉取镜像并启动容器
docker run -v -d redisinsight:/db -p 8001:8001 redislabs/redisinsight:latest
```
```shell
 kert@192  ~  docker ps
CONTAINER ID   IMAGE                           COMMAND                  CREATED          STATUS             PORTS                                                                     NAMES
c2772f255565   redislabs/redisinsight:latest   "bash ./docker-entry…"   37 minutes ago   Up 37 minutes      0.0.0.0:8001->8001/tcp, :::8001->8001/tcp                                 xenodochial_taussig
8aeeca3792c4   php_dnmp_nginx                  "/docker-entrypoint.…"   4 weeks ago      Up About an hour   0.0.0.0:80->80/tcp, :::80->80/tcp                                         nginx
057860a8a2ff   php_dnmp_php                    "docker-php-entrypoi…"   4 weeks ago      Up 34 hours        9000/tcp, 0.0.0.0:9501-9504->9501-9504/tcp, :::9501-9504->9501-9504/tcp   php
```
> 看到上面的8001端口和status就表明我们已经成功的安装好了。


2. 访问并初始化配置

搭建好服务之后，我们通过浏览器器访问`http://127.0.0.1:8001/`。第一次会出现下面的界面，我们全部勾选上即可。

![Snipaste_2021-07-26_00-22-00](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627230198623-Snipaste_2021-07-26_00-22-00.png)

接下来，我们如果使用本地的Redis服务直接选择左侧按钮即可，如果你本地没有Redis，选择右侧按钮进行一步一步操作即可，也是非常简单的。

接下来，设置我们Redis的链接信息。我这里默认是本地，主机地址就写`127.0.0.1`,端口号就写`6379`即可，Name字段则是给连接创建一个名字，可以随意填写。

![Snipaste_2021-07-26_00-24-47](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627230412062-Snipaste_2021-07-26_00-24-47.png)

![Snipaste_2021-07-26_00-26-41](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627230420018-Snipaste_2021-07-26_00-26-41.png)

3. 效果初览

经过上面的步骤，我们已经创建好所有的基础工作，接下来直接使用即可。首先进入的是总览页面。

![Snipaste_2021-07-26_00-28-47](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627230537544-Snipaste_2021-07-26_00-28-47.png)

> 很直观的显示了，客户端连接数、内存使用总量、总key数量、命中率、服务启动时间等信息。

4. 浏览左右的key

点击左侧的brower可以浏览数据库中存在的key。

![Snipaste_2021-07-26_00-31-19](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627230687101-Snipaste_2021-07-26_00-31-19.png)

5. 使用cli工具。

点击左侧的CLI，我们可以直接使用Redis的命令。不仅仅执行Redis命令，而且还将命令的一些搜索模式给显示出来。

![Snipaste_2021-07-26_00-32-23](https://gitee.com/bruce_qiq/picture/raw/master/2021-7-26/1627230751278-Snipaste_2021-07-26_00-32-23.png)

后面的功能就不一一介绍了，大家直接界面的操作进行使用就可以了。

![](https://img.soogif.com/94kTxftf5beFshfxV295IXlbKjewbJtz.gif?scope=mdnice)












