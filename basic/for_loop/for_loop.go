package for_loop

import "fmt"

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	result := n
	for n > 1 {
		n--
		result *= n
	}
	return result
}

func factorial2(n int) int {
	result := 1
	if n < 0 {
		return 0
	}
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}

func PrintLoop() {
	fmt.Println(factorial(5))
	fmt.Println(factorial2(5))
}
