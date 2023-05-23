package Order

import (
	"fmt"
)

var (
	a = c + b
	b = f()
	_ = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func PrintEvaluationOrder() {
	// 10 4 6 6
	fmt.Println(a, b, c, d)
}
