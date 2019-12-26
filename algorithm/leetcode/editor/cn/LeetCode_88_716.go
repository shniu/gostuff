package cn
//给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。 
//
// 说明: 
//
// 
// 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。 
// 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。 
// 
//
// 示例: 
//
// 输入:
//nums1 = [1,2,3,0,0,0], m = 3
//nums2 = [2,5,6],       n = 3
//
//输出: [1,2,2,3,5,6] 
// Related Topics 数组 双指针

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := m - 1, n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[i+j+1] = nums1[i]
			i--
		} else {
			nums1[i+j+1] = nums2[j]
			j--
		}
	}
	for j >= 0 {
		nums1[i+j+1] = nums2[j]
		j--
	}
}

//algorithm submit region begin(Prohibit modification and deletion)
func merge(nums1 []int, m int, nums2 []int, n int)  {
	tmpRes := make([]int, m + n)

	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			tmpRes[i+j] = nums1[i]
			i++
		} else {
			tmpRes[i+j] = nums2[j]
			j++
		}
	}

	for i < m {
		tmpRes[i+n-1] = nums1[i]
		i++
	}
	for j < n {
		tmpRes[j+m-1] = nums2[j]
		j++
	}

	copy(nums1, tmpRes)
}
//algorithm submit region end(Prohibit modification and deletion)
