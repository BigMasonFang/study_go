package study_go

import (
	"fmt"
	"math"
)

func printName(name string) {
	fmt.Println(name)
}

func Sum(a, b float64) float64 {
	return a + b
}

func Squres(a, b float64) (float64, float64) {
	return math.Pow(a, 2), math.Pow(b, 2)
}

func PrintResult() {
	printName("fangxiang")
	fmt.Println(Sum(2, 3))
	fmt.Println(Squres(11, 17))
	x, y := 11.1, 17.1
	squareSum := Sum(Squres(x, y))
	fmt.Println(squareSum)
}
