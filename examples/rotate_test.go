package examples

import (
	"testing"
	"fmt"
)

func TestRotate1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 2)
	fmt.Print(nums)

	nums2 := []int{1, 2, 3, 4, 5, 6}
	rotate2(nums2, 2)
	fmt.Print(nums2)
}
