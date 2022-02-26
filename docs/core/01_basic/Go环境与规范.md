# Golang 环境与开发规范

## Go 安装和环境搭建

1. [Golang Install](https://golang.org/doc/install)

## Golang 开发规范

可参考 

1. [Uber go guide](https://github.com/uber-go/guide)
2. https://golang.org/doc/effective_go
3. https://github.com/golang/go/wiki/CodeReviewComments
4. https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins

- [go 的 package 指南](https://rakyll.org/style-packages/)

在使用 golang 编写代码时，package 是分离和组织代码的方式。良好的代码组织有利于可读性、可维护性和后期的迁移成本。
通常我们可以遵循如下的一些原则：

1. 尽可能在 package 中使用多个文件来组织逻辑上松耦合、独立的逻辑，尽可能有更好的可读性
2. 将类型保持在更靠近其使用位置的位置。这样，任何维护人员（不仅是原始作者）都可以轻松找到类型。 `Header` 结构类型的好地方可能在 `headers.go` 中，并且尽可能的将类型定义放在文件的开始处
3. 在 Go 中尽可能的使用功能职责来组织代码
4. 在包的 API 设计的早期阶段使用 `godoc` 是一个很好的实践，以了解您的概念将如何在 doc 上呈现。有时，可视化也会对设计产生影响。Godoc 是用户使用软件包的方式，因此可以对事物进行调整以使其更易于访问。运行 `godoc -http = <hostport>` 在本地启动godoc服务器
5. 提供示例以帮助用户发现和理解如何使用
6. 主软件包不可导入，因此不需要从主软件包中导出标识符。如果要将包构建为二进制文件，请不要从主包中导出标识符。

软件包名称和导入路径都是软件包的重要标识符，代表了软件包包含的所有内容。规范地命名软件包不仅可以提高代码质量，而且可以改善用户的代码质量。

1. package 命名时只用小写
2. 简短但具有代表性的名字
3. 避免将你的自定义存储库结构暴露给用户，与 GOPATH 约定保持一致。避免在导入路径中包含 src/，pkg/部分
4. 不使用复数
5. 如果要导入多个具有相同名称的软件包，则可以在本地重命名软件包名称, 重命名应遵循上面提到的相同规则

文档化 package
有时，打包文档会变得很冗长，尤其是当它们提供用法和指南的详细信息时。将软件包 godoc 移至 doc.go 文件
doc.go 可以参考 https://github.com/googleapis/google-cloud-go/blob/master/datastore/doc.go