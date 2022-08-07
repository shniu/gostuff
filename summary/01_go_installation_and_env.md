# Golang

- Go 官方网站：https://go.dev/ 

## Installation

[Go Doc: install](https://go.dev/doc/install)

```shell
# Download url, version is optional.
#  All versions here: https://go.dev/dl/
wget https://golang.org/dl/go1.16.2.linux-amd64.tar.gz

mkdir -p $HOME/go
tar -zxvf go1.16.2.linux-amd64.tar.gz -C $HOME/go
mv $HOME/go/go $HOME/go/go1.16.2

# Config path and go envs
export GOVERSION=go1.16.2                    # Go 版本设置
export GO_INSTALL_DIR=$HOME/go               # Go 安装目录
export GOROOT=$GO_INSTALL_DIR/$GOVERSION     # GOROOT 设置
export GOPATH=$WORKSPACE/golang              # GOPATH 设置
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH    # 将 Go 语言自带的和通过 go install 安装
export GO111MODULE="on"                      # 开启 Go moudles 特性
export GOPROXY=https://goproxy.cn,direct     # 安装 Go 模块时，代理服务器设置
export GOPRIVATE=
export GOSUMDB=off                           # 关闭校验 Go 依赖包的哈希值

# test
go version
```