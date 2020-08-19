# gRPC 和 Protocol Buffers 3

触发我去了解 gRPC 和 Protocol Buffers 3 的契机是来自 [Jiajun Huang](https://jiajunhuang.com/aboutme) 的一篇文章: [自己动手写一个反向代理](https://jiajunhuang.com/articles/2019_06_09-write_you_a_simple_reverse_proxy.md.html). Demo 的仓库内用了 gRPC 和 Protocol Buffers 3. [相关代码仓库](https://github.com/jiajunhuang/natrp)

#### 开始前的准备

- 需要安装 protoc (Protocol Buffers Compiler)
- `.proto` 文件(文件后缀)的编译成 Go 代码需要安装 `protoc-gen-go` 插件
- 克隆 [相关代码仓库](https://github.com/jiajunhuang/natrp)

#### 如何编译 `.proto` 的 Go 代码

基本格式是:
```bash
protoc --go_out=<存储Go代码的目录> <.proto 文件路径>
```

如果你克隆了 natrp 这个项目的代码, 这个编译的命令记录在 Makefile 内. 内容如下:

```make
default:
	protoc -I pb/ serverpb/server.proto --go_out=plugins=grpc:pb/
	go build -o bin/client client/main.go
	go build -o bin/server server/main.go server/server.go
```

#### 关于 protoc Golang 代码生成插件

截止 `2020-08-06 11:34`. 该插件正在从旧版本到新版本的过渡中, 这意味着在不同的文档里会有两种不同的 `.proto` 文件编译命令, 比如

旧版本编译命令
```bash
protoc -I pb/ serverpb/server.proto --go_out=plugins=grpc:pb/
```

新版本编译命令
```bash
protoc -I pb/ --go_out=./ --go-grpc_out= serverpb/server.proto
```

#### Protocol Buffer 3 的作用

类似与 `XML` 的应用间数据交换格式. 相对于 `XML` 有空间占用小. 性能高的有点. 基于二进制的格式. 是 Interface Definition Language (IDL: 接口定义语言), 即定义接口用的

Protocol Buffer 3 官方仓库有一个 [example](https://github.com/protocolbuffers/protobuf/tree/master/examples) 文件夹, 里面有 Golang 的演示代码. 结合官方的[ 文档 ](https://developers.google.com/protocol-buffers/docs/gotutorial). 是一个不错的开始

#### Protocol Buffers 3 一些基本用法

为了便于理解. Golang 代码经过了精简. 这里的 `.proto` 代码来自 [example](https://github.com/protocolbuffers/protobuf/tree/c6493970296fa5c5b4a81a37248a328579fe9662/examples)

##### `message` 字段
```proto
message Person {
  string name = 1;
  int32 id = 2;
  string email = 3;
}
```

对应的 Golang 数据代码
```go
type Person struct {
  name  string
  id    int32
  email string
}
```

##### `enum` 字段
```proto
enum PhoneType {
  MOBILE = 0;
  HOME = 1;
  WORK = 2;
}
```

对应的 Golang 代码
```go
type PhoneType int32

const (
  MOBILE PhoneType = 0
  HOME PhoneType   = 1
  WORK PhoneType   = 2

)
```

##### `package` 字段
```proto
package tutorial;
```

对应的 Golang 代码
```go
package tutorial
```

`package` 字段还用于 `import` 字段载入其他 `.proto` 文件时避免名字冲突. 相关的解释见后面的 `import` 字段

##### `import` 字段
```proto
import "google/protobuf/timestamp.proto";
message Person {
  google.protobuf.Timestamp last_updated = 5;
}
```

对应的 Golang 代码
```go
import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)
type Person struct {
	LastUpdated *timestamp.Timestamp
}
```

`.proto` 中的 `import` 命令会跟 Golang 中的 `import` 命令做关联. 在 `timestamp.proto` 包中有:
```proto
package google.protobuf;`
option go_package = "github.com/golang/protobuf/ptypes/timestamp";
```

上面的两个字段保证了在生成的 Go 代码中 `import` 指定的目的地址是正确的指向包所在地址的. 并且包不同 `.proto` 文件同名的 `message` 或者其他指令不会冲突. 举个例子: 在 `import` 字段相关的 `.proto` 代码片段中, `last_updated` 的类型是 `google.protobuf.Timestamp`, 即 `<package>.<message-name>` 的形式. 如果自己刚好也有一个 `Timestamp` 类型, 由于 `<package>` 的存在, 从而避免冲突

##### 类型内嵌. 这里以 `message` 类型内嵌为例
```proto
message Person {
  string name = 1;
  int32 id = 2;
  string email = 3;

  message PhoneNumber {
    string number = 1;
  }
}
```

对应的 Golang 代码
```go
type Person struct {
  name  string
  id    int32
  email string
}

type Person_PhoneNumber struct {
  Number string
}
```

##### `repeated` 字段
```proto
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
  }

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;
}
```

对应的 Golang 代码
```go
type Person struct {
  name   string
  id     int32
  email  string
"
  phones []*Person_PhoneNumber
}

type Person_PhoneType int32
const (
  Person_MOBILE Person_PhoneType = 0
  Person_HOME   Person_PhoneType = 1
  Person_WORK   Person_PhoneType = 2
)

type Person_PhoneNumber struct {
  Number string
  Type   Person_PhoneType
}
```

从上面的代码可以看出, `repeated` 字段用于构建数组. 有一个细节需要注意. 如果 `repeated` 对应的类型是 `message`. 因为 `message` 对应 Golang 中的 `struct` 类型. 因此. 对应的数组是指向结构体的指针数组`[]*Person_PhoneNumber` 而不是 `[]Person_PhoneNumber`

##### `service` 字段

[example](https://github.com/protocolbuffers/protobuf/tree/c6493970296fa5c5b4a81a37248a328579fe9662/examples) 中没有用到 `service`, 因此我们以 [natrp/pb/serverpb/server.proto](https://github.com/jiajunhuang/natrp/blob/master/pb/serverpb/server.proto) 为例, 并结合[官方文档](https://developers.google.com/protocol-buffers/docs/proto3#services) 和 [gRPC tutorial](https://grpc.io/docs/languages/go/basics/) 来看看怎么用

关于 `service` 的概念, 可以参考 [gRPC core concept](https://grpc.io/docs/what-is-grpc/core-concepts/)

gRPC 的 `service` 有四种类型. 分别是:

客户端发送一个请求给服务端, 服务端返回一个响应
```proto
rpc SayHello(HelloRequest) returns (HelloResponse);
```

客户端发送一个请求给服务端, 并得到一个数据流, 可以读取一系列的返回消息
```proto
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
```

客户端以 stream 的形式发送一系列的消息给服务端, 并等待服务端读取这些信息并(在结束后)返回响应消息
```proto
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
```

客户端和服务端以一个 read-write 的 stream 的形式发送和接受一系列的消息. 这里, 读写的秩序是由开发者自己控制的. 这一段的解释参考:

> Bidirectional streaming RPCs where both sides send a sequence of messages using a read-write stream. The two streams operate independently, so clients and servers can read and write in whatever order they like: for example, the server could wait to receive all the client messages before writing its responses, or it could alternately read a message then write a message, or some other combination of reads and writes. The order of messages in each stream is preserved.

```proto
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```

当我们定义一个 service, gRPC 会生成相关的 interface. 比如下面的 service:
```proto
service RouteGuide {
  rpc GetFeature(Point) returns (Feature) {}
  rpc ListFeatures(Rectangle) returns (stream Feature) {}
  rpc RecordRoute(stream Point) returns (RouteSummary) {}
  rpc RouteChat(stream RouteNote) returns (stream RouteNote) {}
}
```

会生成如下的 interface, 服务端需要实现的 interface 默认名字是 `service` 的名字加上 `Server` 后缀, 客户端需要实现的 interface 默认名字是 `service` 的名字加上 `Client` 后缀:
```go
// 服务端需要实现的 interface
type RouteGuideServer interface {
	GetFeature(context.Context, *Point) (*Feature, error)
	ListFeatures(*Rectangle, RouteGuide_ListFeaturesServer) error
	RecordRoute(RouteGuide_RecordRouteServer) error
	RouteChat(RouteGuide_RouteChatServer) error
	mustEmbedUnimplementedRouteGuideServer()
}
```

```go
// 客户端需要实现的 interface
type RouteGuideClient interface {
	GetFeature(ctx context.Context, in *Point, opts ...grpc.CallOption) (*Feature, error)
	ListFeatures(ctx context.Context, in *Rectangle, opts ...grpc.CallOption) (RouteGuide_ListFeaturesClient, error)
	RecordRoute(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RecordRouteClient, error)
	RouteChat(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_RouteChatClient, error)
}
```

其中. 生成的代码默认上已经实现了客户端的上述方法. 因此我们只需要根据接口调用指定的方法即可. 服务端的方法需要我们自己去实现, 但上述接口的方法中涉及到的, 比如: `RouteGuide_RecordRouteClient` 接口对应的结构体, gRPC 都已经自动实现了, 这包括了客户端流, 服务端流, 双向流的方法. 所以, 我们只需要调用对应的模型的方法即可. 因此, gRPC 的作用也就很明朗了. 我们在整个编写代码的过程中, 基本上没有涉及到网络底层的编程, 我们的绝大多数操作都是对业务的编程.

有一个特殊的方法, `mustEmbedUnimplementedRouteGuideServer()` 不需要我们单独实现, 这个是由 gRPC 自动生成的 ` pb.UnimplementedRouteGuideServer` 结构体实现的, 我们的自定义的结构体内嵌 `pb.UnimplementedRouteGuideServer` 结构体即可

#### gRPC 的作用

将网络做一层封装. 通过 Protocol Buffer 作为信息编码的载体(类似 XML). 让开发者专注于构建业务逻辑代码

#### gRPC 基本流程

gRPC 的生命周期见: [gRPC core concept/RPC life cycle](https://grpc.io/docs/what-is-grpc/core-concepts/)

#### Go 项目如何使用 gRPC

这一块可以看两份文档, 分别是:
- [A basic tutorial introduction to gRPC in Go](https://grpc.io/docs/languages/go/basics/) 通过一个 demo app 来演示用法
- [Go Generated-code Reference](https://grpc.io/docs/languages/go/generated-code/) 讲解生成的代码的作用和用法, 更具体

安装命令:
```bash
go get google.golang.org/grpc
```

#### natrp 项目实现流程
- 客户端
    - 在一个匿名函数内
        - 客户端通过 gRPC 建立连接. 发送客户端的 metadata (认证相关, 包含认证是否成功的判断)
        - 客户端调用 Msg 方法得到一个 stream
        - 通过 `net.Dial` 建立向本地监听端口的本地连接(比如本地的 nginx 服务会监听在 80 端口, 这一步就是建立一个本地 80 端口的连接)
        - 在一个循环内
        - 在一个 `goroutine` 内
            - 接收来自服务端的 `stream` 信息
            - 通过 `*Conn.Write` 方法将上一步接收到的数据转发给本地连接
            - 在发生错误时, 记录错误并退出匿名函数, 关闭本地连接
            - `goroutine` 内的操作结束
    - 在一个循环内
        - 调用 `*Conn.Read` 方法加载本地连接的响应数据(在 `goroutine` 内我们转发了来自服务器的请求, nginx 接收到请求会响应, 这里就是处理 nginx 响应的)
        - 调用 `stream.Send` 方法发送数据
        - 在发生错误时, 记录错误并退出匿名函数, 关闭 gRPC 连接
    - 重试相关的逻辑代码
- 服务端
    - 在 `server/server.go` 实现了 `ServerServiceServer` 接口
    - 我们主要用到了 Msg 这个 `service` 结构体的方法, 在这个方法内:
        - 获取到 `metadata` 数据. 这个数据用于确定客户端的连接要绑定到网络地址 address:port
        - 根据上一步的信息构建监听并重用这个监听的端口(发送给服务器 port 端口的请求, 响应也将从 port 端口发回给客户端)
        - 等待连接的建立
        - 在一个 goroutine 内:
            - 在一个循环内:
                - 接受来自客户端的数据
                - 将来自 gRPC 客户端的数据转发给通过 TCP 跟服务器建立的连接上
                - 如果出现错误. 记录错误, 关闭通过 TCP 跟服务器建立的连接, 退出当前 goroutine
        - 在一个循环内:
            - 读取通过 TCP 跟服务器建立的连接上的数据
            - 发送给 gRPC 客户端
            - 如果出现错误, 记录错误, 关闭通过 TCP 跟服务器建立的连接, 退出程序(因为已经是 main goroutine)
#### 参考
- [Protocol Buffers 4 官方文档](https://developers.google.com/protocol-buffers/docs/overview)
- [Protocol Buffers 3 中文翻译文档](https://colobu.com/2017/03/16/Protobuf3-language-guide/)
- [Protobuf 终极教程-鸟窝](https://colobu.com/2019/10/03/protobuf-ultimate-tutorial-in-go/)
- [gogo-protobuf](https://github.com/gogo/protobuf)
- [gRPC Golang Quick Start](https://grpc.io/docs/languages/go/quickstart/)
