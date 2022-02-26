# gRPC

- https://grpc.io/
- https://developers.google.com/protocol-buffers/docs/overview

使用 Protocol Buffers

1. 定一个 proto 文件：用于定义你想要表达、表示、传输交换的数据结构
2. 使用 message 定义一个消息
3. message 需要有个名字，比如 message HelloRequest
4. 结构体中每行定义一个 name-values 组成的 field

```protobuf
// 定义 message
message PersonReply {
  string name = 1;
  int32  id = 2;
  bool has_children = 3;
}

message QueryPersonRequest {
  string name = 1;
}

// 定义 service
service PersonService {
  rpc QueryPersons (QueryPersonRequest) returns (PersonReply);
}

```

生成代码

```shell
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```

- gRPC
  - 服务定义
    - 第一种：(Unary) 单个输入，单个返回 `rpc SayHello (HelloRequest) returns (HelloResponse)`
    - 第二种：(Server streaming RPC) client 发送 request，获取到一个 stream，客户端读数据直到没有数据可读，gRPC 保证了数据的顺序性 `rpc LotsOfReplies (HelloRequest) returns (stream HelloResponse)`
    - 第三种：(Client streaming RPC) client 发送一个 stream 获取到一个返回 `rpc LotsOfGreetings (stream HelloRequest) returns (HelloResponse)`
    - 第四种：(Bidirectional streaming RPC) 双向流 `rpc BidHello (stream HelloRequest) returns (stream HelloResponse)`
  - 根据 .proto 文件自动生成服务端和客户端代码
    - 服务端实现接口定义的功能，并使用 gRPC Server 供客户端调用，gRPC 的基础设施会解码请求，执行服务方法，编码响应返回客户端
    - 在客户端，会有一个 stub 代理所有发给服务端的请求；客户端调用 stub 实现的 method，stub 会调用远程的 Server 实现的 method
    - Client 和 Server 通信的协议和数据表示是使用 Protocol Buffers
  - 同步 RPC 和异步 RPC
    - 同步 RPC 阻塞调用直到服务器响应到来；但是网络本质上是异步的，在启动 RPC 后而不阻塞当前线程是更加好的处理方式
    - 每种编程语言都提供了同步和异步两种 API
    



