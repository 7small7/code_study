## 事务基础路线
![m_5d609cce116fe7465ba28dd3ebac1443_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725490029-m_5d609cce116fe7465ba28dd3ebac1443_r.png)


## 事务定义

事务就是一组 DML 语句的集合。事务保证了对数据库中数据的一致性操作。

## 存储引擎

在日常开发中，我们常用的存储引擎有 InnoDB 和 MyISAM 两种存储引擎。然而 MyISAM 是不支持事务操作的。

## 事务的提交方式

### 示例代码
```mysql
CREATE TABLE `user`  (
  `id` int(10) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT,
  `name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `age` int(11) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

mysql root@127.0.0.1:demo> desc user;
+-------+------------------+------+-----+---------+-------+
| Field | Type             | Null | Key | Default | Extra |
+-------+------------------+------+-----+---------+-------+
| id    | int(10) unsigned | NO   | PRI | <null>  |       |
| name  | varchar(10)      | NO   |     | <null>  |       |
| age   | int(11)          | YES  |     | 0       |       |
+-------+------------------+------+-----+---------+-------+
3 rows in set
Time: 0.012s

-- 插入语句
insert into `user`(`id`,`name`, `age`) values 
(1,'张三', 12),(2,'李四', 13),(3,'王五',14),(4,'赵六',15);
```

MySQL 中的事务默认采用的是自动提交，即一条 SQL 语句就是一个事务。默认提交方式可以通过参数 auto_commit 参数进行设置。MySQL 中的事务提交方式主要分为手动提交和自动提交方式。隐式提交属于自动提交的一种特殊方式。
![m_b74aaf1f08fe6ad3e6dd0813d36cf7d4_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725510409-m_b74aaf1f08fe6ad3e6dd0813d36cf7d4_r.png)

### 自动提交

下面的 SQL 语句，针对 user 表做一个 update 操作，当MySQL 执行这条 SQL 语句之后，数据就进行了持久化存储，数据就不能进行回滚了，除非我们手动的 update 回原来的数据。
```mysql
// 执行 update 语句之前表数据
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.009s
// 执行 update 语句
mysql root@127.0.0.1:demo> update `user` set name = '张三1' where id = 1;
Query OK, 1 row affected
Time: 0.004s
// 查询update 语句之后的表数据
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.010s
```
> 通过上面的示例代码，可以看出，在默认的情况下，我们执行一个 SQL 语句，就会自动的存储到表中。

### 手动提交

手动处理事务的提交方式，需要用到commit 关键字。下面罗列几个 MySQL 事务中需要涉及到的几个关键词。

begin: 开启一个事务,也可以使用 start transaction，两者都是等价的。

rollback: 回滚一个事务的操作 ，也可以使用 rollback work，两者都是等价的。

commit: 提交一个事务的操作， 也可以使用 commit work，两者都是等价的。

savepoint transactionId: 保存一个事务点。

release transactionId: 释放一个事务点。

rollback transactionId: 回滚到指定的事务点。

1.此时我们新建一个客户端A，开启一个事务操作。
```mysql
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行 update 语句
mysql root@127.0.0.1:demo> update `user` set name = '张三' where id = 1;
Query OK, 1 row affected
Time: 0.002s
// 查看当前事务修改结果
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.010s
```
2.此时新建一个客户端 B，开启一个事务操作。
```mysql
// 开启一个事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行查询语句
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.010s
```
> 通过上面的对比，发现在客户端 A 执行了一个 update 操作，事务内查询id=1 的 name 值为张三，然而在客户端 B 查询到 id=1 的 name 值为张三 1。这就是事务之间的隔离性(可重复读)。是因为客户端 A 还没执行 commit 操作。

### 隐式提交

这里的隐式提交指的是事务的嵌套，就是一个事务里面包含着另外的一个事务。当事务里面在开启事务时，第一个事务默认会执行 commit 操作。
1.此时我们新建一个客户端A，开启一个事务操作。
```mysql
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
mysql root@127.0.0.1:demo> update `user` set name = '张三1' where id = 1;
Query OK, 1 row affected
Time: 0.002s
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行 sql 查询语句
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.009s
```
> 此时在客户端 A 的事务内又开启了另外一个事务，执行 SQL 查询语句后，发现数据是修改后的值。

2.此时新建一个客户端 B，无需开启事务。
```mysql
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.009s
```
> 从理论上分析，由于客户端 A 的事务没有提交，此时我们查询到的数据不应该是事务内修改后的值，然而我们执行查询时，发现数据表的值是修改后的值，这证明嵌套事务会自动提交(此时的事务隔离级别是可重复读)。 如果你认为当前的事务隔离级别是未提交读，那你可以尝试关闭 MySQL 服务这极端的操作在来查询数据，你的到的结果也是如此。

### 提交点

所谓的提交点，就是指在一个事务中，做一个类似于数据快照一样的操作。在数据为提交之前，我们可以针对不同的数据快照进行回退。
代码演示:
```mysql
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 查看修改前的数据
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.010s
// 创建一个 savepoint
mysql root@127.0.0.1:demo> savepoint trans_1;
Query OK, 1 row affected
Time: 0.002s
// 执行修改操作
mysql root@127.0.0.1:demo> update `user` set name = '张三1' where id = 1;
Query OK, 1 row affected
Time: 0.002s
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.010s
// 回滚savepoint
mysql root@127.0.0.1:demo> rollback trans_1;
// 查询数据
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.010s
```
> 执行回滚快照之后，update 语句执行的修改操作被撤回了。

## 事务的四大特性

事务的四大特性分别指的是 ACID，只有具备这四大特性的数据库，才可具备事务的操作。这四大特性分别是原子性、持久性、隔离性和持久性。
![m_3a69f58f5631b0409070539742b63094_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725539105-m_3a69f58f5631b0409070539742b63094_r.png)


### 原子性

定义：原子性指的是事务操作具备原子操作，一个事务里面的 SQL 操作要么全部成功要么全部失败，不能存在一些 SQL 成功，一些 SQL 执行失败。
场景：银行转账，你给小明转账 100 块，此时给小明的账户增加 100 块是一个 SQL 语句，你账户上减少 100 块是一个 SQL 语句，原子性就是指的这两个 SQL 要全部成功。
示例: 
```mysql
update user_account set money = money - 100.00 where user_name = '自己';
update user_account set money = money + 100.00 where user_name = '小明';
```
![m_7a3328aae26bb8f5d1f283eda63d6979_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725552725-m_7a3328aae26bb8f5d1f283eda63d6979_r.png)

### 一致性

定义：一致性指的是事务操作前后必须满足业务约束。
场景：银行转账，你当前账户有 200 块，小明有 100 块。你给小明转账 100 块之后，你账户有 100 块，小明账户有 200 块。此时不管是转账失败还是成功，账户的总金额还是和转账之前的总金额保持一致的。
示例：
```mysql
// 计算事务操作之前的总和
select sum(money) as sum_money from user_account where user_name = '自己' and user_name = '小明';
|-----sum_money-----|
|        300.00            |
update user_account set money = money - 100.00 where user_name = '自己';
update user_account set money = money + 100.00 where user_name = '小明';
// 计算事务操作之后的总和
select sum(money) as sum_money from user_account where user_name = '自己' and user_name = '小明';
|-----sum_money-----|
|        300.00            |
```
![m_c1832dccc1d9fd64b16d6b3645e4902c_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725571182-m_c1832dccc1d9fd64b16d6b3645e4902c_r.png)

### 隔离性

定义：隔离性指的是多个事务之间是相互隔离的，事务之间是互不受影响的。但这种场景需要考虑事务的隔离级别。文章后面内容页会针对事务之间的隔离级别做特别的演示。
场景：银行转账，你当前账户存在 200 块钱，你此时给小明转账 100 块，同时你女朋友用你的账户在执行支付操作，这两种场景属于两个事务操作，两种操作是互不受影响的。
示例：
```mysql
// 转账小明操作
update user_account set money = money - 100.00 where user_name = '自己';
// 女朋友支付操作
update user_account set money = money - 100.00 where user_name = '自己';
```
![m_1b90a7f5777600d881fbe47a8e1fa17b_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725593662-m_1b90a7f5777600d881fbe47a8e1fa17b_r.png)


### 持久性
定义：持久性指的是事务一旦提交，就不能进行回滚(撤回)，永久的保存在磁盘中。
场景：银行转账，你当前给小明转账 100，点击了确认操作，此时你后悔转账了，此时是没法撤回的，因为这个操作已经被提交，也就表明小明账户上已经多 100 了。
```mysql
// 开启事务
begin;
update user_account set money = money - 100.00 where user_name = '自己';
update user_account set money = money + 100.00 where user_name = '小明';
// 提交操作
commit;
select money from user_account where user_name = '自己'; // 100.00
select money from user_account where user_name = '小明'; // 200.00
// 执行回滚
rollback
//此时金额是不发生改变
select money from user_account where user_name = '自己'; // 100.00
select money from user_account where user_name = '小明'; // 200.00
```

## 事务隔离级别

事务的隔离级别是针对两个或两个以上的事务之间是否相互是隔离的，这里的隔离其实就是指的事务对数据做了操作，在事务在commit 之前，其他的未提交事务是否可以读到相应修改的数据。大致逻辑如下图：
![m_b414510f123f598da1de713509cec751_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725611828-m_b414510f123f598da1de713509cec751_r.png)

1.当事务1 开启时，数据表中的 id=1 的数据，age=2；

2.事务1执行 update 操作，将 id=1的数据，age 设置为 1；

3.此时在事务 1 还未提交时，开启事务 2；

4.事务2 对数据表id=1 的数据进去 select；

5.由于事务 1 执行了 update 操作，当事务 2 去执行 select 操作时，是返回 age=1？还是 age=2？这就涉及到事务的隔离级别。

### 系统隔离级别
MySQ默认使用的可重复读隔离级别。如下可以查看 MySQL 事务隔离级别。
```mysql
mysql root@127.0.0.1:demo> show variables like '%iso%';
+-----------------------+-----------------+
| Variable_name         | Value           |
+-----------------------+-----------------+
| transaction_isolation | REPEATABLE-READ |
| tx_isolation          | REPEATABLE-READ |
+-----------------------+-----------------+
2 rows in set
Time: 0.011s
```
> transaction_isolation是tx_isolation 的别名，推荐使用transaction_isolation选项，因为在 MySQL8 开始，tx_isolation将被废弃。

### 未提交读

定义：所谓的未提交读，指的是当一个事务 A内对数据做了操作还没有执行commit ，另外一个事务B是可以读到事务A修改后的数据。如上图中，事务二去做查询时，返回的值就是 2。
设置隔离级别:
```mysql
mysql root@127.0.0.1:demo> set session transaction isolation level READ UNCOMMITTED;
Query OK, 0 rows affected
Time: 0.002s
mysql root@127.0.0.1:demo> show variables like '%iso%';
+-----------------------+------------------+
| Variable_name         | Value            |
+-----------------------+------------------+
| transaction_isolation | READ-UNCOMMITTED |
| tx_isolation          | READ-UNCOMMITTED |
+-----------------------+------------------+
2 rows in set
Time: 0.010s
```
代码演示：
**1.开始事务 A。**
```mysql
// 开启事务前查询数据
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.010s
// 设置当前会话的隔离级别
mysql root@127.0.0.1:demo> set session transaction isolation level READ UNCOMMITTED;
Query OK, 0 rows affected
Time: 0.002s
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行修改操作
mysql root@127.0.0.1:demo> update `user` set name = '张三' where id = 1;
Query OK, 1 row affected
Time: 0.002s
// 查询数据已被修改(但当前事务还未被提交)
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.011s
```
**2.开启事务 B。**
```mysql
// 设置当前会话的隔离级别
mysql root@127.0.0.1:demo> set session transaction isolation level READ UNCOMMITTED;
Query OK, 0 rows affected
Time: 0.002s
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行 SQL 查询
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.011s
```
> 通过上面的事务 A 和事务 B 可以看得出，在事务 A 还没提交时，事务 B 就能读取到事务A 修改后的数据。在演示代码时，需要注意的是，需要在事务 A 开启之后执行 update 操作之前，将事务 B 开启。

### 提交读

定义：所谓的提交读，指的是当一个事务 A内对数据做了操作并且提交了，另外一个事务B在还未提交时，是可以读到事务A修改后的数据。如上图中，事务二去做查询时，如果此时事务一已经执行了 commit，此时返回的结果仍是 2。
设置隔离级别:
```mysqll
mysql root@127.0.0.1:demo> set session transaction isolation level READ COMMITTED;
Query OK, 0 rows affected
Time: 0.002s
```
代码演示：
**1.开始事务 A。**
```mysql
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.011s
// 设置当前会话的事务隔离级别
mysql root@127.0.0.1:demo> set session transaction isolation level READ COMMITTED;
Query OK, 0 rows affected
Time: 0.002s
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行修改操作
mysql root@127.0.0.1:demo> update `user` set name = '张三' where id = 1;
Query OK, 1 row affected
Time: 0.002s
// 提交事务
mysql root@127.0.0.1:demo> commit;
Query OK, 0 rows affected
Time: 0.003s
// 查看事务执行结果
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.011s
```
**2.开启事务 B。**
```mysql
// 设置当前会话的事务隔离级别
mysql root@127.0.0.1:demo> set session transaction isolation level READ COMMITTED;
Query OK, 0 rows affected
Time: 0.002s
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 查看数据(此时事务 A 一定要处于未执行  commit 语句)
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.011s
// 再次查看数据(此时事务 A 一定要处于执行 commit 语句)
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.011s
```
> 通过上面的示例代码，可以看出，当事务 A 提交之后，事务 B 内就可以读到修改后的数据。记住在事务 B 内执行 select 语句时，事务 A 的提交状态。

### 可重复读

定义：所谓的可重复读，指的是当一个事务 A内对数据做了操作并且提交了，另外一个事务B在还未提交时，读到的数据永远是事务B开始时的状态,是不会读取到事务 A 提交后的数据。如上图中，事务二去做查询时，如果此时事务一已经执行了 commit，此时返回的结果是 1。
设置隔离级别:
```mysql
mysql root@127.0.0.1:demo> set session transaction isolation level REPEATABLE-READ;
Query OK, 0 rows affected
Time: 0.002s
```
**1.开启事务 A。**
```mysql
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.011s
// 设置当前会话的事务隔离级别
mysql root@127.0.0.1:demo> set session transaction isolation level REPEATABLE-READ;
Query OK, 0 rows affected
Time: 0.002s
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行修改操作
mysql root@127.0.0.1:demo> update `user` set name = '张三' where id = 1;
Query OK, 1 row affected
Time: 0.002s
// 提交事务
mysql root@127.0.0.1:demo> commit;
Query OK, 0 rows affected
Time: 0.003s
// 查看事务执行结果
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.011s
```
**2.开启事务 B。**
```mysql
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
+----+-------+-----+
4 rows in set
Time: 0.011s
// 设置当前会话的事务隔离级别
mysql root@127.0.0.1:demo> set session transaction isolation level REPEATABLE-READ;
Query OK, 0 rows affected
Time: 0.002s
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行查询操作
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.011s
// 提交事务
mysql root@127.0.0.1:demo> commit;
Query OK, 0 rows affected
Time: 0.003s
// 查看事务执行结果
mysql root@127.0.0.1:demo> select * from user;
+----+------+-----+
| id | name | age |
+----+------+-----+
| 1  | 张三 | 12  |
| 2  | 李四 | 13  |
| 3  | 王五 | 14  |
| 4  | 赵六 | 15  |
+----+------+-----+
4 rows in set
Time: 0.011s
```
> 通过上面的示例代码，可以看出，当事务 A 提交之后，事务 B 内是不可以读到修改后的数据，当事务 B 提交之后(事务 B 没错 DML 语句操作)，但是查询到数据和事务 A 提交后的结果一致。

### 串行
定义：所谓的串行，指的是当一个事务 A内对数据做了操作并且处于未提交状态，另外一个事务B在去操作数据时，是不会被立马执行的，而是需要等到事务 Acommit 之后，才可以进行数据操作。如上图中，事务二去做查询时，如果此时事务一处于未提交状态，事务二会一直处于等待状态，直到事务 A commit 之后返回结果值 2。
设置隔离级别:
```mysql
set session transaction isolation level serializable;
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
```
代码演示：
**1.开始事务 A**
```mysql
// 设置当前会话的隔离级别
mysql root@127.0.0.1:demo>set session transaction isolation level serializable;
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行修改
mysql root@127.0.0.1:demo> update `user` set name = '张三1' where id = 1;
Query OK, 1 row affected
Time: 0.002s
// 提交事务
mysql root@127.0.0.1:demo> commit;
Query OK, 0 rows affected
Time: 0.004s
// 检测修改
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
| 5  | 小QI  | 16  |
+----+-------+-----+
5 rows in set
Time: 0.009s
```
**2.开始事务 B**
```mysql
// 设置当前会话的隔离级别
mysql root@127.0.0.1:demo>set session transaction isolation level serializable;
// 开启事务
mysql root@127.0.0.1:demo> begin;
Query OK, 0 rows affected
Time: 0.002s
// 执行查询(此时事务 A 一定要还未执行 commit 命令)
mysql root@127.0.0.1:demo> select * from user;
+----+-------+-----+
| id | name  | age |
+----+-------+-----+
| 1  | 张三1 | 12  |
| 2  | 李四  | 13  |
| 3  | 王五  | 14  |
| 4  | 赵六  | 15  |
| 5  | 小QI  | 16  |
+----+-------+-----+
5 rows in set
Time: 14.159s(查询消耗时间，是由于 sessionA处于阻塞过程中)
```
> 此时一定要特别注重事务 B 在执行 select 语句所用的时间，这里为什么是 14s 多，就是因为在这段时间中，事务 A 还未处于commit 状态，因此事务 B 在执行查询时，一直处于等待状态，直到事务 A 进行了 commit 操作。

### 隔离级别的几种现象

| 事务隔离级别 | 脏读 | 不可重复度 | 幻读 | 加锁度 |
| :---: | :---: | :---: | :---: | :---: |
| 未提交读 | √ | √ | √ | × |
| 提交读 | × | √ | √ | × |
| 可重复读 | × | × | √ | × |
| 串行 | × | × | × | √ |

脏读：指的是事务处于未提交状态，对数据做了操作，其他的事务是可以读取到当前事务所做的操作。
![m_eae303d7c16a2c08f4e677ed9669b6fe_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725656060-m_eae303d7c16a2c08f4e677ed9669b6fe_r.png)

不可重复度：一个事务 A对数据做了操作并进行了提交。其他事务对于事务A在提交前和提交后读到的数据不一致。侧重(update/delete)语句。
![m_d90d7261c949b2246ea45306e99d1f7f_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725673393-m_d90d7261c949b2246ea45306e99d1f7f_r.png)

幻读：一个事务 A对数据做了操作并进行了提交。其他事务对于事务A在提交前和提交后读到的数据不一致。侧重(insert)语句。
![m_8e8d6bbb66be3ace59b55a900484f2e9_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725709026-m_8e8d6bbb66be3ace59b55a900484f2e9_r.png)

可重复读：事务 B 在提交前，不管事务 A 是否提交，读到的数据永远是事务 B 开始时的状态。
![m_ae96c565a0fabd92412f56ef2d6ff057_r](https://gitee.com/bruce_qiq/picture/raw/master/2021-3-26/1616725688209-m_ae96c565a0fabd92412f56ef2d6ff057_r.png)


#### 隔离级别优缺点分析

a.通过上面几种情况的分析，可重复读的隔离级别是最优的选择。真正做到了事务之间是相互隔离，避免了脏读、不可重复读和幻读的问题。

b.串行是最大程度上保证了数据的一致性，但是属于枷锁处理，并发性不高。

c.未提交读和不可重复度，虽然存在数据不一致的情况，但效率更高，适合用在对数据一致性要求不高的场景中。