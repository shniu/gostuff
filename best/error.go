package best

import (
	"database/sql"
	"errors"
	"fmt"

	xerrors "github.com/pkg/errors"
)

// 总结 Golang 中异常处理的最佳实践

// 1. Sentinel Error
// 这种方式在基础包中被大量使用，但是在我们自己的程序代码中不建议使用
// 使用 errors.New 创建一个新的异常，返回对 errorString 的指针引用
var ErrEOF = errors.New("best/error: EOF Exception.")
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)

// 2. Error Types
// 自己定义一个 struct 来实现 error 接口
// Error 接口
// type error interface {
//	Error() string
// }

// 比如 os.PathError
type MyError struct {
	Text string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d %s", e.File, e.Line, e.Text)
}

func test() error {
	return &MyError{"", "", 0}
}

func checkErr() {
	err := test()

	switch err.(type) {
	case nil:
		// nothing to do
	case *MyError:
		fmt.Println("MyError")
	default:
		fmt.Println("unknown")
	}
}

// 3. Wrap errors
// pkg/errors 或者 Go1.13 的新特性

func doQuery(q *string) error {
	return sql.ErrNoRows
}

func WrapSqlErr() {
	querySql := "select * from user"
	err := doQuery(&querySql)

	wrappedErr := fmt.Errorf("Sql exception: %v, %w", querySql, err)

	if errors.Is(wrappedErr, sql.ErrNoRows) {
		fmt.Println("error is sql.ErrNoRows")
		fmt.Printf("%+v\n", wrappedErr)
	}
}

type MyDbError struct {
	Msg string
	Err error
}

func (e *MyDbError) Error() string {
	return e.Msg
}

func (e *MyDbError) Unwrap() error {
	return e.Err
}

func doExecute(s *string) error {
	// do something
	return &MyDbError{
		Msg: "Executing sql error",
		Err: sql.ErrNoRows,
	}
}

func MyDbErrorDemo() {
	s := "select * from users"
	err := doExecute(&s)

	if errors.Is(err, sql.ErrNoRows) {
		// ....
	}
}

// pkg/errors
var errMy = errors.New("my error")
func test0() error {
	return xerrors.Wrapf(errMy, "test0 failed")
}

func test1() error {
	return test0()
}

func test2() error {
	return test1()
}

func PkgErrors() {
	err := test2()
	fmt.Printf("main: %+v\n", err)
}
