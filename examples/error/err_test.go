package error

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestErr(t *testing.T) {
	helloErr()
}

func TestMyError_Error(t *testing.T) {
	err := test()
	fmt.Printf("%+v", err)

	switch err := err.(type) {
	case nil:
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
		// do nothing
	}
}

func TestOsPathError(t *testing.T) {
	pathError := os.PathError{"CreateFile", "/tmp/camp/000", errors.New("file not exists.")}
	fmt.Println(pathError)
}
