# Go 百万连接

如何使用 Go 实现百万的 TCP 连接，是一个值得研究的问题。比如解决服务端向客户端做实时数据推送的场景，我们一般使用 WebSocket，或者直接维护
和客户端的 TCP 连接，无论哪一种落地方式，核心的原理和技术都是一样的。

## Reference

- [Golang 实现轻量、快速的基于 Reactor 模式的非阻塞 TCP 网络库](https://juejin.cn/post/6844903945907748872) and [here](https://note.mogutou.xyz/articles/2019/09/19/1568896693634.html)
- [百万 Go TCP 连接的思考: epoll方式减少资源占用](https://colobu.com/2019/02/23/1m-go-tcp-connection/)
- [百万 Go TCP 连接的思考2: 百万连接的吞吐率和延迟](https://colobu.com/2019/02/27/1m-go-tcp-connection-2/)
- [百万 Go TCP 连接的思考: 正常连接下的吞吐率和延迟](https://colobu.com/2019/02/28/1m-go-tcp-connection-3/)
- [Going Infinite, handling 1M websockets connections in Go](https://speakerdeck.com/eranyanay/going-infinite-handling-1m-websockets-connections-in-go)
- [A Million Websocket and Go](https://www.freecodecamp.org/news/million-websockets-and-go-cc58418460bb/)
- [使用4种框架分别实现百万 Websocket 连接](https://colobu.com/2015/05/22/implement-C1000K-servers-by-spray-netty-undertow-and-node-js/)
- [七种WebSocket框架的性能比较](https://colobu.com/2015/07/14/performance-comparison-of-7-websocket-frameworks/)
- [100万并发连接服务器笔记系列](http://www.blogjava.net/yongboy/category/54842.html)
- [千万级并发实现的秘密：内核不是解决方案，而是问题所在！](https://www.csdn.net/article/2013-05-16/2815317-The-Secret-to-10M-Concurrent-Connections)
- [字节跳动在 Go 网络编程上的实践](https://www.infoq.cn/article/fea7chf9moohbxbtyres)
- [第 65 期 Go 原生网络模型 vs 异步 Reactor 模型](https://bytemode.github.io/reading/65-2019-10-31-go-net/)

### Code in GitHub

- An example code: https://github.com/eranyanay/1m-go-websockets/
- An example code: https://github.com/smallnest/1m-go-tcp-server
