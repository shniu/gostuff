# Go 并发编程

Golang 支持 `go` 关键字，可以快速让一个函数创建为 `goroutine`. `main` 函数就是一个特殊的 `goroutine`, 即使使用这个单一的
逻辑处理器和操作系统线程，也可以调度数十万 `goroutine` 以惊人的效率和性能并发运行。

## Goroutine

- 空的 `select` 语句将永远阻塞

```
// 将永远阻塞
select {}
```

- 启动一个 http server

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ...
	})
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

- Leave concurrency to the caller 把并发的选择留给调用者

```
// 同步返回目录列表，缺点是在大目录的情况下会返回很多数据，并且占用大量内存
func ListDir(dir string) ([]string, error)

// 以 chan string 的方式不断返回目录，当通道关闭时，这表示不再有目录。由于在ListDirectory返回后发生通道的填充，
// ListDirectory 可能内部启动 goroutine 来填充通道。
func ListDir(dir string) chan string
```

ListDir chan 这个版本存在几个问题：

1. 通过一个关闭的 chan 作为不再需要处理目录的信号，无法告诉调用者返回的目录不完整，因为遇到了错误，调用发无法区分空目录于完全从目录读取的错误之间的区别
2. 调用者必须从通道读取目录信息，直到关闭，这是调用者停止的唯一方法；调用者及时已经获取到想要的结果，也必须等到所有目录都返回才能结束，
效率上并不比返回 slice 快多少
   
一个改进版本

```
// 使用一个函数来接收目录中的文件
func ListDir(dir string, fn func (string))

// 调用方式
ListDir("/tmp", func (name string) {
   // do something with name
})
```

- Never start a goroutine without knowing when it will stop

在启动一个 goroutine 时，我们需要知道：它合适停止运行，以及可以阻止运行吗，否则就会出现 goroutine 泄漏问题

如果要启动两个端口的服务端程序，该怎么做？

```go
package main

import (
	"context"
	"fmt"
	"net/http"
)

func serve(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()

	// 这里会被阻塞，不会返回任何东西，当出现错误时返回 error
	return s.ListenAndServe()
}

func main() {
	done := make(chan error, 2)
	stop := make(chan struct{})
	go func() {
		// 当有返回值时，会放入 done channel 中，说明可以终止服务了
		done <- serve(":8080", nil, stop)
	}()
	go func() {
		// 当有返回值时，会放入 done channel 中，说明可以终止服务了
		done <- serve(":8081", nil, stop)
	}()

	var stopped bool
	for i := 0; i < cap(done); i++ {
		// 这里会一直阻塞，直到 done 中有数据可读
		if err := <-done; err != nil {
			fmt.Printf("error: %v\n", err)
		}
		
		// 触发一次 stop
		if !stopped {
			stopped = true
			// 关闭 channel 的信号
			close(stop)
        }
	}
}
```