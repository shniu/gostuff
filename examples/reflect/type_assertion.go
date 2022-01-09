package reflect

import (
	"fmt"
	"reflect"
)

// ref: https://golang.org/ref/spec#Type_assertions

func typeAssert() {
	a, b := 0, 0
	// fmt
	fmt.Printf("type: %T, %T\n", a, b)

	// 反射
	ta := reflect.TypeOf(a).String()
	tb := reflect.TypeOf(b).String()
	fmt.Printf("type: %s, %s\n", ta, tb)
}


// 类型推断
func typeOf(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	default:
		_ = t
		return "Unknown"
	}
}
