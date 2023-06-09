package generics

import "fmt"

type Number interface {
	int64 | float64
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

func SumIntsorFloats[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func PrintGenric() {
	ints := map[string]int64{
		"first":  22,
		"second": 44,
	}
	floats := map[string]float64{
		"first":  22.3,
		"second": 87.1,
	}
	fmt.Printf("Generic Sums: %v and %v\n",
		// can write with type
		SumIntsorFloats[string, int64](ints),
		// or auto inference
		SumIntsorFloats(floats))
}
