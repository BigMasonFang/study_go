package composite

import "fmt"

type MyStruct struct {
	name   string
	age    int
	sex    string
	weight float64
}

func PrintStructComposite() {
	// bad: build with values
	s1 := &MyStruct{"fx", 17, "M", 80.0}
	// good: build struct with filed-value
	// 1. field order can be arbitry
	// 2. not explicit field will initialized with zero value
	s2 := &MyStruct{age: 17}
	fmt.Println(s1)
	fmt.Println(s2)
}
