package leetcode

func twoSum(nums []int, target int) []int {

	// return loop(nums, target)
	return hashgo(nums, target)
}
func loop(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			return []int{i, j}
		}
	}

	return []int{}
}

func hashgo(nums []int, target int) []int {
	var m = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		delta := target - nums[i]
		_, ok := m[delta]
		if ok {
			return []int{m[delta], i}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}