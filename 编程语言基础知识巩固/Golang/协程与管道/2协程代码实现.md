[TOC]

## 同步阻塞演示

下面的代码，未使用协程，属于一个同步阻塞执行(串行)的代码。我们可以看到下面的输出也是存在一个先后顺序进行输出。
```go
package main

import (
	"fmt"
)

func Hello()  {
	fmt.Println("Hello")
}

func main() {
	Hello()
	fmt.Println("Main")
}
```
```go
// output
go run goroutine.go
Hello
Main
```
## 协程简单实现

Golang中实现协程是非常简单的，直接在指定函数前面加一个go 关键字即可。
```go
package main

import (
	"fmt"
)

func Hello()  {
	fmt.Println("Hello")
}

func main() {
	// 开启一个协程
	go Hello()
	fmt.Println("Main")
}
```
```go
// output
go run goroutine.go
Main
```
> 通过上面的代码,我们在Hello函数前添加了一个关键字go，但发现Hello函数中的输出并未按照预期的输出。这是因为在启动程序时，首先有一个主线程，我们这里在主线程运行过程中开启了一个协程，当Hello函数还未输出，主线程就结束了，因此开启的协程也随之结束了。简而言之，就是协程还没执行，主线程就结束了，协程根本没有执行的机会。下面两种方式讲解如何解决此问题。

## 协程方式一

方式一，我们可以通过线程之前同步的方式。也就是说每个协程执行结束之后，主线程才结束。

```go
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hello(i int) {
	defer wg.Done() // 减少WaitGroup计数器的值，应在线程的最后执行。
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 向协程计数器加1。
		go Hello(i)
	}
	wg.Wait() // 阻塞协程，直到协程计数器为0。
	fmt.Println("所有的协程已经结束了，主线程可以结束了。")
}
```
```go
go run goroutine.go
Hello Goroutine! 2
Hello Goroutine! 9
Hello Goroutine! 0
Hello Goroutine! 3
Hello Goroutine! 7
Hello Goroutine! 1
Hello Goroutine! 8
Hello Goroutine! 5
Hello Goroutine! 4
Hello Goroutine! 6
所有的协程已经结束了，主线程可以结束了。
```
> 通过上面的输出结果，我们可以看到协程对应的内容进行输出了。同时，我们可以看到启动的协程并未是按照顺序输出的，而是随机输出。是因为协程在执行过程中就是随机执行，而不是顺序执行。

## 实现方式二

该方式是通过睡眠时间，就是让主线程处于睡眠等待而不是立即结束，在等待的时间段内，协程就可以执行。

```go
package main

import (
	"fmt"
	"time"
)

func Show(ch1 chan int) {
	val := <-ch1
	fmt.Println("通道的值是:", val)
}

func main() {
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		go Show(ch1)
		ch1 <- i
	}
	time.Sleep(time.Second * 10)
	close(ch1)
	fmt.Println("所有的协程已经结束了，主线程可以结束了。")
}
```
```go
go run goroutine.go
通道的值是: 4
通道的值是: 0
通道的值是: 1
通道的值是: 7
通道的值是: 9
通道的值是: 6
通道的值是: 8
通道的值是: 3
通道的值是: 2
通道的值是: 5
所有的协程已经结束了，主线程可以结束了。
```
> 通过该方式，我们能够确保协程进行执行。但强烈推荐不要使用该方式。存在协程执行完，主线程的睡眠时间还未到，程序还是处于执行状态。同时，如果主线程睡眠时间到了，但是协程任务还未执行完毕，剩下的任务就不会被执行。