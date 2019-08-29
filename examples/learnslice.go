package examples

import (
	"fmt"
	"reflect"
)

func ArrayToSlice(a [3]int, b []int) {
	// Check type
	fmt.Printf("type: %T, %T\n", a, b)

	ta := reflect.TypeOf(a).String()
	tb := reflect.TypeOf(b).String()
	fmt.Printf("type: %s, %s\n", ta, tb)

	// Array to slice
	s := a[:]
	fmt.Printf("type: %T\n", s)

	// Copy
	var d = make([]int, len(s), cap(s) * 2 + 1)
	copy(d, s)
	fmt.Printf("copy: %v\n", d)

	// Append
	var e = []int{4, 5, 6}
	d = append(d, e...)
	d = append(d, 7)
	fmt.Printf("append: %v\n", d)

	// append a slice of nil
	var n []int
	n = append(n, d...)
	fmt.Printf("append nil: %v\n", n)
	d[0] = 11
	fmt.Printf("after modified: %v\n", d)
	fmt.Printf("after modified: %v\n", n)

	// reslice a slice
	var dd = d
	fmt.Printf("before reslice: %v, %v\n", len(dd), cap(dd))
	fmt.Printf("before reslice: %v, %v\n", len(d), cap(d))
	fmt.Printf("type: %T, %T\n", dd, d)
	d = d[1:]
	fmt.Printf("after reslice: %v, %v\n", len(d), cap(d))
	fmt.Printf("after reslice: %v, %v\n", len(dd), cap(dd))
}

func typeOf(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	default:
		_ = t
		return "Unknown"
	}
}
