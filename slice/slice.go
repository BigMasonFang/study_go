package slice

import "fmt"

func PrintSlice() {
	// init
	s1 := make([]int, 5)
	s2 := make([]int, 3, 5)
	// modify value of element
	s1[0] = 3
	fmt.Println(s1)
	// copy
	s3 := make([]int, 3)
	copy(s3, s1[0:3])
	fmt.Println(s1, s2, s3)
}
