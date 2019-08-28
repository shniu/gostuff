package leetcode

//m := make(map[string]string)
//m[")"] = "("
//m["}"] = "{"
//m["]"] = "["

func isValid(s string) bool {
	m := map[string]string{")": "(", "}": "{", "]": "["}

	var stack []string

	for _, v := range s {
		vv, ok := m[string(v)]
		if ok {
			if len(stack) == 0 || stack[len(stack) - 1] != vv {
				return false
			}
			stack = stack[:len(stack) - 1]
		} else {
			stack = append(stack, string(v))
		}
	}

	return len(stack) == 0
}
