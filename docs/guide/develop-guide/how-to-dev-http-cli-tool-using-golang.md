# 使用 Golang 写一个实用的 CLI 工具

## 需求

实现一个 Go 版本的 httpie ([httpie 的介绍看这里](https://httpie.io/))。

## 分析

httpie 是一个 CLI 工具，简化访问 http API 服务。

要实现一个 Go 的 HTTPie CLI 小工具需要如下功能：

- 解析命令行参数，校验用户的输入，处理子命令和参数等，并且有 help 查看命令使用
- 根据参数发起 http 请求，获得响应
- 以友好的方式输出，比如格式化、多彩显示等

根据以上的功能，在使用 Go 开发时，需要借助一些小工具

- CLI 解析工具
  - Go 内置了 Flag 命令行解析工具，但是功能有限
  - Cobra 是一个实用广泛的第三方库，可以考虑使用
- HTTP 客户端工具，Go 内置了 http client
  - https://www.loginradius.com/blog/engineering/tune-the-go-http-client-for-high-performance/
  - https://mailazy.com/blog/http-request-golang-with-best-practices/
  - https://stuartleeks.com/posts/connection-re-use-in-golang-with-http-client/
- 命令行多彩输出
  - https://www.infoq.cn/article/jjqljfltft8b4ogijoof
  - https://github.com/gookit/color/blob/master/README.zh-CN.md
  - https://github.com/fatih/color
- json 格式处理
  - encoding/json
- 异常处理设计
- mime 处理
- go http client 异步处理
  - https://github.com/freshcn/async
  - https://mj-go.in/golang/async-http-requests-in-go
  - https://zetcode.com/web/async-http-requests/
  - https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1
  - https://github.com/tomcatzh/asynchttpclient/blob/master/asynchttpclient.go


## 设计

- TODO
  - 使用 Cobra 构建命令行工具，需要同时弄清楚 Flags 和 Cobra 的用法

## 实现


