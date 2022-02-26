# Go select

Go 语言中的 select 能够让 Goroutine 同时等待多个 Channel 可读或者可写，在多个文件或者 Channel状态改变之前，
select 会一直阻塞当前线程或者 Goroutine。

特殊情况

1. 如果 select 中包括 default，就会执行 default，但不会阻塞，也就是非阻塞的

select 是与 switch 相似的控制结构，与 switch 不同的是，select 中虽然也有多个 case，但是这些 case 中的表达式必须都是 Channel 的收发操作。

```go
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
```

## Reference

- [golang select](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-select/)