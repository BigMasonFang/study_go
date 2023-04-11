package main

import (
	// inner import

	v "github.com/BigMasonFang/study_go/variables"

	// web import
	arraylist "github.com/emirpasic/gods/lists/arraylist" // alias + import
)

func main() {
	// mod and package
	array1 := arraylist.New()
	array1.Add("a")
	// array1.Add(1)
	// spew.Println(array1.Values())
	// spew.Dump(array1.Values()...)

	// type
	// t.PrintType()
	// string&rune
	// s_u.PrintStringRune()
	// variables
	// v.PrintVars()
	v.PrintResult()
}
