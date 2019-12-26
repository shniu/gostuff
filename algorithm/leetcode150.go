package algorithm

import "strconv"

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/

type stack []int

func (s *stack) push(val int) {
	*s = append(*s, val)
}

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) pop() int {
	res := (*s)[s.len()-1]
	*s = (*s)[:s.len()-1]
	return res
}

func newStack() *stack {
	return &stack{}
}

func evalRPN(tokens []string) int {
	os := newStack()

	for i := 0; i < len(tokens); i++ {

		if tokens[i] == "+" {
			// 最后两个元素出栈
			top1 := os.pop()
			top2 := os.pop()

			// 加法
			os.push(top1 + top2)
		} else if tokens[i] == "-" {
			// 最后两个元素出栈
			top1 := os.pop()
			top2 := os.pop()

			// 减法
			os.push(top2 - top1)
		} else if tokens[i] == "*" {
			// 最后两个元素出栈
			top1 := os.pop()
			top2 := os.pop()

			// 乘法
			os.push(top1 * top2)
		} else if tokens[i] == "/" {
			// 最后两个元素出栈
			top1 := os.pop()
			top2 := os.pop()

			// 除法
			os.push(top2 / top1)
		} else {
			// 压栈
			val, _ := strconv.Atoi(tokens[i])
			os.push(val)
		}
	}

	return os.pop()
}
