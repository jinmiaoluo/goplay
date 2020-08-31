# Go init 执行顺序

#### demo

目录树
```
├── a.go
├── b.go
└── d
    └── d.go
```

在当前文件夹执行下面的命令可以查看调用顺序:
```go
go run a.go b.go
```

- 先初始化 package variables
  - a import b, and b import c
    - variables from c
    - variables from b
    - variables from a
  - package variables 类似于全局变量
- 然后调用 init 函数
  - 不同 package 中的 init 函数
    - a import b, and b import c
      - 调用顺序如下:
        - init function from c
        - init function from b
        - init function from a
  - 同一个包内的 init 函数
    - 在同一个文件内
      - 按照定义先后顺序
    - 在不同文件内
      - 按照编译时传入文件的顺序(比如上面的 `go run a.go b.go` 命令, 在 a.go 的 init 函数就会比 b.go 中的 init 函数先调用)
