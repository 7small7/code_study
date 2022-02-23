## 数组的含义

数组是**相同类型**的变量的集合。也可以理解为**内存地址连续**的一组变量的集合。

## 数据结构
![go-array1](http://qiniucloud.qqdeveloper.com/mweb/go-array1.png)

## 数组的定义

### 基本格式

```go
// 指定初始值
数组名称 := [数组长度]数组类型 {初始值1,初始值2,......,初始值n}
// 初始化一个数组
var 数组名 [数组长度]数组类型
```

#### 情况分析

1. 数组定义时长度和类型不能省略。如果长度是"..."，此时数组的长度会根据初始值的元素个数而定。

2. 数组的长度可以是一个**常量表达式**，但是编译阶段的数组长度值必须是一个整数类型的值，不能是其他的类型。比如编译后的值是一个字符串的0("0"),这种是错误的。

3. 数组的小标是从0开始的。数组可以创建一个空数组，也可以创建一个分配好的值的数据。

4. 数组的类型和数组的长度是数组中的一部分，因此如果类型相同并且数组元素也相同的数组，但是长度不一样，是不同的数组类型。

```go
// 这两个数组不是不是同一类型的数组
var arrray1 = [10]int {1, 2}
var arrray2 = [11]int {1, 2}
// 这两组数组是同一类型的数组
var arrray1 = [10]int {1, 2}
var arrray2 = [...]int {1, 2}
```

5. 当未设置初始值时，根据数组的长度和数组的类型，编译时自动设置默认值。如果设置的是"..."，编译时则默认根据类型自动设置一个初始值。

#### 默认值 

1. 当数组类型是数字类型时，初始值是0，

2. 当数组类型是string，初始值为空，

3. 当数组的类型是bool类型时，初始值为false

#### 代码演示

```go
func IndexArray() {
	// 长度和初始值都不设置，打印一个空数组,[]
	array1 := [...]int{}
	fmt.Println(array1)
	// 长度设置，初始值不设置，打印一个长度为1，初始值为0的数组,[0]
	array2 := [1]int{}
	fmt.Println(array2)
	// 长度不设置，设置初始值，打印出初始值[1 2]
	array3 := [...]int{1,2}
	fmt.Println(array3)
	// 长度和初始值都设置，打印出初始值[1 2 3 4]
	array4 := [4]int{1,2,3,4}
	fmt.Println(array4)
	// 长度设置，初始值设置的个数小于长度,未设置的初始值根据数组类型编译时自动设置默认值，打印出[1 2 0 0]
	array5 := [4]int{1,2}
	fmt.Println(array5)
	// 长度设置，指定索引设置初始值，其他未设置初始值的编译时自动设置默认值，打印出[0 0 0 0 0 1]
	array6 := [6]int{5: 1}
	fmt.Println(array6)
	// 长度和初始值设置，指定索引,打印出[2 2 2]
	array7 := [3]int{0:2,1:2,2:2}
	fmt.Println(array7)
}
```

## 多维数组

### 多维数组定义

```go
var 数组名称 [数组长度][二维的长度]....[n维的长度] 数组类型
```

1. 不管是多少维度的数组，数据类型必须和定义时保持一致。

### 代码演示

```go
// 多维数组的定义
func MoreArray() {
	var array1 [1][2]int
	var array2 = [1][2]int{
		{1, 2},
	}
	// 正确定义
	var array3 = [3][2][3]int{
		{{1, 2, 3}, {3, 2, 1}},
	}
	// 正确定义
	var array3 = [3][2][3]int{
		{
			{1, 2, 3}, 
			{3, 2, 1}},
	}
	// 错误定义
	var array3 = [3][2][3]int{
		{
		  {1, 2, 3}, 
		  {3, 2, 1},
		},
	}
}
```
```go
// output
[[[1 2 3] [3 2 1]] [[0 0 0] [0 0 0]] [[0 0 0] [0 0 0]]]
```

> 多维数组的常见集中情况和上面一维数组的常见集中情况一样，但是有如下不同的一个地方。多维数组的顶层必须设置长度

```go
// 多维数组需要注意的地方
// 错误的定义方式
var array3 = [...][2][3]int{
    {{1, 2, 3}, {3, 2, 1}},
}
// 正确的定义方式(注意第二行结尾的","不能省略)
var array3 = [][2][3]int{
    {{1, 2, 3}, 
    {3, 2, 1}},
}
```
```go
//output
[
   [[1 2 3] [3 2 1]]
   [[0 0 0] [0 0 0]] 
   [[0 0 0] [0 0 0]]
]
```

## 数组的常用操作

### 循环

1. for方式
```go
func LoopArray() {
	array := [4]int{1,2,3,4}
	len := len(array)
	for i := 0; i < len; i++ {
		fmt.Printf("数组array的索引%d对应的值是%d", i, array[i])
		fmt.Println()
	}
}
```

2. range方式
```go
arr := [...]int{1, 2, 3, 4}
	for index, value := range arr {
		fmt.Println(index, value)
	}
```

### 计算长度

```go
len := len(数组名)
```
> 当数组是一个多维数组时，直接使用len(数组名)，计算的是顶层的长度

### 赋值

```go
数组名[下标] = 值
array := [4]int{1,2,3}
array[3] = 100
```
### 比较

如果两个数组类型相同（包括数组的长度，数组中元素的类型）的情况下，我们可以直接通过较运算符（==和!=）来判断两个数组是否相等，只有当两个数组的所有元素都是相等的时候数组才是相等的，不能比较两个类型不同的数组，否则程序将无法完成编译。
```go
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"
d := [3]int{1, 2}
fmt.Println(a == d) // 编译错误：无法比较 [2]int == [3]int
```
### 数组与函数的使用

```go
func 函数名(形参名称 [数组长度]数组类型) 函数返回值 {
    // 函数体
}
```
> 形参中，函数的长度为可写

```go
// 数组与函数的使用
func Function(arr [] int) int {
    return len(arr)
}

Function([5]int{})
```

### 排序

#### 冒泡排序

```go
// 从大到小进行排序
func BubbleSort() {

	array := [...]int{2,9,3,5,1,2}
	len := len(array)
	fmt.Println(array)

	for i := 1; i < len; i++ {// 控制外层循环次数
		for j := 0; j < len-i; j++ { // 控制内部每次循环个数
			if array[j] < array[j+1] {
				tmp := array[j]
				array[j] = array[j+1]
				array[j+1] = tmp
			}
		}
	}
	fmt.Println(array)
}
```

#### 插入排序

```go
// 插入排序
func InsertSort() {
	array := [...]int{2,9,3,5,1,2}
	len := len(array)
	fmt.Println(array)
	// var j int = 0
	// for i := 1; i < len; i++ {
	// 	tmp := array[i]
	// 	for j = i; j > 0 && tmp < array[j - 1]; j-- {
	// 		array[j] = array[j - 1]
	// 	}
	// 	array[j] = tmp
	// }
	for i := 1; i < len; i++ {// 控制外层循环个数
		tmp := array[i]
		for j := i-1; j >= 0; j-- {// 针对已经排序好的数组再次排序
			if tmp > array[j] {
				array[j+1] = array[j]
				array[j] = tmp
			}
		}
	}
	fmt.Println(array)
}
```