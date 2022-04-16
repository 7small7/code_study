[TOC]

## 定义

defer是golang语言中的一种延迟机制。经过defer定义的语句会在**函数执行结束之后被执行**。

## 使用场景

处理业务或逻辑中涉及成对的操作是一件比较烦琐的事情，比如打开和关闭文件、接收请求和回复请求、加锁和解锁等。在这些操作中，最容易忽略的就是在每个函数退出处正确地释放和关闭资源。

## 实现原理

defer实现的原理是，当编译器发现某段语句被defer修饰，就会将该段语句**放入一个独立的栈**之中，当前函数被执行完之后，接着执行该独立栈中的语句。
如果一个函数中存在多个defer语句，则**按照栈先进后出的特点依次被执行**。

## 代码演示

```go
package main

import "fmt"

func sum(n1, n2 int) int {
  // 此时会将defer语句放入栈中 
	defer fmt.Println("n1 = ", n1)
	defer fmt.Println("n2 = ", n2)
	sum := n1 + n2
	fmt.Println("sum = ", sum)
	// return之后，函数被执行完毕，接着执行栈中的语句
	return sum
}

func main()  {
	result := sum(10, 20)
	fmt.Println("result = ", result)
}
```
```go
//output
sum =  30
n2 =  20
n1 =  10
result =  30
```

### 注意事项

在被defer定义之后的语句，也会将值进行拷贝(值拷贝)放入栈中。
```go
package main

import "fmt"

func sum(n1, n2 int) int {
	defer fmt.Println("n1 = ", n1)
	defer fmt.Println("n2 = ", n2)
	n1++
	n2++
	sum := n1 + n2
	fmt.Println("sum = ", sum)
	return sum
}

func main()  {
	result := sum(10, 20)
	fmt.Println("result = ", result)
}
```
```go
// output
sum =  32
n2 =  20
n1 =  10
result =  32
```
> 此时你会发现sum的值和result的值变了，但是n1和n2输出的值未变。