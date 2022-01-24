## 说说php的同步模式与swoole的携程之间的区别？
1. 首先，Swoole 只能运行在命令行（Cli）模式下，所以我们开发调试都是使用命令行，而不是 php-fpm/apache 等。在 Swoole 中，我们可以使用`\Swoole\Coroutine::create()`创建协程，或者你也可以使用简写`go()`。
2. 我们一直在说 Swoole 协程适合用于 I/O 密集场景，在同样的硬件配置环境下，它会比传统的同步模式承载更多的访问量。我们熟悉的文件读写、网络通讯请求（MySQL、Redis、Http等）都是属于 I/O 密集型场景。
假设一次 SQL 查询为 100ms，在传统同步模式下，当前进程在这 100ms 的时间里，是不能做其它操作的。如果要执行十次这个 SQL，可能需要耗费 1s 以上。而如果用协程，虽然不同协程之间也是按顺序执行，但是在前一个等待 100ms 期间，底层会调度 CPU，去执行其它协程的操作。也就是说，可能第一个查询还没返回结果，其它几个查询就已经发送给了 MySQL 并正在执行中了。如果开启十个协程，分别执行这个 SQL，可能只需要耗费 100+ms 即可完成。

## php-fpm与swoole之间有什么区别？
php-fpm与swoole介绍：
1. 早期版本的 PHP 并没有内置的 WEB 服务器，而是提供了 SAPI（Server API）给第三方做对接。现在非常流行的 php-fpm 就是通过 FastCGI 协议来处理 PHP 与第三方 WEB 服务器之间的通信。
2. Swoole 采用的也是 Master/Worker 模式，不同的是 Master 进程有多个 Reactor 线程，Master 只是一个事件发生器，负责监听 Socket 句柄的事件变化。Worker 以多进程的方式运行，接收来自 Reactor 线程的请求，并执行回调函数（PHP 编写的）。启动 Master 进程的流程大致是：
  初始化模块。
  初始化请求。因为 swoole 需要通过 cli 的方式运行，所以初始化请求时，不会初始化 PHP 的全局变量，如 $_SERVER, $_POST, $_GET 等。
  执行 PHP 脚本。包括词法、语法分析，变量、函数、类的初始化等，Master 进入监听状态，并不会结束进程。

Swoole 加速的原理
由 Reactor（epoll 的 IO 复用方式）负责监听 Socket 句柄的事件变化，解决高并发问题。通过内存常驻的方式节省 PHP 代码初始化的时间，在使用笨重的框架时，用 swoole 加速效果是非常明显的。

php-fpm与swoole区别
1. PHP-FPM是Master 主进程 / Worker 多进程模式。
2. 启动 Master，通过 FastCGI 协议监听来自 Nginx 传输的请求。
3. 每个 Worker 进程只对应一个连接，用于执行完整的 PHP 代码。
4. PHP 代码执行完毕，占用的内存会全部销毁，下一次请求需要重新再进行初始化等各种繁琐的操作。
5. 比较适用于HTTP Server。

1. Swoole是Master 主进程（由多个 Reactor 线程组成）/ Worker 多进程（或多线程）模式。
2. 启动 Master，初始化 PHP 模块，由 Reactor 监听 Socket 句柄的事件变化。
3. Reactor 主线程负责子多线程的均衡问题，Manager 进程管理 Worker 多进程，包括 TaskWorker 的进程。
4. 每个 Worker 接受来自 Reactor 的请求，只需要执行回调函数部分的 PHP 代码。
5. 只在 Master 启动时执行一遍 PHP 初始化代码，Master 进入监听状态，并不会结束进程。
6. 不仅可以用于 HTTP Server，还可以建立 TCP 连接、WebSocket 连接。