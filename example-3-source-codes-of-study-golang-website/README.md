# Study Golang Website
Golang 中文网是一个基于 Go 和相关组建构建的 Golang 编程语言相关的中文站点

#### 如何实现配置文件重新加载

衍生问题:
- golang 服务如何捕获特定信号

```go
go run demo0.go
```

流程概述:
- 定义一个 buffered channel
- 只捕获 `syscall.SIGUSR1`
- 在一个循环内
  - 从 buffered channel 中读取信号(中断)
    - 如果是 `syscall.SIGUSR1`, 重新加载配置
    - 可以通过 `kill -USR1 <pid>` 来触发这个信号, `<pid>` 可以通过 `ps -ef |grep demo0` 找到. 或者直接执行 `kill -USR1 $(pidof demo0)`

#### 如何管理 API

`echo` 库 demo
```bash
go run demo1.go
# 通过 `curl localhost:1323/` 访问
```

另外, 参见 [echo guide 文档][1]
##### 中间件
TL;DR: 中间件是在 request-response 周期中对 `echo#Context` 进行处理的链状函数( a 中间件处理完, b 继续处理 )

可以看看 [echo 文档][2] 中的解释

分为 4 个阶段, 分别是:
- Root Level (Before router)
- Root Level (After router)
- Group Level
- Route Level

##### 用了哪些中间件

```go
	e.Use(thirdmw.EchoLogger()) // 日志
	e.Use(mw.Recover()) // 异常恢复
	e.Use(pwm.Installed(filterPrefixs)) // 程序首次安装的流程
	e.Use(pwm.HTTPError()) // http 错误码处理. 比如 404 页面的渲染等
	e.Use(pwm.AutoLogin())
```

#### echo 中间件机制
<!--todo: 如何实现-->


##### 组路由
<!--todo: 什么是组路由-->

##### 组路由中间件

#### 初始化流程
初始化流程指代码第一次部署. 通过客户通过 web 界面提供数据库等配置信息

#### 如何管理数据库

#### 如何实现中文分词

```go
go run demo1
```

#### 如何使用 pprof

#### 如何配置定时服务

#### 如何实现用户登陆

#### 参考文档
- [1]: https://echo.labstack.com/guide
- [2]: https://echo.labstack.com/middleware

