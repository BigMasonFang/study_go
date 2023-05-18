package composite

import "fmt"

func PrintArraySliceComposite() {
	numbers := [...]int{'a': 8, 'b': 7, 'c': 4, 'd': 3, 'e': 2, 'y': 1, 'x': 5}
	fnumbers := [...]float64{-1, 4: -0.5, -0.2, 7: -0.1, 9: 0}
	fmt.Println(numbers)
	fmt.Println(fnumbers)
}
