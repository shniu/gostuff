# Go module

## 演进过程

- go get 
- vendor
- dep
- vgo
- go module

## Go module

Go module 是 module-aware 模式，不同于 gopath 模式

Go module 特性在 Go 1.11 开启实验特性，GO111MODULE 就是这个特性的实验开关，取值有三个：on, off, auto.

- 特点

1. module-aware 模式是 Go 目前包依赖管理的生产标准
2. 一个 module 就是一组相关包组成的一个独立的版本单元
3. 一个仓库只定一个一个 module，最好不要放过多的 module
4. module-aware 模式下包的缓存和存放路径是 $GOPATH/pkg/mod 下
5. 从 Go1.16 开始，GO111MODULE 默认值是 on

- go module 的依赖版版本选择

1. `go list -m -json all` 可以查看 build list 的信息，也就是构建当前 module 所需的所有相关包信息的列表
2. main module 是指 go.mod 所在的 root module
3. go module 在选择版本时使用 MVS (Minimal Version Selection)

- 常用命令

```shell
# Go module
go mod init
go mod tidy
go mod verify

# Clear go mod cache
go clean -modcache
```

- 在 module-aware 模式下，要升级主版本好，怎么做：

```shell

```

- 在 module-aware 模式下，要升级某个依赖的版本，怎么做：

```shell
# 1. find all versions
# go list -m -versions <module path>, e.g.
$ go list -m -versions github.com/spf13/cobra
github.com/spf13/cobra v0.0.1 v0.0.2 v0.0.3 v0.0.4 v0.0.5 v0.0.6 v0.0.7 v1.0.0 v1.1.0 v1.1.1 v1.1.2 v1.1.3 v1.2.0 v1.2.1 v1.3.0 v1.4.0 v1.5.0

# 2. upgrade or degrade
$ go get github.com/spf13/cobra@v1.5.0
$ go mod tidy

# 3. Upgrade to the newest version
$ go get -u github.com/spf13/cobra
$ go mod tidy
```

- go module proxy

GOPROXY=https://goproxy.cn
GOSUMDB=

- 私有 module

在获取企业内部的代码服务器时，可能需要配置 credentials 信息

1. 使用 .netrc 配置 access token

```shell
# ~/.netrc
machine github.com
login shniu
password [personal access tokens]
```

2. 还需要配置绕过 go.sum 校验，私有 module 如何绕过 GOPROXY 呢？可以配置 GOPRIVATE，这样就不需要 GOPROXY 下载，也不需要通过 GOSUMDB 验证校验和

```shell
GOPRIVATE=github.com/shniu/private
```

但是要注意 GONOPROXY 和 GONOSUMDB

3. 此外还有 ssh 的方式，将自己本地的主机公钥添加到 github.com 的 SSH Keys 中，然后配置 .gitconfig

```shell
# ~/.gitconfig
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```

- 升级 module 的主版本号

Go import 包的总原则：如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的；也就是说，如果新旧两个包不兼容，那么应该采用不同的导入路径

```go
// Example
import (
	"github.com/foo/bar"
	barV2 "github.com/foo/v2/bar"
)
```

- 升级自己 module 的主版本号，采用 major branch 方案

主要的方式是通过建立 vN 分支并基于 vN 分支打 vN.x.y 标签的方式进行主版本号的升级

```shell
# 1. 编写自己的 module，完成后
# 2. 可以对当前的代码打标签 v1.0.0
$ git tag v1.0.0
$ git push --tag origin master

# 3. 做不兼容修改，module 升级为 v2.0.0
# 比如删除了某个函数，和老版本是不兼容的
# 接下来要修改 go.mod 文件
module github.com/foo/v2

$ git tag v2.0.0
$ git push --tag origin master
```

一般步骤：

1. 在 go.mod 文件中升级 module 的根路径，增加 vN
2. 建立 vN.x.y 的标签，（也可以不，但这样会使用伪版本号）
3. 特别注意，在修改了主版本号之后，如果内部有相互的包引用，也要跟着一起修改，否则会出现新包引用旧版本的问题

## 自定义 Go 包的导入路径

可以使用 govanityurls 
