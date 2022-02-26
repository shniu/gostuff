package error

import "fmt"

func helloErr() {
	defer func() {
		fmt.Println("defer func is called.")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	panic("a panic is triggered.")
}

type MyError struct {
	Msg string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func test() error {
	return &MyError{"Something happened", "server.go", 42}
}
