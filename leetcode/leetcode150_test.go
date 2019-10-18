package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEvalRPN(t *testing.T) {
	v := evalRPN([]string{"2", "1", "+", "3", "*"})
	fmt.Println(v)
}

type queue []int

func (q *queue) offer(val int) {
	fmt.Println(reflect.TypeOf(q))
	fmt.Println(reflect.TypeOf(*q))
	fmt.Println(reflect.TypeOf(&*q))
	fmt.Println(reflect.TypeOf(&q))
	fmt.Println(reflect.TypeOf(**&q))
}

func (q queue) poll() {
	q = append(q, 1)
}

func TestTypeQueue(t *testing.T) {
	fmt.Println(reflect.TypeOf(queue{}))

	q := &queue{}
	q.offer(123)
}
