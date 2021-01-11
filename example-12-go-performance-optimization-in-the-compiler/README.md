# Go 编译器中的性能优化

这是对

- 文章: https://github.com/go101/go101/wiki/The-perceivable-optimizations-made-by-the-standard-Go-compiler-%28gc%29
- - 视频: https://talkgo.org/t/topic/702
-
- 的 demo 记录和总结

#### 字符串和字节切片之间的转换

```go
go run demo0.go
```

