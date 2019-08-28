package leetcode

import (
	"strconv"
	"strings"
)

// Solution: stack
func decodeString(s string) string {
	// 倍数
	var multi int

	// 结果
	var res []string

	// stack
	var multiStack []int
	var resStack []string

	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 数字
		if v, err := strconv.Atoi(string(s[i])); err == nil {
			multi = multi * 10 + v
			continue
		}

		// [
		if string(s[i]) == "[" {
			multiStack = append(multiStack, multi)
			resStack = append(resStack, strings.Join(res, ""))

			multi = 0
			res = res[:0]
			continue
		}

		// ]
		if string(s[i]) == "]" {
			// 取出倍数
			multi = multiStack[len(multiStack) - 1]

			var tmp []string
			for i := 0; i < multi; i++ {
				tmp = append(tmp, strings.Join(res, ""))
			}

			top := resStack[len(resStack) - 1]
			res = []string{top}
			res = append(res, strings.Join(tmp, ""))

			multiStack = multiStack[:len(multiStack) - 1]
			resStack = resStack[:len(resStack) - 1]
			multi = 0
			continue
		}

		// 字母
		res = append(res, string(s[i]))
	}

	return strings.Join(res, "")
}
