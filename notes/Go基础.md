
## Golang 基础知识

### Array & Slices

Golang 中数组类型指定长度和元素类型；数组不需要显示初始化，会按照指定类型的默认值初始化；数组是值类型，数组变量代表了整个数组，
而不是第一个元素的指针，在传递数组时，传递的是数组的拷贝，而不是指针引用.

切片提供了更好的灵活性，所以在 Golang 中经常被使用。切片的零值是 nil，len 和 cap 都是0；数组可以转变成切片

一些基本函数可用于切片：
1. len
2. cap
3. copy
4. append

- [Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals)

这是一篇关于 Go Slices 的官方博客，用来介绍 Slice 的底层实现及使用

- [Slice 的使用技巧](https://segmentfault.com/a/1190000018015717)

```go
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

### Errors

Golang 的错误类型

- [go error are values](https://blog.golang.org/errors-are-values)