[TOC]

## vim命令如何复制当前文件全部内容

```shell
按esc然后输入ggyG
```

## 如何将一个文件复制到多个目录下面

```shell
echo 目录一 目录二 目录三... | xargs -n 1 cp -v 文件名称
```

## 如何在同一目录创建多个目录

```shell
mkdir 目录一 目录二 目录三
```

## 如何递归创建目录

```shell
mkdir 一级目录/二级目录/三级目录/四级目录/....
```

## 统计Apache或者Nginx日志文件的访问信息

```shell
# 统计Nginx访问日志中IP前十的数据。
awk '{print $1}' /var/log/nginx/access.log | sort | uniq -c | sort -nr -k1 | head -n 10
# 根据访问IP统计UV
awk '{print $1}' access.log|sort | uniq -c |wc -l
# 统计访问URL统计PV
awk '{print $7}' access.log|wc -l
# 查询访问最频繁的URL
awk '{print $7}' access.log|sort | uniq -c |sort -nk 1 -r|more
# 查询访问最频繁的IP
awk '{print $1}' access.log|sort | uniq -c |sort -nk 1 -r|more
# 根据时间段统计查看日志
cat access.log| sed -n '/14\/Mar\/2015:21/,/14\/Mar\/2015:22/p'|more
awk '{ print $1}'：取数据的低1域（第1列）
sort：对IP部分进行排序。
uniq -c：打印每一重复行出现的次数。（并去掉重复行）
sort -nr -k1：按照重复行出现的次序倒序排列,-k1以第一列为标准排序。
head -n 10：取排在前5位的IP 。
```