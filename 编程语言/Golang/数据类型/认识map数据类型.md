[TOC]
## 语法定义

map是一种键值类型的数组，也可以称为集合。

## 定义

```go
var mapName map[keyType]valueType
```
1. mapName 为 map 的变量名。
2. keyType 为键类型。
3. valueType 是键对应的值类型。

> [keyType] 和 valueType 之间允许有空格。

## 创建方式

### make声明

```go
// 1.
// 不设置容量的情况下,会根据实际的内容进行扩升,每次扩容的长度是1
map1 := make(map[string]string)
fmt.Println(map1)
```

### 变量声明

```go
// 2.使用变量声明
// 此时还未分配内存
var map2 map[string]string
// 必须使用make()函数    
map2 = make(map[string]string) 
map2["name"] = "张三"
map2["age"] = "12"
map2["sex1"] = "女"
map2["sex"] = "男"
map2["id_card"] = "123423423452345234523452345"
fmt.Println("map2", map2)
```

### 使用变量(默认初始值)

```go
// 3. var name map[keyType]valueType = map[keyType]valueType{"键":"值","键":"值"}
map4 := map[string]string{"id": "1", "sex": "男", "name": "张三", "age": "12"}
fmt.Println(map4)

map3 := make(map[string]string, 3) // 容量会自动的增加，每次进行加1
fmt.Println("切片的长度是:", len(map3))
map3["name"] = "lisa"
map3["name1"] = "lisa"
map3["name2"] = "lisa"
map3["name3"] = "lisa"
fmt.Println(map3)
fmt.Println("切片的长度是:", len(map3))
```

## 注意事项

1.map键类型，可以是数字、字符串、boolean、指针、channel、接口、结构体、数组。但一定不能是slice，map和function。

2.map值类型，和键类型一致。

3.map是无序的，也就是说，在输出map时，输出的循序不是按照赋值的顺序来的，而是按照a-z字母的顺序进行排序。

4.map的键是不能重复的，如果重复会覆盖前面的键；map的值是可以重复的。

5.使用变量方式声明的map，是一个空的map，是没有分配内存空间的，必须使用make()函数。

6.使用len()函数可以计算出map的长度。

## 常见操作

### 循环
```go

	map5 := map[string]string{
		"name":    "张三",
		"age":     "12",
		"sex":     "男",
		"address": "上海市宝山区", // 如果{}里面的值换行，最后一个","是不能省略的。
	}
	for k, v := range map5 {
		fmt.Println(k, "=>", v)
	}
```

### 追加

```go
	map5["id_card"] = "12311313131331"
	fmt.Println("追加后的map内容是", map5)
```

### 修改

```go
map5["sex"] = "女"
fmt.Println("修改map的值", map5)
```

### 删除

```go
delete(map, "删除的键")
delete(map5, "id_card")
fmt.Println("删除后的map内容是", map5)
```

### 排序

map是不能直接进行排序，一般的方式是根据排序的条件，将键/值追加到slice，然后对slice进行排序。
```go
// slice1 := make([]string, 10),这种方式因为在定义时，
// 默认初始化了一个为空的切片，因此在下面使用for循环时，会在切片的尾部今夕追加。
var slice1 []string
fmt.Println(slice1, len(slice1), cap(slice1))
for k, _ := range map5 {
	slice1 = append(slice1, k)
}
fmt.Println(slice1, len(slice1), cap(slice1))
sort.Strings(slice1)
fmt.Println(slice1)
```

### 清空

```go
map5 = make(map[string]string)
fmt.Println(map5)
```

### map切片

```go
slice2 := make([]map[string]string, 1)
//  手动追加map到切片内容（使用该方式，当切片的长度超过定义时的长度，需要改动定义的切片长度）
slice2[0] = map[string]string{
	"name": "张三",
	"age":  "23",
}
// 动态追加（利用切片的动态扩容特点）
newMap := map[string]string{
	"name": "张三",
	"age":  "23",
}
slice2 = append(slice2, newMap)
fmt.Println(slice2)
	```