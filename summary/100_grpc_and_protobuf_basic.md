# gRPC and Protobuf

- [http://google.github.io/proto-lens/installing-protoc.html](http://google.github.io/proto-lens/installing-protoc.html)
- [https://developers.google.com/protocol-buffers](https://developers.google.com/protocol-buffers)

## Installation

```shell
### Macos
brew install protobuf
brew upgrade protobuf

# or 下载安装包安装
PROTOC_ZIP=protoc-3.14.0-osx-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

### Linux
PROTOC_ZIP=protoc-3.14.0-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.14.0/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP

# 然后，安装 protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go
```

The protocol buffer compiler requires a plugin to generate Go code, and protoc-gen-go is the tool, 详细内容: https://developers.google.com/protocol-buffers/docs/reference/go-generated 

