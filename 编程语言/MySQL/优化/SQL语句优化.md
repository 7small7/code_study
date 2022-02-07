## SQL语句优化
### 插入优化
1. 合并多条SQL语句

```php
insert into `user`
(`nickname`, `age`, `score`, `activity_id`, `user_id`) value
("张三", 12, 12.09, 1, 2);
insert into `user`
(`nickname`, `age`, `score`, `activity_id`, `user_id`) value
("张三", 12, 12.09, 2, 2);
```
合并为下面的语句，可以减少每一次内部解析的过程。
```php
insert into `user`
(`nickname`, `age`, `score`, `activity_id`, `user_id`) values
("张三", 12, 12.09, 1, 2),
("张三", 12, 12.09, 2, 2);
```
2. 主键按照顺序插入。
如果在执行批量的insert操作，并且insert带有id字段。id的顺序应该是有序的，这样减少插入时维护索引顺序带来的时间消耗。
1. 开启事务提交。
将数据分段进行提交。
1. 关闭唯一索引检测。在插入之后，打开唯一索引检测，针对重复的数据使用手动检测的方式进行处理。

### 查询优化
1. 选择合适的查询字段进行查询。

```php
select * from `user`;
```
> 这里应该根据实际要查询的字段，指定查询的字段。而不是使用*查询全部字段。

1. 减少大字段的查询。

假设content就是一个`text`的字段类型，下面的select语句应该尽量减少该字段的查询，尽可能的单独去查询该字段。
```php
select `id`, `title`, `content` from article;
```

2. order by查询。
如果order by是通过索引查询，需要按照`索引创建时的顺序进行排序，并且排序的排序规则需要保持一致`。例如存在这样的一个复合索引(user_id, activity_id)。否则会出现filesort的情况，因为索引是按照循序排列的，此时走filesort，会发生随机I/O操作。

a. 未按照索引创建时的顺序查询，排序规则保持一致。

```php
explain select 
id ,
activity_id, 
user_id 
from user_1 order by activity_id desc,user_id desc\G;
***************************[ 1. row ]***************************
id            | 1
select_type   | SIMPLE
table         | user_1
partitions    | <null>
type          | ALL
possible_keys | <null>
key           | <null>
key_len       | <null>
ref           | <null>
rows          | 4
filtered      | 100.0
Extra         | Using index; Using filesort
```
b. 未按照索引创建时的顺序查询，排序规则未保持一致。

```php
explain select 
id ,
activity_id, 
user_id 
from user_1 order by activity_id desc,user_id asc\G;
***************************[ 1. row ]***************************
id            | 1
select_type   | SIMPLE
table         | user_1
partitions    | <null>
type          | ALL
possible_keys | <null>
key           | <null>
key_len       | <null>
ref           | <null>
rows          | 4
filtered      | 100.0
Extra         | Using index; Using filesort
```
> 通过上面两种查询方式，可以看出没有按照索引创建的顺序进行排序，就会出现某一个索引走索引排序，但是另外一个索引列会走文件排序。在测试的过程中，如果按照索引创建的顺序进行排序，但是排序的规则不一致，也会出现Using filesort的情况。

c. 按照索引创建顺序进行排序，并且排序规则保持一致。

```php
explain select 
id ,
activity_id, 
user_id 
from user_1 order by user_id desc, activity_id desc\G;
***************************[ 1. row ]***************************
id            | 1
select_type   | SIMPLE
table         | user_1
partitions    | <null>
type          | index
possible_keys | <null>
key           | idx_u_a
key_len       | 8
ref           | <null>
rows          | 4
filtered      | 100.0
Extra         | Backward index scan; Using index
```
> 此时按照索引创建的顺序进行排序，并且排序规则保持一致。这时候就直接通过索引进行排序，而不会走filesort。在测试过程中将desc改为asc也是一样的结果，使用的using index排序，而不会走filesort。

1. group by查询。

使用group by查询，尽量在索引列上进行group查询并在后面添加一个order by null。

1. in查询。

in查询，如果在走了索引，一般也会使用using index查询。尽量使用join方式进行查询。

1. or 查询。

or查询，即使多列都存在索引，or后续的列的索引也会失效。尽量使用union的方式进行查询。

1. 使用分页查询。

```php
select * from `user_1` limit 100000, 10;
```
上面的语句，在数据量大情况下，查询也会特别耗时。应该使用下面的方式，该方式在做分页时减少了磁盘查找，先在索引中查询出对应的ID，走了聚集索引。
第一种方案是用wherein语句。
```php
select * from `user_1` where id in (select id form `user_1` limit 100000, 10);
```
第二种方案是用join查询。
```php
select a.* from `user_1` as a 
(select id from `user_1` limit 100000, 10) b
inner join a.id = b.id;
```
> 不过前面也提到in的性能，相对join来说低一些。因此，推荐使用join语句查询。