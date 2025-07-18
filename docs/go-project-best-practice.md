# 优雅的 Go 项目

经常遇到的代码问题，比如：

- 代码不规范，难以阅读；
- 函数共享性差，代码重复率高；
- 不是面向接口编程，代码扩展性差，代码不可测；
- 代码质量低下

究其原因，是因为这些代码的开发者很少花时间去认真研究如何开发一个优雅的 Go 项目，更多时间是埋头在需求开发中。

## 如何写出优雅的 Go 项目

一个优雅的 Go 应该具备的特点：

1. 符合 Go 编码规范和最佳实践；
2. 易阅读、易理解，易维护；
3. 易测试、易扩展；
4. 代码质量高；

写出一个优雅的 Go 项目，就是用“最佳实践”的方式去实现 Go 项目中的 Go
应用、项目管理和项目文档。具体来说，就是编写高质量的 Go 应用、高效管理项目、编
写高质量的项目文档。

### 编写高质量的 Go 应用

- 代码结构
- 代码规范
- 代码质量：通过单元测试和 Code Review 来实现
- 编程哲学：Go 语言有很多设计哲学，对代码质量影响比较大的，我认为有两个：面向接口编程和面向“对象”编程
- 软件设计方法：设计模式（Design pattern）和 SOLID 原则

### 高效管理项目

制定一个高效的开发流程、使用 Makefile 管理项目和将项目管理自动化。

- 高效的开发流程
- 使用 Makefile 管理项目
- 自动生成代码
- 善于借助工具
- 对接 CI/CD：当前比较流行的 CI/CD 工具有 Jenkins、GitLab、Argo、Github Actions、JenkinsX 等

### 编写高质量的项目文档

README.md、安装文档、开发文档、使用文档、API 接口文档、设计文档等等
