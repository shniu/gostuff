package examples

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)

	// 记录移动次数
	moveCnt := 0
	startPos := 0

	for moveCnt < length {
		// 记录
		currentPos := startPos
		currentValTmp := nums[startPos]

		// 因为存在一种情况是：在找next位置交换时，会回到开始的位置，形成环
		nextPos := (currentPos + k) % length
		for nextPos != startPos {
			nextPos = (currentPos + k) % length

			// 换到nextPos上
			nextValTmp := nums[nextPos]
			nums[nextPos] = currentValTmp

			// 移动到nextPos位置，准备进入下一轮交换
			currentPos = nextPos
			currentValTmp = nextValTmp

			// 移动次数+1
			moveCnt++
		}

		startPos++
	}
}

func rotate2(nums []int, k int) {
	length := len(nums)
	k %= length
	temp := append(nums[length-k:], nums[:length-k]...)
	for i := 0; i < length; i++ {
		nums[i] = temp[i]
	}
}
