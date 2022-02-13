开发了一款免费的程序员答题小程序，上面的所有问题都是来自于真实面试题。涵盖了Redis、MySQL、Golang、PHP等等技术栈。点击即可进入答题。
## 包介绍

在使用Golang开发中，我们不可能把所有的项目文件都放在一个目录下面。这就需要根据实际的项目，将程序文件进行归类，不同的功能放在不同的目录。这就是包的作用之一，Golang中的包就像PHP中的命名空间类似。

## 语法

### 定义包

假设我们创建了一个名字叫做pack1的目录，此时我们就需要把这个目录下面的文件都定义为pack1包名。

```go
pack pack1
```

1. 使用关键字`pack` + 包名。包名不能使用Golang自带的关键字。

2. 包名必须放在程序文件的第一行，在定义包之前不能有任何内容。

3. main包作为程序的入口文件，其他的程序目录不能定义为main包，但是与main包同目录下面的程序文件可以定义包名为main包。举个例子，你定义了一个程序文件a.go，同时在该目录下面定义了程序文件b.go。下面这种写法是正确的。

```go
// a.go
pack main

func main() {
  // todo
}
```
```
// b.go
pack main

func show() {
  // todo
}
```

### 包的使用

当前包要使用其他包的元素(变量、函数、结构体等等)，就需要导入到当前包。这和PHP中的use语法一样。Golang中导入包使用关键字`import`。
```
// 只到如一个包
import "go_demo/src/demo/pack1"

// 导入多个包
import (
	"go_demo/src/demo/pack1"
	"go_demo/src/demo/pack2"
)
```
> 上面的两种写法都是正确的，主要是看你导入包的数量，不过推荐使用第二种，毕竟一个包基本都会导入多个外部包。

### 导包的注意事项

1. 导入的包，根据包的类型进行分组。一般我们把包分为Golang内置包和第三方包。推荐使用下面的导包风格。中间多一空行，增加代码的可阅读性。

```go
import (
	"os"

	"go_demo/src/demo/pack1"
	"go_demo/src/demo/pack2"
)
```

2. 尽可能导入需要使用的包，不用的包就不要导入。如果导入的包不使用，在编译时会报错，要解决这个问题，可以使用下面的语法，添加一个"_"。

```
import (
	_ "go_demo/src/demo/pack1"
	"go_demo/src/demo/pack2"
)
```

3. 包别名。当引入第三方包和自身定义的包名，发生冲突时，可以针对包做一个别名。下面的`githubPacke`就是一个包别名。

```go
import (
	"go_demo/src/demo/pack1"
	githubPacke "github.com/demo/pack1"
)
```

4. 调用包中的属性时，可以省略包名，但是不推荐这种方式，很容易发生使用冲突。假设pack1包中有一个Show()函数，pack2中有一个Run()函数。

```go
import (
	"go_demo/src/demo/pack1"
	"go_demo/src/demo/pack2"
)
func Demo() {
 // 直接省略掉包名。
  Show()
  Run()
}
```

1. 包属性。这里的属性是指，包内的变量、结构体、方法、函数等等。如果在开头字母是小写，则只能在本包中调用，不能被外部包调用。否则会出现`cannot refer to unexported name pack1.demo2`类似的错误信息。下面的Name变量、Demo()1函数就能被外部包调用。

```go
var Name string
var age int

func Demo1() {
	pack3.Demo1()
	fmt.Println("pack1->Demo1()")
}

func demo2() {
	fmt.Println("pack1->demo2()")
}
```

1. 调用外部包时，会自动调用包中的init()函数。越底层的包，越先执行。下面有main.go文件，packe1.demo1.go文件，pack2.demo1.go和pack3.demo1.go文件。当main.go去调用pack1和pack2的时候，发现pack1调用了pack3，于是最先调用的包应该就是pack3中的init()函数。

![Snipaste_2022-02-14_00-48-18](https://gitee.com/bruce_qiq/picture/raw/master/2022-2-14/1644770912715-Snipaste_2022-02-14_00-48-18.png)

main.go文件。
```go
import (
	"go_demo/src/demo/pack1"
	"go_demo/src/demo/pack2"
)

func main() {
	pack1.Demo1()
	pack2.Demo1()
}
```
packe1.demo1.go文件
```golang

func init() {
	fmt.Println("pack1->init()")
}

func Demo1() {
	pack3.Demo1()
	fmt.Println("pack1->Demo1()")
}
```
packe2.demo1.go文件
```golang
func init() {
	fmt.Println("pack2->init()")
}

func Demo1() {
	fmt.Println("pack2->Demo1()")
}
```
packe2.demo1.go文件
```golang
func init() {
	fmt.Println("pack3->init()")
}

func Demo1() {
	fmt.Println("pack3->Demo1()")
}
```
输出结果为：
```go
pack3->init()
pack1->init()
pack2->init()
pack3->Demo1()
pack1->Demo1()
pack2->Demo1()
```