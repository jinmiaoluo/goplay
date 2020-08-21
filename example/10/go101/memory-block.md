- 原文地址: https://github.com/go101/go101/blob/836588903a76be49da373349e4b69d523d2a0e3b/articles/memory-block.html
- 作者: go101 项目组
- 译者: Jinmiao Luo

# Memory Blocks 内存块

Go is a language which supports automatic memory management, such as automatic memory allocation and automatic garbage collection. So Go programmers can do programming without handling the underlying verbose memory management. This not only brings much convenience and saves Go programmers lots of time, but also helps Go programmers avoid many careless bugs.

Although knowing the underlying memory management implementation details is not necessary for Go programmers to write Go code, understanding some concepts and being aware of some facts in the memory management implementation by the standard Go compiler and runtime is very helpful for Go programmers to write high quality Go code.

This article will explain some concepts and list some facts of the implementation of memory block allocation and garbage collection by the standard Go compiler and runtime. Other aspects, such as memory apply and memory release in memory management, will not be touched in this article.

Go 是一种支持自动内存管理的语言，例如自动内存分配和自动垃圾收集。因此 Go 程序员可以在不进行底层详细内存管理的情况下进行编程。这不仅带来了很多便利，并节省了 Go 程序员大量的时间，而且还帮助 Go 程序员避免了许多粗心导致的错误。

尽管了解底层的内存管理实现细节对于 Go 程序员编写 Go 代码不是必需的，但是了解标准的 Go 编译器和运行时的一些概念并了解内存管理实现中的一些事实对于 Go 程序员编写高质量的代码非常有帮助。

本文将解释一些概念，并列出通过标准 Go 编译器和运行时实现内存块分配和垃圾收集的一些事实。本文将不涉及其他方面，例如内存管理中的内存申请和内存释放。

### Memory Blocks 内存块

A memory block is a continuous memory segment to host value parts at run time. Different memory blocks may have different sizes, to host different value parts. One memory block may host multiple value parts at the same time, but each value part can only be hosted within one memory block, no matter how large the size of that value part is. In other words, for any value part, it never crosses memory blocks.

There are many reasons when one memory block may host multiple value parts. Some of them:

- a struct value often have several fields. So when a memory block is allocated for a struct value, the memory block will also host (the direct parts of) these field values.
- an array values often have many elements. So when a memory block is allocated for a array value, the memory block will also host (the direct parts of) the array element values.
- the underlying element sequences of two slices may be hosted on the same memory block, the two element sequences even can overlap with each other.

<!--todo: 这里应该替换为值部分译文-->
内存块是一个连续的内存段，用于在运行时托管[值部分](https://go101.org/article/value-part.html)。不同的内存块可能具有不同的大小，以承载不同的值部分。一个内存块可以同时托管多个值部分，但是每个值部分只能托管在一个内存块中，无论该值部分的大小如何。换句话说，对于任何值部分，它都不能跨多个内存块。

一个内存块可能包含多个值部分, 有很多原因, 其中的几个：

- 结构值通常具有多个字段。因此，当为结构值分配内存块时，该内存块还将托管这些字段值（的直接部分）。
- 数组值通常包含许多元素。因此，当为数组值分配内存块时，该内存块还将托管数组元素值（的直接部分）。
- 两个 slice 的基础元素序列可能位于同一内存块上，这两个元素序列甚至可以相互重叠。

### A Value References the Memory Blocks Which Host Its Value Parts 值引用内存块(承载其值部分)

We have known that a value part can reference another value part. Here, we extend the reference definition by saying a memory block is referenced by all the value parts it hosts. So if a value part v is referenced by another value part, then the other value will also reference the memory block hosting v, indirectly.

我们知道一个值部分可以引用另一个值部分。在这里，我们扩展引用定义为, 内存块被其托管的所有值部分引用。因此，如果一个值部分 `v` 被另一个值部分引用，则另一个值也将间接引用托管 `v` 的内存块。

### When Will Memory Blocks Be Allocated? 何时分配内存块？

In Go, memory blocks may be allocated but not limited at following situations:

- explicitly call the new and make built-in functions. A new call will always allocate exact one memory block. A make call will allocate more than one memory blocks to host the direct part and underlying part(s) of the created slice, map or channel value.
- create maps, slices and anonymous functions with corresponding literals. More than one memory blocks may be allocated in each of the processes.
- declare variables.
- assign non-interface values to interface values (when the non-interface value is not a pointer value).
- concatenate non-constant strings.
- convert strings to byte or rune slices, and vice versa, except some special compiler optimization cases.
- convert integers to strings.
- call the built-in append function (when the capacity of the base slice is not large enough).
- add a new key-element entry pair into a map (when the underlying hash table needs to be resized).

在Go中，在下面的情况下将分配内存块，但不限于以下情况：

- 显式调用 `new` 和 `make` 内置函数。一个 `new` 调用将始终分配确切的一个内存块。一个 `make` 调用将分配多个内存块，以承载创建的 slice，map 或通道值的直接部分和基础部分。
- 创建具有相应字面量的 map，slice 和匿名函数。在每个过程中会分配一个以上的内存块。
- 声明变量。
- 将非接口值分配给接口值（当非接口值不是指针值时）。
- 连接非常量字符串。
<!--todo: 这里应该替换为字符串#转换优化一章译文-->
- 将字符串转换为字节片或符文片，反之亦然，除了[一些特殊的编译器优化情况](https://go101.org/article/string.html#conversion-optimizations)。
- 将整数转换为字符串。
- 调用内置的 `append` 函数（当底层的 slice 的容量不够大时）。
- 在 map 中添加一个新的键元素条目对（当需要调整基础哈希表的大小时）。

### Where Will Memory Blocks Be Allocated On? 内存块将分配到哪里？

For every Go program compiled by the official standard Go compiler, at run time, each goroutine will maintain a stack, which is a memory segment. It acts as a memory pool for some memory blocks to be allocated from/on. The initial stack size of each goroutine is small (about 2k bytes on 64-bit systems). The stack size will grow and shrink as needed in goroutine running.

(Please note, for the standard Go compiler, there is a limit of stack size each goroutine can have. For standard Go compiler 1.11, the default maximum stack size is 1 GB on 64-bit systems, and 250 MB on 32-bit systems. We can call the SetMaxStack function in the runtime/debug standard package to change the size.)

对于由官方标准 Go 编译器编译的每个 Go 程序，在运行时，每个 goroutine 将维护一个栈，这是一个内存段。它充当一些内存块的内存池，以便从中分配内存。每个 goroutine 的初始栈大小很小（在 64 位系统上约为 2k 字节）。栈大小将在 goroutine 运行中根据需要增大和缩小。

（请注意，对于标准 Go 编译器，每个 goroutine 都有限制的栈大小。对于标准 Go 编译器 1.11 版本，默认最大栈大小在 64 位系统上为 1 GB，在 32 位系统上为 250 MB 我们可以在 `runtime/debug` 标准包中调用 `SetMaxStack` 函数来更改大小。）

Memory blocks can be allocated on stacks. Memory blocks allocated on the stack of a goroutine can only be used (referenced) in the goroutine internally. They are goroutine localized resources. They are not safe to be referenced crossing goroutines. A goroutine can access or modify the value parts hosted on a memory block allocated on the stack of the goroutine without using any data synchronization techniques.

Heap is a singleton in each program. It is a virtual concept. If a memory block is not allocated on any goroutine stack, then we say the memory block is allocated on heap. Value parts hosted on memory blocks allocated on heap can be used by multiple goroutines. In other words, they can be used concurrently. Their uses should be synchronized when needed.

内存块可以被分配在栈上。在 goroutine 栈上分配的内存块只能在 goroutine 内部使用（引用）。它们是 goroutine 本地化资源。跨 goroutine 引用内存块是不安全的(译者注: 比如 A goroutine 应用 B goroutine 中的内存块)。 goroutine 可以访问或修改托管在 goroutine 栈上分配的内存块上的值部分，而无需使用任何数据同步技术。

堆是每个程序中的一个单例。这是一个虚拟的概念。如果一个内存块没有被分配在任何一个 goroutine 的栈上，那么我们说该内存块分配在堆上。托管在分配给堆的内存块上的值部分可以被多个 goroutine 使用。换句话说，它们可以同时使用。使用堆时应该按照需求进行同步。

Heap is a conservative place to allocate memory blocks on. If compilers detect a memory block will be referenced crossing goroutines or can't easily confirm whether or not the memory block is safe to be put on the stack of a goroutine, then the memory block will be allocated on heap at run time. This means some values can be safely allocated on stacks may be also allocated on heap.

In fact, stacks are not essential for Go programs. Go compiler/runtime can allocate all memory block on heap. Supporting stacks is just to make Go programs run more efficiently:

- allocating memory blocks on stacks is much faster than on heap.
- memory blocks allocated on a stack don't need to be garbage collected.
- stack memory blocks are more CPU cache friendly than heap ones.

堆是一个保守的地方，可以在上面分配内存块。如果编译器检测到某个内存块将被多个 go​​routine 引用，或者无法轻松地确认该内存块放置在 goroutine 的栈中是否安全时，则该内存块将在运行时被分配在堆上。这意味着一些值可以安全地分配在栈上，也还是被分配在堆上。

实际上，栈对于 Go 程序不是必需的。 Go 编译器/运行时可以在堆上分配所有内存块。支持栈只是为了使Go程序更有效地运行：

- 在栈上分配内存块比在堆上分配快得多。
- 在栈上分配的内存块不需要进行垃圾收集。
- 栈内存块比堆内存块对CPU缓存更友好。

If a memory block is allocated somewhere, we can also say the value parts hosted on the memory block are allocated on the same place.

If some value parts of a local variable declared in a function is allocated on heap, we can say the value parts (and the variable) escape to heap. By using Go Toolchain, we can run `go build -gcflags -m` to check which local values (value parts) will escape to heap at run time. As mentioned above, the current escape analyzer in the standard Go compiler is still not perfect, many local value parts can be allocated on stacks safely will still escape to heap.

An active value part allocated on heap still in use must be referenced by at least one value part allocated on a stack. If a value escaping to heap is a declared local variable, and assume its type is T, Go runtime will create (a memory block for) an implicit pointer of type *T on the stack of the current goroutine. The value of the pointer stores the address of the memory block allocated for the variable on heap (a.k.a., the address of the local variable of type T). Go compiler will also replace all uses of the variable with the dereferences of the pointer value at compile time. The *T pointer value on stack may be marked as dead since a later time, so the reference relation from it to the T value on heap will disappear. The reference relation from the *T value on stack to the T value on heap plays an important role in the garbage collection process which will be described below.

如果在某处分配了内存块，我们也可以说存储在内存块中的值部分分配在同一位置。

如果在函数中声明的局部变量的某些值部分分配在堆上，则可以说值部分（和变量）逃逸到堆上。通过使用 Go 工具链，我们可以运行 `go build -gcflags -m` 来检查哪些本地值（值部分）将在运行时逃逸到堆上。如上所述，标准 Go 编译器中的逃逸分析器仍然不够完善，许多局部值可以安全地分配在栈上但仍逃逸到堆上。

仍在使用的堆上分配的存活的值部分, 必须至少被一个在栈上分配的值部分所引用。如果逃逸到到堆的值是被声明为局部变量，假定其类型为 `T`, Go 运行时将在当前 goroutine 的栈上创建 `*T` 类型的隐式指针（作为其内存块）。指针的值存储为堆上的变量分配的内存块的地址（也就是类型 `T` 的局部变量的地址）。 Go 编译器还将在编译时用指针值解引用(得到的值)替换所有使用的变量。之后，栈上的 `*T` 指针值会被标记为无效，因此从它到堆上 `T` 值的引用关系将消失。从栈上的 `*T` 值到堆上的 `T` 值的引用关系在垃圾收集过程中起着重要作用，下面将对此进行描述。

Similarly, we can view each package-level variable is allocated on heap, and the variable is referenced by an implicit pointer which is allocated on a global memory zone. In fact, the implicit pointer references the direct part of the package-level variable, and the direct part of the variable references some other value parts.

A memory block allocated on heap may be referenced by multiple value parts allocated on different stacks at the same time.

Some facts:
- if a field of a struct value escapes to heap, then the whole struct value will also escape to heap.
- if an element of an array value escapes to heap, then the whole array value will also escape to heap.
- if an element of a slice value escapes to heap, then all the elements of the slice will also escape to heap.
- if a value (part) v is referenced by a value (part) which escapes to heap, then the value (part) v will also escape to heap.
- A memory block created by calling new function may be allocated on heap or stacks. This is different to C++.

同样，我们可以看到每个程序包级别的变量在堆上分配，并且该变量被全局内存区域上分配的隐式指针引用。实际上，隐式指针引用了程序包级变量的直接部分，而变量的直接部分则引用了其他一些值部分。

分配在堆上的内存块可以被同时被分配在不同栈上的多个值部分引用。

一些事实：
- 如果一个 struct 值的字段逃逸到堆上，那么整个 struct 值也将逃逸到堆上。
- 如果数组值的一个成员逃逸到堆上，那么整个数组值也将逃逸到堆上。
- 如果 slice 值的成员逃逸到堆上，那么 slice 的所有成员也将逃逸到堆上。
- 如果值（部分）`v` 被逃逸到堆的值（部分）引用，则值（部分）`v` 也将逃逸到堆上。
- 通过调用新函数创建的内存块可以分配在堆或栈上。这与 C++ 不同。

When the size of a goroutine stack changes, a new memory segment will be allocated for the stack. So the memory blocks allocated on the stack will very likely be moved, or their addresses will change. Consequently, the pointers, which must be also allocated on the stack, referencing these memory blocks also need to be modified accordingly.

当 goroutine 栈的大小更改时，将为该栈分配一个新的内存段。因此，分配在栈上的内存块很可能会移动，换句话说, 它们的地址将会更改。因此，引用栈里面的内存块的指针必须分配在栈上, 并且指针值必须得到相应地修改(译者注: 如果出现前面提到的 goroutine 栈大小发生改变的情况)。

### When Can a Memory Block Be Collected? 内存块何时会被收集?
Memory blocks allocated for direct parts of package-level variables will never be collected.

The stack of a goroutine will be collected as a whole when the goroutine exits. So there is no need to collect the memory blocks allocated on stacks, individually, one by one. Stacks are not collected by the garbage collector.

For a memory block allocated on heap, it can be safely collected only if it is no longer referenced (either directly or indirectly) by all the value parts allocated on goroutine stacks and the global memory zone. We call such memory blocks as unused memory blocks. Unused memory blocks on heap will be collected by the garbage collector.

Here is an example to show when some memory blocks can be collected:

分配给包级变量的直接部分的内存块将永远不会被收集。

当 goroutine 退出时，goroutine 的堆栈将作为一个整体收集。因此，无需一一收集在栈上分配的内存块。栈不是由垃圾收集器收集的。

对于在堆上分配的内存块，只有在 goroutine 栈和全局内存区域上分配的所有值部分都不再（直接或间接）引用它时，才能安全地收集它。我们称这类内存块为未使用的内存块。垃圾收集器将收集堆上未使用的内存块。

这是显示何时可以收集一些内存块的示例：

```go
package main

var p *int

func main() {
	done := make(chan bool)
	// "done" will be used in main and the following
	// new goroutine, so it will be allocated on heap.

	go func() {
		x, y, z := 123, 456, 789
		_ = z  // z can be allocated on stack safely.
		p = &x // For x and y are both ever referenced
		p = &y // by the global p, so they will be both
		       // allocated on heap.

		// Now, x is not referenced by anyone, so
		// its memory block can be collected now.

		p = nil
		// Now, y is als not referenced by anyone,
		// so its memory block can be collected now.

		done <- true
	}()

	<-done
	// Now the above goroutine exits, the done channel
	// is not used any more, a smart compiler may
	// think it can be collected now.

	// ...
}
```

Sometimes, smart compilers, such as the standard Go compiler, may make some optimizations so that some references are removed earlier than we expect. Here is such an example.

有时，诸如标准 Go 编译器之类的智能编译器可能会进行一些优化，以使某些引用的删除时间比我们预期的要早。这是一个例子。

```go
package main

import "fmt"

func main() {
	// Assume the length of the slice is so large
	// that its elements must be allocated on heap.
	bs := make([]byte, 1 << 31)

	// A smart compiler can detect that the
	// underlying part of the slice bs will never be
	// used later, so that the underlying part of the
	// slice bs can be garbage collected safely now.
	fmt.Println(len(bs))
}
```

Please read value parts to learn the internal structures of slice values.

By the way, sometimes, we may hope the slice bs is guaranteed to not being garbage collected until `fmt.Println` is called, then we can use a `runtime.KeepAlive` function call to tell garbage collectors that the slice bs and the value parts referenced by it are still in use.

For example,

<!--todo:这里应该替换为值部分译文地址-->
请阅读[值部分](https://go101.org/article/value-part.html)以了解切片值的内部结构。

顺便说一句，有时候我们希望可以确保在调用 `fmt.Println` 之前 slice bs 不会被垃圾收集，我们可以调用 `runtime.KeepAlive` 函数来告诉垃圾收集器 slice bs 和被 bs 引用的值部分仍在使用中。

例如，

```go
package main

import "fmt"
import "runtime"

func main() {
	bs := make([]int, 1000000)

	fmt.Println(len(bs))

	// A runtime.KeepAlive(bs) call is also
	// okay for this specified example.
	runtime.KeepAlive(&bs)
}
```

`runtime.KeepAlive` function calls are often needed if unsafe pointers are involved.

<!--todo: 这里应该替换为类型不安全的指针译文地址-->
如果涉及到 [unsafe pointers](https://go101.org/article/unsafe.html)，通常需要运行时间 `run.KeepAlive` 函数。

### How Are Unused Memory Blocks Detected? 如何检测未使用的内存块？
The current standard Go compiler (version 1.15) uses a concurrent, tri-color, mark-sweep garbage collector. Here this article will only make a simple explanation for the algorithm.

A garbage collection (GC) process is divided into two phases, the mark phase and the sweep phase. In the mark phase, the collector (a group of goroutines actually) uses the tri-color algorithm to analyze which memory blocks are unused.

当前的标准 Go 编译器（版本 1.15）使用并发的三色(译者注: 一个算法, 见参考)标记清除(译者注: 一个算法, 见参考)垃圾收集器。在这里，本文仅对算法进行简单说明。

垃圾收集（GC）过程分为两个阶段，即标记阶段和清除阶段。在标记阶段，收集器（实际上是一组 goroutine）使用三色算法分析未使用的内存块。

The following quote is token from a Go blog article, in which an objects is either value parts or memory blocks.
> At the start of a GC cycle all objects are white. The GC visits all roots, which are objects directly accessible by the application such as globals and things on the stack, and colors these grey. The GC then chooses a grey object, blackens it, and then scans it for pointers to other objects. When this scan finds a pointer to a white object, it turns that object grey. This process repeats until there are no more grey objects. At this point, white objects are known to be unreachable and can be reused.

下面的引文来自 [Go 博客文章](https://blog.golang.org/go15gc)，其中**对象**是值部分或内存块。
> 在 GC 周期开始时，所有对象均为白色。 GC访问所有根，这些根是应用程序可直接访问的对象，例如全局变量和栈中的东西，这些东西会被标记为灰色。然后，GC 选择一个灰色对象将其标记为黑色，然后对其进行扫描以寻找指向其他对象的指针。当此扫描找到指向白色对象的指针时，指针指向的白色对象将变为灰色。重复此过程，直到没有其他灰色对象为止。在这一点上，已知白色物体是无法触及的并且可以被复用(译者注: 可以被收集)。(译者注: 这里的译文可能有误)
<!--todo: 确保译文关于垃圾收集器三色算法解释的正确性-->

About why the algorithm uses three colors instead of two colors, please search "write barrier golang" for details. Here only provides two references: eliminate STW stack re-scanning and mbarrier.go.

In the sweep phase, the marked unused memory blocks will be collected.

The GC algorithm is a non-compacting one, so it will not move memory blocks to rearrange them.

关于该算法为何使用三种颜色而不是两种颜色的原因，请搜索 "write barrier golang" 以获取详细信息。这里仅提供两个参考：[eliminate STW stack re-scanning](https://github.com/golang/proposal/blob/master/design/17503-eliminate-rescan.md) 和 [mbarrier.go](https://golang.org/src/runtime/mbarrier.go)

在清除阶段，将收集标记的未使用的内存块。

GC算法是非紧凑算法，因此它不会移动内存块来重新排列它们。

### When Will an Unused Memory Block Be Collected? 什么时候会收集未使用的内存块？
Unused heap memory blocks are viewed as garbage by Go runtime and will be collected to reuse or release memory. The garbage collector is not always running. It will start when a threshold is satisfied. So an unused memory block may be not collected immediately when it becomes unused. Instead, it will be collected eventually. Currently (Go Toolchain 1.15), the threshold is controlled by GOGC environment variable:

> The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely.

The value of this environment variable determines the frequency of garbage collecting, and it can be modified at run time by calling runtime/debug.SetGCPercent

Go运行时将未使用的堆内存块视为垃圾，(这些内存块)将被收集以重用或释放内存。垃圾收集器并不是一直运行的。当满足阈值时它才会启动。因此，未使用的内存块在变为未使用时不一定会被立即收集。取而代之的结果是确保最终被收集。当前（Go 工具链 1.15），阈值由[GOGC环境变量](https://golang.org/pkg/runtime/#hdr-Environment_Variables)控制：

> GOGC变量设置初始垃圾收集目标百分比。当新分配的数据与上一个收集之后剩余的实时数据之比达到此百分比时，将触发收集。默认值为GOGC=100。设置GOGC=off将完全禁用垃圾收集器。

此环境变量的值确定垃圾收集的频率，可以在运行时通过调用 [runtime/debug.SetGCPercent](https://golang.org/pkg/runtime/debug/#SetGCPercent) 函数对其进行修改。较小的值导致更频繁的垃圾收集。负百分比将禁用自动垃圾收集。

A garbage collection process can also be started manually by calling the runtime.GC function.

One more thing needs to note is, for the current official Go runtime (v1.15), a new garbage collection process will start automatically if garbage collection has not run for two minutes.

The gargage collection strategies might change in later official Go runtime versions.

An unused memory block may not be released to OS immediately after it is collected, so that it can be reused for new some value parts. Don't worry, the official Go runtime is much less memory greedy than most Java runtimes.

垃圾收集过程可以通过调用 [runtime.GC](https://golang.org/pkg/runtime/#GC) 函数来手动启动。

需要注意的另一件事是，对于当前的官方Go运行时（v1.15），[新的垃圾收集过程将自动启动, 如果垃圾收集未运行两分钟](https://github.com/golang/go/blob/895b7c85addfffe19b66d8ca71c31799d6e55990/src/runtime/proc.go#L4481-L4486)。

垃圾收集策略可能会在后来的官方 Go 运行时版本中更改。

未使用的内存块在收集之后可能不会立即释放到 OS，因此可以将其重新用于某些值部分。不用担心，官方的 Go 运行时比大多数 Java 运行时在内存贪婪上要少很多。
- 参考
    - https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/
