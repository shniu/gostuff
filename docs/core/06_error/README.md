
## Golang 的异常处理

Golang 使用 error 来处理异常，需要显示的来处理异常，这个有别于 Java 的 Exception 机制，与 c/c++ 也不同。
Golang 支持多参数返回，将 err 放入返回值中，调用完成后需要对 err 进行立即处理。

结论：
1. error 和 exception 是有区别的：使用多个返回值和一个简单的约定， Go 解决了让程序员知道什么时候出了问题，并为真正的异常情况保留了 panic
2. Golang 使用带有 error 的多返回值来实现异常处理，这有别于 Java 的 Exception 机制，对于返回的 error 要立即处理，不做任何假设；
除了 error，Golang 中还有 panic，为我们提供一些程序无法处理的异常的解决方案，如果抛出 panic 说明程序就是挂了，无法再运行了。 
3. Golang 的异常处理简洁明确。

代码在：[best/error.go](../../../best/error.go)

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

思考：errors.New(...) 为何返回 *string, 而不是 string
返回 *string 可以避免很多不必要的麻烦，程序会更加健壮少犯错误，返回 string 比较的是值，返回 *string 比较的是指针地址

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

### Error 类型

- sentinel error
- Error types

实现了 error 接口的自定义类型

- 不透明错误 (Opaque errors)

不透明错误处理的全部功能：只需返回错误而不假设其内容
使用行为来判断错误，而不是类型

```
// 比如
// 自定义的
func timeout interface {
    // ...
    Timeout bool
}
func (t *timeout)Error() string {
    return "Read timeout"
}

func IsTimeout(err error) bool {
    t, ok := err.(timeout)
    return ok && t.Timeout()
}
```

- Wrap errors

sentinel error, error types, 不透明的 error 处理在应用代码中都存在缺陷，Wrap errors 是一个好的解决方案。

pkg/errors 以及 Go 1.13 的 errors 新特性(Unwrap() / Is() / As()) 实现了 Wrap errors 的能力。

关于 error 处理的一些原则：
1. 应当只处理错误一次，处理一个异常意味着检查错误值，并作出一个决定：要么打印日志，让流程继续；要么返回 error，退出流程
2. Go 中的错误处理契约规定，在出现错误的情况下，不能对其他返回值的内容做出任何假设
3. 日志记录与错误无关且对调试没有帮助的信息应被视为噪音，应予以质疑。记录的原因是因为某些东西失败了，而且日志包含了答案
  1. 错误要被日志记录
  2. 应用程序处理错误，保证 100% 完整性
  3. 之后不再报告当前错误

通过使用 pkg/errors 包，可以向错误值添加上下文，这种方式既可以由人来处理也可以由机器检查 (一定是在应用层代码才更加推荐使用 wrap errors)

### 如何选择

在处理异常时有几种情况：

- 在自己写应用代码时，代码内部的一些错误处理，推荐使用 errors.New(...) 或者 errors.Errorf(...) 返回错误

```
func parseArgs(args []string) error {
    if len(args) < 3 {
        return errors.Errorf("Not enough arguments, excepted at least three args.")
    }
    // ...
}

```

- 如果调用其他包内的函数，只需要简单直接返回
- 如果和其他库进行协作，比如基础库和第三方库等，考虑使用 errors.Wrap 或者 errors.Wrapf 保存堆栈信息

````
f, err := os.Open(path)
if err != nil {
    return errors.Wrapf(err, "failed to open %q", path)
}
````

- 直接返回错误，而不是每个错误产生的地方到处打日志
- 在程序的入口或者 goroutine 顶部，使用 %+v 打印堆栈信息

```
func main() {
    err := app.Run()
    if err != nil {
        fmt.Printf("FATAL: %+v\n", err)
        os.Exit(1)
    }
    // ...
}
```

- 使用 errors.Cause 获取 root error，再和 sentinel error 判定
- Packages that are reusable across many projects only return root error values.

只有 applications 才选择 wrap error；如果一个可重用性非常高的包，不要使用 wrap error，而应该返回原始错误，比如自定义的 sentinel error
比如 sql.BadSQLError

- If the error is not going to be handled, wrap and return up the call stack.
- Once an error is handled, it is not allowed to be passed up the call stack any longer.

一旦确定函数 / 方法将处理错误，错误就不再是错误。如果函数 / 方法仍然需要发出返回，则它不能返回错误值。它应该只
返回零（比如降级处理中，你返回了降级数据，然后需要return nil）。

目前还是需要 pkg/errors

### errors after Go1.13

```
// Before Go1.13
// 1
var ErrNotFound = errors.New("Not Found")
if err == ErrNotFound {
    // ...
}

// 2
type TimeoutError struct {
    Name string
    Err error
}
func (e *TimeoutError) Error() string {
    return e.Name + ": timeout"
}

// 使用
if e, ok := err.(*TimeoutError); ok {
    // TimeoutError 处理逻辑
}

// 3
// 打印错误栈
if err != nil {
    return fmt.Errorf("decompress %v: %v", name, err)
}

// 或者
type QueryError struct {
    Err error
    Query string
}
func (e *QueryError) Error() string {
    return ...
}
// 使用
if e, ok := err.(*QueryError); ok && e.Err = ErrNotFound {
    // do something
}
```

After Go1.13

Go1.13 引入了新特性：
1. Unwrap, Is, As 等
2. fmt 增加 %w 

```
// if err == ErrNotFound { ... }
// 类似
if errors.Is(err, ErrNotFound) {
    // something wasn't found
}

// if e, ok := err.(*QueryError); ok { ... }
// 类似
var e *QueryError
if errors.As(err, &e) {
    // err is a *QueryError
}
```

## Reference

- [Go errors with additional details](https://romanyx90.medium.com/go-errors-with-additional-details-66873577f3a9)

### Go2 Error Propasal

- https://go.googlesource.com/proposal/+/master/design/29934-error-values.md
- https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md
- https://github.com/golang/go/wiki/Go2ErrorValuesFeedback
- https://go.googlesource.com/proposal/+/master/design/go2draft-error-handling-overview.md

### Blogs related error handling

- [Go 语言的错误处理机制引发争议](https://www.infoq.cn/news/2012/11/go-error-handle/)

重点讨论了 Go 的错误处理机制和其他语言的不同，Go 选择了多返回值 + error 的方式，painc + recover 是针对特殊情况设计的
Go 的设计遵循了对错误清洗简单的处理，我们明确知道什么情况下应该如何处理异常，如果不处理就会有潜在的风险
此外，Go 是针对大型软件而设计的，try + catch 方式的异常/错误捕获方式可能是一种心智负担或者拖累，Go 语言返回错误的方式，对于调用者
来说可能不是很方便，但这样会让程序中可能出错的地方显的更加明显。

- [Why does Go not have exceptions?](https://go.dev/doc/faq#exceptions)

Go 的开发者认为，将异常和控制结构耦合在一起，会导致代码复杂化，将过多的错误标记为异常也不是一种好的实践。

