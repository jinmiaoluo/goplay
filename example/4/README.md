# 打印调用栈
打印调用栈是调试代码常用的手段

# 目录

<!-- vim-markdown-toc GFM -->

* [打印嵌套函数的调用栈](#打印嵌套函数的调用栈)
* [pprof 打印 goroutine 调用栈](#pprof-打印-goroutine-调用栈)
* [通过 runtime.Stack 打印 goroutine 调用栈](#通过-runtimestack-打印-goroutine-调用栈)
* [通过 SIGNQUIT 信号触发打印调用栈](#通过-signquit-信号触发打印调用栈)
* [通过 panic 和 GOTRACEBACK 打印调用栈](#通过-panic-和-gotraceback-打印调用栈)
* [参考](#参考)

<!-- vim-markdown-toc -->

#### 打印嵌套函数的调用栈

```bash
go run ./demo1.go
```

结果
```bash
Here is function doo
goroutine 1 [running]:
runtime/debug.Stack(0x15, 0x0, 0x0)
        /usr/local/go/src/runtime/debug/stack.go:24 +0x9d
runtime/debug.PrintStack()
        /usr/local/go/src/runtime/debug/stack.go:16 +0x22
main.doo()
        /home/jinmiaoluo/repo/goplay/example/4/main.go:10 +0x7a
main.coo(...)
        /home/jinmiaoluo/repo/goplay/example/4/main.go:14
main.boo(...)
        /home/jinmiaoluo/repo/goplay/example/4/main.go:18
main.main()
        /home/jinmiaoluo/repo/goplay/example/4/main.go:22 +0x21
```

#### pprof 打印 goroutine 调用栈

```bash
go run demo2.go
```

结果
```bash
goroutine profile: total 2
1 @ 0x4b58e5 0x4b5700 0x4b24ca 0x4bcf75 0x4bcefb 0x4bcefc 0x4bcef6 0x433372 0x45e031
#       0x4b58e4        runtime/pprof.writeRuntimeProfile+0x94  /usr/local/go/src/runtime/pprof/pprof.go:694
#       0x4b56ff        runtime/pprof.writeGoroutine+0x9f       /usr/local/go/src/runtime/pprof/pprof.go:656
#       0x4b24c9        runtime/pprof.(*Profile).WriteTo+0x3d9  /usr/local/go/src/runtime/pprof/pprof.go:329
#       0x4bcf74        main.m3+0x64                            /home/jinmiaoluo/repo/goplay/example/4/demo2.go:20
#       0x4bcefa        main.m2+0x3a                            /home/jinmiaoluo/repo/goplay/example/4/demo2.go:17
#       0x4bcefb        main.m1+0x3b                            /home/jinmiaoluo/repo/goplay/example/4/demo2.go:14
#       0x4bcef5        main.main+0x35                          /home/jinmiaoluo/repo/goplay/example/4/demo2.go:11
#       0x433371        runtime.main+0x211                      /usr/local/go/src/runtime/proc.go:203

1 @ 0x4bcfa1 0x45e031
#       0x4bcfa0        main.a+0x0      /home/jinmiaoluo/repo/goplay/example/4/demo2.go:23
```

#### 通过 runtime.Stack 打印 goroutine 调用栈

```bash
go run ./demo3.go
kill -SIGUSR1 <pid>
```

结果
```bash
=== BEGIN goroutine stack dump ===
goroutine 17 [running]:
main.DumpStacks()
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:40 +0x6d
main.setupSigusr1Trap.func1(0xc00005a180)
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:34 +0x34
created by main.setupSigusr1Trap
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:32 +0xab

goroutine 1 [sleep]:
time.Sleep(0x34630b8a000)
        /usr/local/go/src/runtime/time.go:198 +0xba
main.m3(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:24
main.m2(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:21
main.m1(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:18
main.main()
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:15 +0x4f

goroutine 6 [syscall]:
os/signal.signal_recv(0x4e1a20)
        /usr/local/go/src/runtime/sigqueue.go:147 +0x9c
os/signal.loop()
        /usr/local/go/src/os/signal/signal_unix.go:23 +0x22
created by os/signal.Notify.func1
        /usr/local/go/src/os/signal/signal.go:127 +0x44

goroutine 18 [sleep]:
time.Sleep(0x34630b8a000)
        /usr/local/go/src/runtime/time.go:198 +0xba
main.a()
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:27 +0x30
created by main.main
        /home/jinmiaoluo/repo/goplay/example/4/demo3.go:14 +0x3a

=== END goroutine stack dump ===
```

#### 通过 SIGNQUIT 信号触发打印调用栈

```bash
go run demo4.go
kill -SIGQUIT <pid>
```

结果
```bash
SIGQUIT: quit
PC=0x456e90 m=0 sigcode=0

goroutine 0 [idle]:
runtime.epollwait(0x3, 0x7ffed8be9830, 0x36ee7f00000080, 0x0, 0x36ee7f, 0x0, 0x0, 0x0, 0x0, 0x0, ...)
        /usr/local/go/src/runtime/sys_linux_amd64.s:705 +0x20
runtime.netpoll(0x34630b87a4e, 0x1)
        /usr/local/go/src/runtime/netpoll_epoll.go:119 +0x92
runtime.findrunnable(0xc000022000, 0x0)
        /usr/local/go/src/runtime/proc.go:2323 +0x72b
runtime.schedule()
        /usr/local/go/src/runtime/proc.go:2520 +0x2fc
runtime.park_m(0xc000001080)
        /usr/local/go/src/runtime/proc.go:2690 +0x9d
runtime.mcall(0x0)
        /usr/local/go/src/runtime/asm_amd64.s:318 +0x5b

goroutine 1 [sleep]:
time.Sleep(0x34630b8a000)
        /usr/local/go/src/runtime/time.go:198 +0xba
main.m3(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo4.go:18
main.m2(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo4.go:15
main.m1(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo4.go:12
main.main()
        /home/jinmiaoluo/repo/goplay/example/4/demo4.go:9 +0x4a

goroutine 5 [sleep]:
time.Sleep(0x34630b8a000)
        /usr/local/go/src/runtime/time.go:198 +0xba
main.a()
        /home/jinmiaoluo/repo/goplay/example/4/demo4.go:21 +0x30
created by main.main
        /home/jinmiaoluo/repo/goplay/example/4/demo4.go:8 +0x35

rax    0xfffffffffffffffc
rbx    0x36ee7f
rcx    0x456e90
rdx    0x80
rdi    0x3
rsi    0x7ffed8be9830
rbp    0x7ffed8be9e30
rsp    0x7ffed8be97e0
r8     0x0
r9     0xc000064050
r10    0x36ee7f
r11    0x246
r12    0xffffffffffffffff
r13    0xe
r14    0xd
r15    0x200
rip    0x456e90
rflags 0x246
cs     0x33
fs     0x0
gs     0x0
exit status 2
```

#### 通过 panic 和 GOTRACEBACK 打印调用栈

```bash
GOTRACEBACK=1 go run demo5.go
```

结果
```bash
panic: panic from m3

goroutine 1 [running]:
main.m3(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo5.go:18
main.m2(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo5.go:15
main.m1(...)
        /home/jinmiaoluo/repo/goplay/example/4/demo5.go:12
main.main()
        /home/jinmiaoluo/repo/goplay/example/4/demo5.go:9 +0x54

goroutine 5 [sleep]:
time.Sleep(0x34630b8a000)
        /usr/local/go/src/runtime/time.go:198 +0xba
main.a()
        /home/jinmiaoluo/repo/goplay/example/4/demo5.go:21 +0x30
created by main.main
        /home/jinmiaoluo/repo/goplay/example/4/demo5.go:8 +0x35
exit status 2
```

关于 GOTRACEBACK 的一点官方解释:

> The GOTRACEBACK variable controls the amount of output generated when a Go program fails due to an unrecovered panic or an unexpected runtime condition. By default, a failure prints a stack trace for the current goroutine, eliding functions internal to the run-time system, and then exits with exit code 2. The failure prints stack traces for all goroutines if there is no current goroutine or the failure is internal to the run-time. GOTRACEBACK=none omits the goroutine stack traces entirely. GOTRACEBACK=single (the default) behaves as described above. GOTRACEBACK=all adds stack traces for all user-created goroutines. GOTRACEBACK=system is like “all” but adds stack frames for run-time functions and shows goroutines created internally by the run-time. GOTRACEBACK=crash is like “system” but crashes in an operating system-specific manner instead of exiting. For example, on Unix systems, the crash raises SIGABRT to trigger a core dump. For historical reasons, the GOTRACEBACK settings 0, 1, and 2 are synonyms for none, all, and system, respectively. The runtime/debug package's SetTraceback function allows increasing the amount of output at run time, but it cannot reduce the amount below that specified by the environment variable

#### 参考
- [鸟窝-调试利器：dump goroutine 的 stacktrace](https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/)
