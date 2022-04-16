[TOC]

## 前言

PHP uniqid()函数可用于生成不重复的唯一标识符，该函数基于微秒级当前时间戳。在高并发或者间隔时长极短（如循环代码）的情况下，会出现大量重复数据。即使使用了第二个参数，也会重复，最好的方案是结合 md5 函数来生成唯一 ID。

## 使用函数

```php
string uniqid ([ string $prefix = "" [, bool $more_entropy = false ]] )
```

获取一个带前缀、基于当前时间微秒数的唯一 ID。
prefix
有用的参数。例如：如果在多台主机上可能在同一微秒生成唯一 ID。
prefix 为空，则返回的字符串长度为 13。more_entropy 为 TRUE，则返回的字符串长度为 23。
more_entropy
如果设置为 TRUE，uniqid() 会在返回的字符串结尾增加额外的煽（使用 combined linear congruential generator）。 使得唯一 ID 更具唯一性。

## PHP uniqid() 生成不重复唯一标识方法一

这种方法会产生大量的重复数据，运行如下 PHP 代码会数组索引是产生的唯一标识，对应的元素值是该唯一标识重复的次数。

```php
<?php
$units = array();
for($i=0;$i<1000000;$i++){
    $units[] = uniqid();
}
$values  = array_count_values($units);
$duplicates = [];
foreach($values as $k=>$v){
    if($v>1){
        $duplicates[$k]=$v;
    }
}
echo '<pre>';
print_r($duplicates);
echo '</pre>';
?>
```

## PHP uniqid() 生成不重复唯一标识方法二

这种方法生成的唯一标识重复量明显减少。

```php
<?php
$units = array();
for($i=0;$i<1000000;$i++){
    $units[] = uniqid('',true);
}
$values  = array_count_values($units);
$duplicates = [];
foreach($values as $k=>$v){
    if($v>1){
        $duplicates[$k]=$v;
    }
}
echo '<pre>';
print_r($duplicates);
echo '</pre>';
?>
```

## PHP uniqid() 生成不重复唯一标识方法三

这种方法生成的唯一标识中没有重复。

```php
<?php
$units = array();
for($i=0;$i<1000000;$i++){
    $units[]=md5(uniqid(md5(microtime(true)),true));
}
$values  = array_count_values($units);
$duplicates = [];
foreach($values as $k=>$v){
    if($v>1){
        $duplicates[$k]=$v;
    }
}
echo '<pre>';
print_r($duplicates);
echo '</pre>';
?>
```

## PHP uniqid() 生成不重复唯一标识方法四

使用 session_create_id()函数生成唯一标识符，经过实际测试发现，即使循环调用 session_create_id()一亿次，都没有出现过重复。
php session_create_id()是 php 7.1 新增的函数，用来生成 session id，低版本无法使用。