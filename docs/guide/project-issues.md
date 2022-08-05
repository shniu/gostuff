# 工程项目相关问题

## go module

- goproxy
 - 如果在国内，没有 vpn 最好配置 goproxy
 - `go env -w GOPROXY="https://goproxy.cn,direct"`
- go.mod 文件
  - go replace 的作用
  - go require
- go 的包文件
  - 使用 go module 后，下载的 go 包位置在 $GOPATH/pkg/mod 下
  - 可以使用 GOMODCACHE 环境变量设置自定义的位置
- 私有仓库
  - GOPRIVATE 控制私有仓库
