### 模式说明

`Hello World`是rabbitMQ中最简单的一种模式。该模式由`生产者`、`消息队列`、`消费者`三部分组成。
![Snipaste_2022-04-16_19-41-49](https://files.mdnice.com/user/5113/62f7063e-8f4c-41b1-8d8c-ad5a08095256.png)
1. 生产者通过channel连接连接到rabbitMQ服务。
2. 通过channel连接到rabbitMQ队列并向队列中发送消息。
3. 消费者通过channel连接连接到rabbitMQ服务，并一直处于监听状态。
4. 当rabbitMQ队列中有数据，则通过channel发送给消费者。

### 安装rabbitMQ包

```go
go get github.com/streadway/amqp
```

### 定义输出函数

```go
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		return
	}
}
```

### 定义生产者

```go
package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"

	"go_demo/src/rabbitmq/com"
)

func main() {
	// 1. 创建套字节
	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		com.FailOnError(err, "RabbitMQ connection fail!")
	}
	defer dial.Close()

	// 2. 创建链接
	channel, err := dial.Channel()
	if err != nil {
		com.FailOnError(err, "RabbitMQ channel fail!")
	}
	defer channel.Close()

	// 3. 连接队列
	declare, err := channel.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		com.FailOnError(err, "RabbitMQ declare queue fail!")
	}

	// 4. 发送消息
	for {
		body := fmt.Sprintf("%s\t%s", "Hello, World!", time.Now())
		err = channel.Publish(
			"",           // 交换机
			declare.Name, // 路由key
			false,        // 强制
			false,        // 立即
			amqp.Publishing{
				ContentType: "text/plain", // 消息类型
				Body:        []byte(body), // 消息内容
			})
		if err != nil {
			com.FailOnError(err, "RabbitMQ send message fail!")
		}
		fmt.Println("send success")
		time.Sleep(time.Second * 2)
	}
}
```

### 发送消息

```go
go run product.go
```
此时没有输出错误信息，则表示生产者执行成功。访问web端管理界面，就可以看到创建的队列以及队列消息。
![Snipaste_2022-04-16_19-57-11](https://files.mdnice.com/user/5113/533ee02a-6651-49a6-b8cb-6eb61b072abc.png)

### 定义消费者

```go
package main

import (
	"fmt"

	"github.com/streadway/amqp"

	"go_demo/src/rabbitmq/com"
)

func main() {
	// 1. 创建套字节
	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		com.FailOnError(err, "RabbitMQ connection fail!")
	}
	defer dial.Close()

	// 2. 创建链接
	channel, err := dial.Channel()
	if err != nil {
		com.FailOnError(err, "连接失败")
	}
	defer channel.Close()

	// 3. 申明队列(当消费者先启动时，如果rabbitMQ中没有该队列，告知rabbitMQ创建好队列。)
	declare, err := channel.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		com.FailOnError(err, "队列失败")
	}

	// 4. 消费信息
	consume, err := channel.Consume(
		declare.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		com.FailOnError(err, "消费失败")
	}

	for d := range consume {
		fmt.Println(string(d.Body))
	}
}
```
### 消费消息

```go
go run consumer.go
```
通过执行上面的命令，我们就可以看到队列中的消息被正常输出了。
![Snipaste_2022-04-16_20-00-01](https://files.mdnice.com/user/5113/5ba65a62-5b0c-453d-a3c8-69be5388d2f4.png)
