# CS:APP

CS:APP 第一周 & 第二周研讨会

Powered by [present](https://github.com/vinayak-mehta/present)

---
### 问题

```go
package main

import "fmt"

func main() {
	var a float32
	a = 0.28
	fmt.Printf("%032.23f\n", a)
	// output: 00000000.28000000119209289550781
}
```
---

### 问题

浮动类型在表示 0.1 0.2 等小数时存在精度问题

0.1(10) = 0.00011001100...(2)
0.2(10) = 0.00110011001...(2)

如何精确的计算小数?

---

### 相关的 issue

https://github.com/golang/go/issues/12127

---

### decimal 包

一个解决浮点类型精确度问题的 Golang 包

https://github.com/shopspring/decimal

通过 Go 定点算术运算来避免精度问题

---

### 什么是定点数

通过整数和缩放因子来表示小数

0.2 = 2 * 10^(-1)

---

### 结构体概述


```go
type Decimal struct {
  value *big.Int
  exp int32
}
```

3.14 会被表示为: 314*10^(-2)

```go
value: 314
exp: -2
```

exp 表示 10 的 -2 次方中的 -2

---

### 加法

0.2 + 0.34

初始状态: `2 * 10^(-1) + 34 * 10^(-2)`

缩放差异: `diff = abs((-1) - (-2)) = 1`

缩放差异值: `s = 10^(diff) = 10`

统一缩放:

```
∵  -2 < -1

∴  2 * 10^(-1) = 2 * s * 10^(-1-diff) = 2 * 10 * 10^(-2) = 20 * 10^(-2)

∴  2 * 10^(-1) + 34 * 10^(-2) = 20 * 10^(-2) + 34 * 10^(-2)
                              = (20 + 34) * 10^(-2)
```

---

### 参考
- https://en.wikipedia.org/wiki/Fixed-point_arithmetic

---

### end

<!-- effect=fireworks -->
