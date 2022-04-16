> 专注于PHP、MySQL、Linux和前端开发，感兴趣的感谢点个关注哟！！！文章整理在[GitHub](https://github.com/bruceqiq/code_study),[Gitee](https://gitee.com/bruce_qiq/code_study)主要包含的技术有PHP、Redis、MySQL、JavaScript、HTML&CSS、Linux、Java、Golang、Linux和工具资源等相关理论知识、面试题和实战内容。

[TOC]

## 说明

本文总结MongoDB操作集合的常用语法总结。

## 集合操作

在MongoDB中的集合，就是MySQL中的数据表。

### 查询集合

```mongodb
show collections
```
> 查询当前数据库下的所有集合

### 创建集合

```mongodb
# 手动创建集合
db.createCollection("collectionName")

# 自动创建集合，在做数据插入操作时，集合不存在时会自动创建
db.collectionName.insertOne()/insertMany()
```
> collectionName为集合的名称, 后面命令涉及到的collectionName也是指的集合名称。

### 删除集合

```mongodb
db.collectionName.drop()
```

### 查询集合

```mongodb
# 查询集合的所有数据
db.collectionName.find({})
# 查询一条集合数据
db.collectionName.findOne({})
```

### 插入操作

```mongodb
# 插入单条数据
db.collectionName.insertOne({"field1": "value1", "field2": "value2",
 ..., "fieldn": "valuen"})

# 插入多条数据
db.collectionName.insertMany([
  {"field1": "value1", "field2": "value2",..., "fieldn": "valuen"},
  {"field1": "value1", "field2": "value2",..., "fieldn": "valuen"},
  {"field1": "value1", "field2": "value2",..., "fieldn": "valuen"},
  {"field1": "value1", "field2": "value2",..., "fieldn": "valuen"},
  {"field1": "value1", "field2": "value2",..., "fieldn": "valuen"},
  ...
])
```

### 修改操作

```mongodb
# 更新单条数据(只会修改一条数据)
db.collectionName.updateOne({"条件字段": "条件字段值"}, {"修改字段": "修改字段值"})

# 批量更新(符合条件的数据都将被修改)
db.collectionName,updateMany({"条件字段": "条件字段值"}, {"修改字段": "修改字段值"})

# 文档替换
db.collectionName.replaceOne({"条件字段", "条件字段值"}, {"修改字段": "修改字段值"})
```

### 删除操作

```mongodb
# 删除单条数据
db.collectionName.dropOne({})

# 批量删除数据
db.collectionName.dropMany({})
```

### 查询操作

```mongodb
# 查询单条数据
db.collectionName.findOne({})

# 查询多条数据
db.collectionName.find({})
```

## 示例演示

我们以一张用户表为例，用户表里面有id、nickname、sex、mobile、address、age、score、

### 创建集合

1. 自动创建

2. 手动创建

### 插入数据

1. 单条插入数据

2. 批量插入数据

### 查询数据

1. 单条查询数据

2. 多条数据查询

3. 等值条件查询

3. in条件查询

4. 比较条件查询

5. like条件查询

### 删除数据

1. 单条删除

2. 全部删除

3. 带条件删除

### 修改数据

1. 单条修改

2. 批量修改

3. 带条件修改