# Go 与算法

#### 目录

<!-- vim-markdown-toc GFM -->

* [Go 与排序算法](#go-与排序算法)
  * [shell 排序](#shell-排序)
  * [堆排序](#堆排序)
  * [快速排序](#快速排序)
  * [归并排序](#归并排序)
  * [选择排序](#选择排序)
  * [冒泡排序](#冒泡排序)
  * [插入排序](#插入排序)
  * [选择排序](#选择排序-1)
* [Go 与分词算法](#go-与分词算法)
* [eisel-lemire parsefloat algorithm](#eisel-lemire-parsefloat-algorithm)
* [Go 与限速算法](#go-与限速算法)
* [参考](#参考)

<!-- vim-markdown-toc -->

#### Go 与排序算法

##### shell 排序
##### 堆排序
##### 快速排序

词汇汇总:

- pivot: 基准
- list: 序列
- quicksort: 快速排序

快速排序的流程:


参考文档:

1. [wikipedia: quicksort](https://en.wikipedia.org/wiki/Quicksort)
2. [golang program for implementation of quick sort](https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html)


##### 归并排序
##### 选择排序
##### 冒泡排序
##### 插入排序
##### 选择排序

#### Go 与分词算法

- 暴力匹配.
- kmp 算法.
- bm 算法.
- sunday 算法.
- rabin karp 算法.

#### eisel-lemire parsefloat algorithm

参考资料:
- https://nigeltao.github.io/blog/2020/eisel-lemire.html

#### Go 与限速算法

- 漏桶算法.
- 令牌桶算法.

参考资料:

- https://gocn.vip/topics/11108
- https://github.com/hpcloud/tail/blob/master/ratelimiter/leakybucket.go
- https://en.wikipedia.org/wiki/Leaky_bucket
- https://en.wikipedia.org/wiki/Token_bucket

漏桶算法中的术语:

- leak rate: 泄漏速率
- average rate: 平均速率
- bucket capacity: 桶容量
- counter: 计数器
- burstiness: 一个阈值. 如果超过这个阈值. 将取消用户发来的包
- generic cell rate algorithm(abbr. GCRA): 基于漏桶算法的调度算法. 常用于网络调度

一些 tips:

- 漏桶算法不一定用于限速. 也可以用于事件触发机制. 比如报警中的阈值机制.

#### 参考
- https://www.liwenzhou.com/posts/Go/go_algorithm/
- https://www.liwenzhou.com/posts/Go/go_algorithm/
