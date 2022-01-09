package basic

import (
	"fmt"
	"strconv"
)

func IterString_fori(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v: %v -> %s\n", i, s[i], string(s[i]))
	}
}

func IterString_range(s string) {
	for i, v := range s {
		fmt.Printf("%v: %v -> %s\n", i, v, string(v))
	}
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("str to int error, str is: %s\n", s)
		return -1
	}
	return i
}

