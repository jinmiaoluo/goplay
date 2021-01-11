# 单例模式

> 单例模式，也叫单子模式，是一种常用的软件设计模式，属于创建型模式的一种。在应用这个模式时，单例对象的类必须保证只有一个实例存在 - 维基百科

#### 解决的问题

- 确保一个类只有一个实例.
- 在 Golang 内对应的是包级别的结构体类型变量如何初始化的问题.

#### 如何实现

- 方案1
  - 在定义了私有的包级别的结构体类型变量后, 立即通过 init() 函数对该变量进行初始化.
  - 提供一个 GetInstance() 函数返回这个私有的包级别的结构体变量的指针.
- 方案2
  - 定义私有的包级别的结构体类型变量.
  - 提供一个 GetInstance() 函数返回这个私有的包级别的结构体变量的指针. 这个函数不同于饿汉模式的函数, 在这个函数内判断私有的包级别的结构体类型变量是否初始化, 如果没有初始化, 进行初始化, (已经初始化过了就什么也不做)直接返回这个私有的包级别的结构体变量的指针.

#### 如何使用

单例模式下会有一个类似 GetInstance() 的公开函数用于在其他函数内访问这个唯一的某个类型的实例

#### 构建方式
- 懒汉方式: 指全局的单例实例在第一次被使用时构建.(即在调用 GetInstance() 时)
- 饿汉方式: 指全局的单例实例在类装载时构建 (在 init 阶段初始化该私有的包级别的变量).

#### 参考
- https://en.wikipedia.org/wiki/Singleton_pattern
- https://zh.wikipedia.org/wiki/%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F
- https://lailin.xyz/post/singleton.html
- https://wiki.jikexueyuan.com/project/java-design-pattern/singleton-pattern.html
- https://www.bilibili.com/video/BV16g4y1q7xi
