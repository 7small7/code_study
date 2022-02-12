> 专注于PHP、MySQL、Linux和前端开发，感兴趣的感谢点个关注哟！！！文章整理在[GitHub](https://github.com/bruceqiq/code_study),[Gitee](https://gitee.com/bruce_qiq/code_study)主要包含的技术有PHP、Redis、MySQL、JavaScript、HTML&CSS、Linux、Java、Golang、Linux和工具资源等相关理论知识、面试题和实战内容。

今天咱们来学习一下golang中的map数据类型，单纯的总结一下基本语法和使用场景，也不具体深入底层。map类型是什么呢？做过PHP的，对于数组这种数据类型是一点也不陌生了。PHP中的数组分为`索引数组`和`关联数组`。例如下面的代码：
```php
// 索引数组【数组的key是一个数字， 从0，1，2开始递增】
$array = [1, '张三', 12];
// 关联数组【数组的key是一个字符串，可以自定义key的名称】
$array = ['id' => 1, 'name' => '张三', 'age' => 12];
```
在golang中，map是一种特殊的数据结构，是一种key对应一个value类型的结构。这种结构可以被称为关联数组和字典。
![Snipaste_2021-11-20_22-41-35](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-20/1637419308059-Snipaste_2021-11-20_22-41-35.png)
在golang中也有切片和数组这样的数据类型，来存储一组数据。
![Snipaste_2021-11-20_22-47-40](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-20/1637419688255-Snipaste_2021-11-20_22-47-40.png)

1. 数组就好比PHP中的一维数组，并且长度是固定的，其中的值类型在定义数组的时候就确定好了。

2. 切片是一种特殊的数组类型。长度是固定的。
```golang
// 数组，类型是int，长度是4。
array := [4]int{1, 2, 3, 4}
// 切片，类型是int，长度不固定。
slice := []int{1, 2, 3, 4}
```
有数组和切片可以存储一组数据，那为什么还有map这样的类型结构呢？map类型具体是啥样的呢？

## 案例

假设我们现在有这样的一个需求，要用golang中的一种数据类型来存储多个用户的数据，这些数据分别用户的ID，name，age，sex...等等字段。我们改用什么数据类型呢？

1. 在PHP中我们可以直接下面的方式定义，操作也是非常简单。

```php
$userInfo = [
  ['id' => 1, 'name' => '张三', 'age' => 12, 'sex' => '男'],  
  ['id' => 2, 'name' => '赵六', 'age' => 22, 'sex' => '男'],
  ['id' => 3, 'name' => '李四', 'age' => 34, 'sex' => '女'],
  ['id' => 4, 'name' => '王麻子', 'age' => 56, 'sex' => '男']
];
```

2. 那如何在golang中实现呢，假设我们用数组和切片实现一下试试。

```golang
// 1. 用数组实现
$user1 := [4]string{"1", "张三", "12", "男"}
$user2 := [4]string{"2", "赵六", "12", "男"}
$user3 := [4]string{"3", "李四", "12", "女"}
$user4 := [4]string{"4", "王麻子", "12", "男"}

// 2. 用户切片实现
$user1 := []string{"1", "张三", "12", "男"}
$user2 := []string{"2", "赵六", "12", "男"}
$user3 := []string{"3", "李四", "12", "女"}
$user4 := []string{"4", "王麻子", "12", "男"}
```
通过上面的示例代码，我们是不是看得出来存在这样几个问题。

a. 可读性低。我们完全不知道1、12这样的值是用户的什么信息，男、张三我们还可以猜测一下是名字和性别。

b. 重复代码。一个用户一个变量，如果存在千万个用户，我们岂不是需要定义千万个变量。

c. 繁琐。相比PHP的实现，是不是非常繁琐。PHP中直接定义一个变量，通过多维数组的方式，就可以定义key和值。清晰并且简单。这也是为什么大家都说PHP中的数组非常强大和好用了。

d. 通过切片和数组实现的方式，我们知道了弊端。那有不有一种数据类型能够像PHP这样简单就能实现呢？这样的场景就可以用map实现PHP这样的定义结构。接下来，我们就具体总结一下map相关的操作。

## map

### map定义

map 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构也称为关联数组或字典。这是一种快速寻找值的理想结构：给定 key，对应的 value 可以迅速定位。
map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。

### map声明

map属于一种引用类型，在使用时我们需要make给其分配内存空间，未分配内存空间的map值是一个nil。
```golang
map声明时，需要指定key的类型和值的类型，并且复制时，必须按照定义时的类型进行复制。
map的值可以是任意类型，可以是切片可以是数组，可以是接口、结构体、指针、字符串等等数据类型。
var map1 map[key类型]值类型

// 声明方式一，完整模式
var map1 map[int]string
map1 = make(map[int]string[, n])

// 声明方式二，段语法风格
map1 := make(map[int]string[, n])
```
> 上面的n都是map的容量，也就是说map可以存储元素的数量。可以省略，map会动态扩容。

示例小案例，我们用map存储一个用户的信息。用户信息包含ID，name，age字段。
```golang
userInfo := make(map[string]string)
userInfo["id"] = "1"
userInfo["name"] = "张三"
userInfo["age"] = "12"
fmt.Println(userInfo)

// output
map[id:1 name:张三 age:12]
```
> 因为用户信息的字段和字段值有字符串和数字类型，定义好类型之后就只能传对应类型的值，因此我们给key和value的类型都定义为string类型。

### map的操作

这里的操作，我们接着上面的小案例来使用。
1. 访问和复制。我们直接使用下标就可以了。

```golang
// 赋值
mapName[key] = "值"
userInfo["name"] = "王五"

// 访问
mapName[key]
name := userInfo["name"]
```

2. 删除元素。删除操作，需要使用到delete()，给该函数传递要删除的key。

```golang
delete(mapName, key)

delete(userInfo, "name")
```

3. 判断某一个值key是否存在。上面我们访问map中的key，直接使用下标就可以了。如果 map 中不存在 key1，val1 就是一个值类型的空值。会导致我们没法区分到底是 key不存在还是它对应的value就是空值。

```golang
value, boolean := mapName[key]
// 如果boolean是一个true则说明，对应的key存在，如果是false，则说明该key不存在。

if value, ok := userInfo["address"]; !ok {
  fmt.Println("address key not found!")
}
因为address这个key不存在，因此会输出“address key not found!”。
```

4. 循环。循环map，我们一般是用到 `for range`来实现。

```golang
// 语法
for value, key := range mapName {
  fmt.Println(key, "=>", value)
}
// 示例
for value, key := range userInfo {
    fmt.Println(key, "=>", value)
}
// output
1 => id
张三 => name
12 => age
```

### 总结

其实对map基本的操作就是这么简单。对它的理解也是这么简单。在日常开发中，我们也经常使用该类型。

回到最上面多个用户的案例，这时候我们是不是就知道怎么使用map实现了。
1. 因为是多个用户，我们是不是需要定义多维的map结构。

2. map的一级key是int，表示当前的用户序号(从0，1，2，3...依次递增)。key对应的值，才是某一个用户的具体信息，我们同样的定义map类型来存储，key和value都是字符串，结构就像map声明中的小案例一样。

3. 因为我们不知道用户的具体个数，我们将一级的key定义为切片。因为切片的长度是不固定的。

```golang
userInfo := make([]map[string]string, 3)

for i := 0; i < 3; i++ {
    userInfo[i] = make(map[string]string, 3)
     userInfo[i]["id"] = "ID"
     userInfo[i]["name"] = "名称"
     userInfo[i]["age"] = "年龄"
}
fmt.Println(userInfo)

// output
[
  map[id:ID name:名称 age:年龄] 
  map[id:ID name:名称 age:年龄] 
  map[id:ID name:名称 age:年龄]
]
```
> 为什么要两次make，因为切片和map在golang中都是引用类型。第一次在make时，是针对切片初始化内存空间，第二次是针对切片的key对应的元素分配内存空间。大致的结构就像下图一样。

![Snipaste_2021-11-20_22-36-42](https://gitee.com/bruce_qiq/picture/raw/master/2021-11-20/1637419026851-Snipaste_2021-11-20_22-36-42.png)
