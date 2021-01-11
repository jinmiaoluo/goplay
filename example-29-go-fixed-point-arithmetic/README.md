# Go 定点算术运算

#### 结构体概述

```go
type Decimal struct {
  value *big.Int
  exp int32
}
```

3.14 会被表示为: 314*10^(-2)

```
value: 314
exp: -2
```
- value 是以整数表示的值.
- exp 表示 10 的 -2 次方.

#### 加法

#### 减法

#### 乘法

#### 除法

#### PPT

这部分使用了 [present](https://github.com/vinayak-mehta/present) 作为 PPT 工具.

1. 安装命令行 PPT 工具及相关的依赖: `make install`.
2. 开始播放 PPT: `make`.

#### 参考
- https://en.wikipedia.org/wiki/Fixed-point_arithmetic
