
# Go Socket

## 博客

- [IO多路复用与Go网络库的实现](https://ninokop.github.io/2018/02/18/go-net/)
- [Go语言: 万物皆异步](https://blog.csdn.net/neosmith/article/details/78884129)
- [深入Go语言网络库的基础实现](http://skoo.me/go/2014/04/21/go-net-core)
- [Go语言TCP网络编程(详细)](https://www.iminho.me/wiki/docs/gopher-reading-list/gocn-2018-03-read-7.md)

## Go 网络库实现

// todo

## 实现一个 RPC 框架

### 一个简单的作业

写一个基于RPC调用的服务，让客户端调用服务端，调用10万次并记录下耗时

设计要点：

1. 异步设计的方法
2. 异步网络 IO
3. 序列化和反序列化方法
4. 设计良好的传输协议
5. 双工通信


