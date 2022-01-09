package stack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack(10)
	stack.Push("abc")
	stack.Push("123")
	stack.Push("789")

	var s = stack.Pop()
	fmt.Print(s)
	assert.Equal(t, "789", s)

	var s2 = stack.Peek()
	assert.Equal(t, "123", s2)

	var s3 = stack.Pop()
	assert.Equal(t, "123", s3)

	stack.Push("hello")
	l := stack.Len()
	assert.Equal(t, 2, l)
}

func TestLinkedStack(t *testing.T) {
	stack := NewLinkedStack()
	stack.Push("1111")
	stack.Push("222")
	stack.Push(12)

	assert.Equal(t, 3, stack.Len())

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	assert.Equal(t, nil, stack.Pop())

	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}
