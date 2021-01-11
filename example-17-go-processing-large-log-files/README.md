# Go 处理大日志文件

#### 关于 `sync.Pool`
官方文档:
A Pool is a set of temporary objects that may be individually saved and retrieved.
`Pool` 是一组可以单独保存和获取的临时对象的集合

Any item stored in the Pool may be removed automatically at any time without notification. If the Pool holds the only reference when this happens, the item might be deallocated.

`Pool` 中存储的任何数据项都可能在没有任何通知的情况下自动删除。如果 `Pool` 含有该数据项唯一的引用, 当发生这种情况时(译者注: 前面提到的自动删除)，则该数据项对应的内存块可能会被释放。

A Pool is safe for use by multiple goroutines simultaneously.

Pool's purpose is to cache allocated but unused items for later reuse, relieving pressure on the garbage collector. That is, it makes it easy to build efficient, thread-safe free lists. However, it is not suitable for all free lists.

`Pool` 可安全地同时用于多个 goroutine。

`Pool` 的目的是缓存已分配但未使用的数据项以供稍后重用，从而减轻了垃圾收集器的压力。也就是说，`Pool` 将使构建高效，线程安全的空闲列表变得容易。但是，它并不适合所有空闲列表。

An appropriate use of a Pool is to manage a group of temporary items silently shared among and potentially reused by concurrent independent clients of a package. Pool provides a way to amortize allocation overhead across many clients.

An example of good use of a Pool is in the fmt package, which maintains a dynamically-sized store of temporary output buffers. The store scales under load (when many goroutines are actively printing) and shrinks when quiescent.

池的适当用法是管理一组临时数据项，这些临时数据项在程序包的并发独立客户端之间静默共享并有可能被重用。`Pool` 提供了一种分摊许多客户端上的内存分配开销的方法。

良好使用 `Pool` 的一个示例是 fmt 软件包，该软件包维护着动态大小的临时输出缓冲区存储。存储空间在负载下会扩张（当许多goroutine正在主动打印时），并且在不活动时会收缩。

On the other hand, a free list maintained as part of a short-lived object is not a suitable use for a Pool, since the overhead does not amortize well in that scenario. It is more efficient to have such objects implement their own free list.

A Pool must not be copied after first use.

另一方面，作为短期对象的一部分维护的空闲 `list` 不适用于 `Pool`，因为在这种情况下开销无法很好地摊销。使此类对象实现其自己的空闲 `list` 更为有效。

`Pool` 在首次使用后不能再被拷贝

官方示例
```shell script
go run demo0.go
```

#### 参考
- http://www.linvon.cn/posts/golang%E5%BF%AB%E9%80%9F%E8%AF%BB%E5%8F%96%E5%A4%84%E7%90%86%E5%A4%A7%E6%97%A5%E5%BF%97%E6%96%87%E4%BB%B6/
- https://medium.com/swlh/processing-16gb-file-in-seconds-go-lang-3982c235dfa2
- [饶全成: 深度解密 Go 语言之 sync.Pool](https://www.cnblogs.com/qcrao-2018/p/12736031.html)
- [第 14 期 2018-08-17 sync.Pool 源码分析及适用场景](https://talkgo.org/t/topic/33)
- [What's happening in Go tip (2014-01-10)](http://dominik.honnef.co/go-tip/2014-01-10/)
- [为什么容器内存占用居高不下，频频 OOM](https://eddycjy.com/posts/why-container-memory-exceed/)
- [浅谈 Golang sync 包的相关使用方法](https://deepzz.com/post/golang-sync-package-usage.html)
- [fasthttp 快在哪里](https://xargin.com/why-fasthttp-is-fast-and-the-cost-of-it/)
