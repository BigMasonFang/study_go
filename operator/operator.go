package operator

import (
	"fmt"
)

func PrintOperator() {
	fmt.Println(3 + 3) // arithmetic
	a, b := 3, 2
	a += 3 // arithmetic assign
	fmt.Println(a)
	a++ // increment
	fmt.Println(a)
	a-- // decrement
	fmt.Println(a)

	fmt.Println(a >= 2)
	fmt.Println(a <= b*2)
	fmt.Println(a <= b*10)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!false)
}
