**TestMain 的用法**

TestMain 在测试开始前和测试结束后可以分别执行代码用于初始化测试的环境

**示例**
```bash
go test -v .
```

**注意事项**
1. 每个 package 只能有一个 TestMain 函数
2. TestMain 不是用于测试 Main 函数的
