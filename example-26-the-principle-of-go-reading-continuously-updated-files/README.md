# Go 读取持续更新的文件原理

#### 通过 `tail` 包捕获 append 的内容

演示通过 golang 代码实现 Linux tail 命令行读取持续更新文件内容的操作.
```go
go run demo0.go
// echo $$ >> demo.log
```

#### 文件系统通知

我们要实时捕获文件中保存的新行, 这里就需要根据文件系统通知, 来触发新行的读取

通过 `fsnotify` 包捕获事件.

```go
go run demo1.go
// echo $$ >> demo.log
// output:
// 2020/09/03 16:05:39 event: "demo.log": WRITE
// 2020/09/03 16:05:39 modified file: demo.log
```

- NewWatcher 函数返回 *Watcher.
- 在一个新的协程中通过 for 循环和 select 捕获 *Watcher.Events 和 *Watcher.Errors 两个 channel 中的数据. 从而捕获事件.
- 在 main 协程为 *Watcher 添加文件.
- 阻塞 main 协程.

#### 参考
- [Go 1.5 os/fsnotify API (草稿第五版)](https://docs.google.com/document/d/1-GQrFdDVrA57-ce0kbzSth4lQqfOMMRKpih3hPJmvoU/edit#heading=h.7ush6jpv61gf)
