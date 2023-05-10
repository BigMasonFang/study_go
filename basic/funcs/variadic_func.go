package funcs

import (
	"fmt"
)

// The function that is called with the varying number of arguments is known as variadic function.
func Find(num int, nums ...int) {
	fmt.Printf("the type of nums is %v\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func PrintVariadicFunc() {
	nums := []int{89, 90, 95}
	Find(89, nums...) // or
	Find(89, 90, 95, 89)
}

// fmt.Printf() append() are examples of variadic funcs
