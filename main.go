package main

import (
	s_u "github.com/BigMasonFang/study_go/string_rune"
	t "github.com/BigMasonFang/study_go/type"

	"github.com/davecgh/go-spew/spew"
	arraylist "github.com/emirpasic/gods/lists/arraylist" // alias + import
)

func main() {
	array1 := arraylist.New()
	array1.Add("a")
	array1.Add(1)
	// fmt.Printf("array1: %v\n", array1)
	spew.Println(array1.Values())
	spew.Dump(array1.Values()...)
	t.PrintType()
	s_u.PrintStringRune()
}
