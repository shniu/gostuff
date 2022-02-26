# Golang

## Golang 介绍与前景

### Golang 为并发而生

Go 语言（或 Golang）是 Google 开发的开源编程语言，诞生于 2006 年 1 月 2 日下午 15 点 4 分 5 秒，于 2009 年 11 月开源，2012 年发布 Go 稳定版。
Go 语言在多核并发上拥有原生的设计优势，Go 语言从底层原生支持并发，无须第三方库、开发者的编程技巧和开发经验。

Go 是非常年轻的一门语言，它的主要目标是“兼具 Python 等动态语言的开发速度和 C/C++ 等编译型语言的性能与安全性”。

很多公司，特别是中国的互联网公司，即将或者已经完成了使用 Go 语言改造旧系统的过程。经过 Go 语言重构的系统能使用更少的硬件资源获得更高的并发
和 I/O 吞吐表现。充分挖掘硬件设备的潜力也满足当前精细化运营的市场大环境。

Go 语言的并发是基于 `goroutine` 的，`goroutine` 类似于线程，但并非线程。可以将 `goroutine` 理解为一种虚拟线程。Go 语言运行时会参与调度 
`goroutine`，并将 `goroutine` 合理地分配到每个 CPU 中，最大限度地使用 CPU 性能。开启一个 `goroutine` 的消耗非常小（大约2KB的内存），
你可以轻松创建数百万个 `goroutine`。

1.`goroutine` 具有可增长的分段堆栈, 这意味着它们只在需要时才会使用更多内存。
2.`goroutine` 的启动时间比线程快。
3.`goroutine` 原生支持利用 channel 安全地进行通信。
4.`goroutine` 共享数据结构时无需使用互斥锁。

### Golang 适合做什么

1. 服务端开发
2. 分布式系统，微服务
3. 网络编程
4. 区块链开发
5. 内存KV数据库，例如boltDB、levelDB
6. 云平台

### 前景

目前 Go 语言已经⼴泛应用于人工智能、云计算开发、容器虚拟化、⼤数据开发、数据分析及科学计算、运维开发、爬虫开发、游戏开发等领域。

Go 语言简单易学，天生支持并发，完美契合当下高并发的互联网生态。Go 语言的岗位需求持续高涨，目前的 Go 程序员数量少，待遇好。

抓住趋势，要学会做一个领跑者而不是跟随者。

国内 Go 语言的需求潜力巨大，目前无论是国内大厂还是新兴互联网公司基本上都会有 Go 语言的岗位需求。

## Golang 基础

### Go 语言思想

Less can be more 大道至简,小而蕴真；让事情变得复杂很容易，让事情变得简单才难 -- 深刻的工程文化。

### Go 优点

1. 自带gc
2. 静态编译，编译好后，扔服务器直接运行
3. 简单的思想，没有继承，多态，类等
4. 丰富的库和详细的开发文档
5. 语法层支持并发，和拥有同步并发的 channel 类型，使并发开发变得非常方便
6. 简洁的语法，提高开发效率，同时提高代码的阅读性和可维护性
7. 超级简单的交叉编译，仅需更改环境变量
8. Go 语言是谷歌 2009 年首次推出并在 2012 年正式发布的一种全新的编程语言，可以在不损失应用程序性能的情况下降低代码的复杂性。
谷歌首席软件工程师罗布派克(Rob Pike)说：我们之所以开发 Go，是因为过去10多年间软件开发的难度令人沮丧。Google 对 Go 寄予厚望，
其设计是让软件充分发挥多核心处理器同步多工的优点，并可解决面向对象程序设计的麻烦。它具有现代的程序语言特色，如垃圾回收，
帮助开发者处理琐碎但重要的内存管理问题。Go 的速度也非常快，几乎和 C 或 C++ 程序一样快，且能够快速开发应用程序。

主要特性：

1.自动垃圾回收。
2.更丰富的内置类型。
3.函数多返回值。
4.错误处理。
5.匿名函数和闭包。
6.类型和接口。
7.并发编程。
8.反射。
9.语言交互性

### Go 语法

#### for loop

```go
package demo
import "fmt"

func goForLoop() {
    // 常规 for loop
    var sum = 0
    for i := 0; i < 10; i++ {
         sum += i
    }

    // while loop
    n := 10
    for n < 10 {
        n *= 2
        n--
    }   
    
    // infinite loop
    sum2 := 0
    for {
        sum2++
    }   
    
    // for each range loop
    strings := []string{"a", "b"}
    for i, s := range strings {
        fmt.Println(i, s)
    }

    // exit a loop
    sum3 := 0
    for i := 1; i < 10; i++ {
        if i % 2 != 0 {
            continue
        }
    
        sum3 += i
    }
    
    // for loop for strings
    for i, j := range "AbCdefg" {
        fmt.Printf("Index %d is %U", i, j)
    }
    
    // for maps
    mmap := map[int]string{
        22: "Geek",
        33: "hello",
    }
    for key, value := range mmap {
        fmt.Println(key, value)
    }

    // for channel
    chn1 := make(chan int)
    go func() {
        chn1 <- 100
        chn1 <- 1000
        chn1 <- 10000
        chn1 <- 100000
        close(chn1)
    }()
    for i := range chn1 {
        fmt.Println(i)
    }
}   
```

- [for ... range 优化](https://www.flysnow.org/2018/10/20/golang-for-range-slice-map.html)


### Go 内置类型和内置函数

TODO

### Go modules

- [Introduction to Go Modules](https://roberto.selbach.ca/intro-to-go-modules/)

### Array & Slices

Golang 中数组类型指定长度和元素类型；数组不需要显示初始化，会按照指定类型的默认值初始化；数组是值类型，数组变量代表了整个数组，
而不是第一个元素的指针，在传递数组时，传递的是数组的拷贝，而不是指针引用.

切片提供了更好的灵活性，所以在 Golang 中经常被使用。切片的零值是 nil，len 和 cap 都是0；数组可以转变成切片

一些基本函数可用于切片：
1. len
2. cap
3. copy
4. append

- [Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices)

- [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) [中文翻译版](https://blog.go-zh.org/go-slices-usage-and-internals)

这是一篇关于 Go Slices 的官方博客，用来介绍 Slice 的底层实现及使用

- [Slice 的使用技巧](https://segmentfault.com/a/1190000018015717)

```golng
// 数组声明
var var_name [SIZE] var_type
e.g.
var mums [10]int
var names [3]string
mums[0] = 1

b := [2]string{"Penn", "Teller”}
b := [...]string{"Penn", "Teller"}

// 声明切片
var a []int
letters := []string{"a", "b", "c", "d”}

// 还可以使用 make
make([]int, 5, 5)

// 数组转切片
s := letters[:]

// 拷贝
func copy(dst, src []T) int

// 在切片后追加
func append(s []T, x ...T) []T

// 引用一个切片或数组
s=ss[:] 
// 清空切片
s=s[:0] 
// 截取接片
s=s[:10] s=s[10:] s=s[10:20] 
// 从切片或数组引用指定长度和容量的切片
s=ss[0:10:20] 
```

#### Maps

- [go maps](https://blog.golang.org/go-maps-in-action)

### Errors

Golang 的错误类型

- [go error are values](https://blog.golang.org/errors-are-values)

### Go 指针

Go 指针的一个用途是将变量的指针作为实参传递给函数，从而在函数内部能够直接修改实参所指向的变量值

- [Go指针](https://www.kancloud.cn/itfanr/go-quick-learn/81640)