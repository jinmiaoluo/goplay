原文地址: https://gfw.go101.org/article/unsafe.html

在源码文件中也有很详细的备注信息. 推荐阅读:

https://golang.org/pkg/unsafe/#Pointer

下面的特性是 `unsafe` 包所有功能的基础依据. 需要记住:

- A pointer value of any type can be converted to a Pointer.
- A Pointer can be converted to a pointer value of any type.
- A uintptr can be converted to a Pointer.
- A Pointer can be converted to a uintptr.


- 任何类型的指针值可以被转换为 `unsafe.Pointer` 类型的值
- `unsafe.Pointer` 类型的值可以被转换为任意类型的指针值
- 一个 `uintptr` 类型的值可以转换为 `unsafe.Pointer` 类型的值
- 一个 `unsafe.Pointer` 类型的值可以转换为 `uintptr` 类型的值

这里面放了文档中的 demo 代码. 

一个使用了 `Sizeof` `Offsetof` `Alignof` 三个函数的例子
```shell script
go run demo1.go
```

传递给 `Offsetof` 函数的实参必须为一个字段选择器形式 `value.field`。 此选择器可以表示一个内嵌字段，但此选择器的路径中不能包含指针类型的隐式字段. demo 如下

```shell script
go run demo2.go
```

如何使用 `unsafe.Pointer` 进行值操作

```shell script
go run demo3.go
```

如何使用 `unsafe.Offset()` 和 `uintptr` 进行基于偏移值的结构体成员访问

```shell script
go run demo4.go
```

如何使用 `unsafe.Sizeof()` 和 `uintptr` 进行基于数组成员大小的成员访问

```shell script
go run demo5.go
```

一个使用reflect.StringHeader的例子

```shell script
go run demo6.go
```

一个使用了reflect.SliceHeader的例子

```shell script
go run demo7.go
```

下面是一个展示了如何通过使用非类型安全途径将一个字符串转换为字节切片的例子。 和使用类型安全途径进行转换不同，使用非类型安全途径避免了复制一份底层字节序列。

```shell script
go run demo8.go
```
