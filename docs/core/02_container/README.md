# Go 容器结构

容器相关的数据结构(等同于 Java 中的集合类)在项目开发中的经常使用

Go 内置的数据结构

- array
- slice
- map
- set 
  
Go 中没有内置的 set 结构，可以使用 map[string]bool 来表示

```
mySet := map[int]bool{}
mySet[0] = true
mySet[1] = false
if _, ok := mySet[0]; ok {
   // ...
}
```

- container 包

其他一些实现

- [GoDS (Go Data Structures)](https://github.com/psampaz/gods) Implementation of various data structures and algorithms in Go.
- [二叉树及遍历 排序算法 栈 字典树](https://github.com/gaopeng527/go_Algorithm)
- [数据结构和算法 golang 实现](https://goa.lenggirl.com/)