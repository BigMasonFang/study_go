package string

import "fmt"

// immutable
func PrintImmutable() {
	var s string = "hello"
	fmt.Println("original string", s)

	// try to change it by reslice
	sl := []rune(s)
	sl[0] = 'z'
	fmt.Println("slice is", string(sl))
	fmt.Println("after reslice it, the original string is", string(s))
}

func PrintNilUseable() {
	var s string = ""
	fmt.Println(s)
	fmt.Println(len(s))
}

func PrintOperators() {
	a := "ABC"
	b := "D"
	d := "ABD"
	fmt.Println(a + b)

	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(a >= b)
	fmt.Println(a <= d)
	fmt.Println(a < b)
	fmt.Println(a != b)
}

func PrintUnicode() {
	s := "这是一个很困难的抉择，后面的道路也很艰辛，但前途是光明的"
	rS := []rune(s)
	bS := []byte(s)
	fmt.Println(len(rS))
	fmt.Println(len(bS))

	for i, v := range rS {
		var utf8Bytes []byte
		for j := i * 3; j < (i+1)*3; j++ {
			utf8Bytes = append(utf8Bytes, bS[j])
		}
		// `%X` is a formatting directive that is used to format an integer value as a hexadecimal string in uppercase.
		fmt.Printf("%s => %X => %X\n", string(v), v, utf8Bytes)
	}
}

func PrintMultipleLines() {
	const s = `
	这是一首诗
	I love pantera
	`
	fmt.Println(s)
}

func PrintString() {
	PrintImmutable()
	PrintNilUseable()
	PrintOperators()
	PrintUnicode()
}
