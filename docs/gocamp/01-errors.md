
## Golang 的异常处理

Golang 使用 error 来处理异常，需要显示的来处理异常，这个有别于 Java 的 Exception 机制，与 c/c++ 也不同。
Golang 支持多参数返回，将 err 放入返回值中，调用完成后需要对 err 进行立即处理。

Golang 的异常处理包括：

1. error 机制
2. panic

### error 机制

- 可以使用 errors.New("buf: EOF") 来创建一个 error
- fmt.Errorf(...) 来创建一个 error
- 自定义 error struct 来实现 error interface

```
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

### panic

Go panic 意味着 fatal error，程序会直接挂掉，这个有别于 exception，panic 是无法被解决的。

使用多个返回值和一个简单的约定，Go 解决了让程序员知道什么时候出了问题，并为真正的异常情况保留了 panic。

对于真正意味的情况，那些表示不可恢复的程序错误，例如索引越界、不可恢复的环境问题、栈溢出，我们才使用 panic。对于其他的错误情况，我们应该是期望
使用 error 来进行判定。

- you only need to check the error value if you care about the result.
- My point isn't that exceptions are bad. My point is that exceptions are too hard and I'm not smart enough to handle them.

Go 的错误处理思想：

- 简单
- 考虑失败，而不是成功（plan for failure, not success)
- 没有隐藏的控制流
- 完全交给你来控制 error
- Error are values

### Error type

- sentinel error

- Error types

实现了 error 接口的自定义类型

- 不透明错误

### Handling error

- Wrap errors

pkg/errors 以及 Go 1.13 的 errors 新特性(Unwrap() / Is() / As())

## Reference

- [Go errors with additional details](https://romanyx90.medium.com/go-errors-with-additional-details-66873577f3a9)