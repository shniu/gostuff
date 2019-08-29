package examples

import "testing"

func TestArrayToSlice(t *testing.T) {
	var a = [3]int{1, 2, 3}
	var b = []int{2, 3}
	ArrayToSlice(a, b)
}
