package basic

import (
	"fmt"
	"testing"
)

func TestArrayToSlice(t *testing.T) {
	var a = [3]int{1, 2, 3}
	var b = []int{2, 3}
	ArrayToSlice(a, b)
	

	s := make([]int, 100)
	s[20] = 100
	s1 := s[10:10]
	s2 := s1[10:20]
	fmt.Println(s1)
	fmt.Println(s2)
}

// https://github.com/thinkeridea/example/blob/master/slice/slice_int_to_string_test.go
func BenchmarkSliceIntToString1(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ss := SliceIntToString1(s)
		_ = ss
	}
}

func BenchmarkSliceIntToString2(b *testing.B) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ss := SliceIntToString2(s)
		_ = ss
	}
}
