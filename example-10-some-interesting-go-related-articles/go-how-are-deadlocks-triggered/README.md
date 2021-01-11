# 如何触发 Golang 死锁

原文地址: https://medium.com/a-journey-with-go/go-how-are-deadlocks-triggered-2305504ac019

原文介绍了作者自己对死锁的理解. 介绍了如何发现死锁(我不是很理解). 然后提供了一些代码用于演示. 作者的演示代码和方法我不是很理解. 所以大概了解了一下 `go tool trace` 和 `go tool pprof` 还有如何打印调度器信息的用法. 然后写几个 demo 记录下来.

#### 关于死锁

死锁是一种状态，当goroutine被阻止而无法解除阻止时，就会发生这种状态。 Go提供了一个死锁检测器，可帮助开发人员避免陷入这种情况。

#### demo

一个简单的死锁:

```bash
go run demo0.go
```

main goroutine 由于等待其他人推送数据到 channel 而被阻塞. 但是由于没有其他的 goroutine 在运行, 所以 main goroutine 将永远不会被解阻塞. 这种情况就会触发死锁.

#### demo1

演示如何捕获特定信号

```bash
go run demo1.go
```

#### demo2

演示如何捕获特定信号并打印跟踪信息用于分析性能

首先我们执行 `demo2.go` 中的代码:

```bash
go run demo2.go
```

然后, 我们打开一个新的命令行窗口, 发送特定的信号, 触发 demo 退出

```bash
kill -USR1 `pidof demo2`
```

本地会生成一份 `trace.out` 文件, 里面包含了跟踪信息, 我们通过 `go tool trace` 加载并通过 web 展示这些信息

```bash
go tool trace trace.out
```

接下来就是通过 web 界面的信息分析性能了. 可以参考 [煎鱼关于 `go tool trace` 的文章](https://eddycjy.com/posts/go/tools/2019-07-12-go-tool-trace/)

#### demo3

演示如何打印调度器信息

```go
GODEBUG=schedtrace=1000 go run demo3.go
```

运行完上面这段命令, 你会在命令行看到很多类似于下面这行的输出, 这是摘录其中一行的解释

> 下面的这段解释出自: https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/ , 我只是根据我这边的输出, 替换了一些数据

```
SCHED 1006ms: gomaxprocs=8 idleprocs=8 threads=7 spinningthreads=0 idlethreads=3 runqueue=0 [0 0 0 0 0 0 0 0]
```

- SCHED：调试信息输出标志字符串，代表本行是goroutine scheduler的输出；
- 1006ms：即从程序启动到输出这行日志的时间；
- gomaxprocs: P的数量；
- idleprocs: 处于idle状态的P的数量；通过gomaxprocs和idleprocs的差值，我们就可知道执行go代码的P的数量；
- threads: os threads的数量，包含scheduler使用的m数量，加上runtime自用的类似sysmon这样的thread的数量；
- spinningthreads: 处于自旋状态的os thread数量；
- idlethread: 处于idle状态的os thread的数量；
- runqueue=0： go scheduler全局队列中G的数量；
- [0 0 0 0 0 0 0 0]: 分别为8个P的local queue中的G的数量。

#### demo4

##### 如何查看火焰图和 CPU profile

```go
go run demo4.go
```

然后分析 CPU 的 profile, 执行下面的命令捕获 5s 内的数据：

```bash
# 命令行
go tool pprof http://localhost:6060/debug/pprof/profile\?seconds\=5

# web 可以看到火焰图
# 煎鱼的文章: `Go 大杀器之性能剖析 PProf' 内有更详细的信息, 建议看看
go tool pprof -http=:8081 http://localhost:6060/debug/pprof/profile\?seconds\=5
```

##### 如何查看跟踪信息

```go
go run demo4.go
```

执行下面的命令捕获 5s 内的 trace 数据, 并打开 trace 分析页面

```bash
curl -o trace.out http://127.0.0.1:6060/debug/pprof/trace?seconds=5
go tool trace trace.out
```

#### 参考
- [煎鱼关于 `go tool trace` 的文章](https://eddycjy.com/posts/go/tools/2019-07-12-go-tool-trace/)
- [GO MEMORY MANAGEMENT PART 3](https://povilasv.me/go-memory-management-part-3/)
- [https://www.alexedwards.net/blog/an-overview-of-go-tooling](https://www.alexedwards.net/blog/an-overview-of-go-tooling)
- [An incomplete list of Go tools](http://dominik.honnef.co/posts/2014/12/go-tools/)
- [sourcegraph 关于 `go tool trace` 的文章](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner/)
- [golang自动检测死锁deadlock的实现](http://xiaorui.cc/archives/5951)
- [golang 执行追踪器文档](https://docs.google.com/document/u/1/d/1FP5apqzBgr7ahCCgFO-yoVhk4YZrNIDNf9RybngBc14/pub)
- [Go 大杀器之性能剖析 PProf](https://eddycjy.com/posts/go/tools/2018-09-15-go-tool-pprof/)
