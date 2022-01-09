package basic

import (
	"testing"
	"fmt"
)

func TestForRange1(t *testing.T) {
	five := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range five {
		five = five[:2]
		fmt.Printf("v[%s]\n", v)
	}
	// output:
	// v[Annie]
	// v[Betty]
	// v[Charley]
	// v[Doug]
	// v[Edward]

	fmt.Print(five) // output: [Annie Betty]
	// 说明以上是 for...range 的值语义，对 five 的修改并不会影响什么，
	// 因为for range操作的是five的副本，不是five本身
}

func TestForRange2(t *testing.T) {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", five[1])

	for i, v := range five {
		five[1] = "Jack"

		if i == 1 {
			fmt.Printf("v[%s]\n", v)
		}
	} // output: Bfr[Betty] : v[Betty]

	fmt.Print(five)  // output: [Annie Jack Charley Doug Edward]
	// 说明可以更改原来的数组，但是不会对当前的循环有任何影响，这就是值语义
}

func TestForRange3(t *testing.T) {
	five := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i := range five {
		fmt.Printf("%v\n", *(&five[i]))
	}
}
