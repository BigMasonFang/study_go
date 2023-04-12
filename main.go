package main

import (
	// inner import
	// t "github.com/BigMasonFang/study_go/type"
	// v "github.com/BigMasonFang/study_go/variables"
	// s "github.com/BigMasonFang/study_go/string_rune"
	// f "github.com/BigMasonFang/study_go/funcs"
	// o "github.com/BigMasonFang/study_go/operator"
	// logic "github.com/BigMasonFang/study_go/if_else_switch"
	// loop "github.com/BigMasonFang/study_go/for_loop"
	// innerInit "github.com/BigMasonFang/study_go/init"
	// method "github.com/BigMasonFang/study_go/method"
	_struct "github.com/BigMasonFang/study_go/struct"

	"fmt"

	// web import
	arraylist "github.com/emirpasic/gods/lists/arraylist" // alias + import
)

func init() {
	// initialization before main
	MAIN_ENV := "DEV"
	fmt.Println("initialization in env ", MAIN_ENV)
}

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
	// v.PrintResult()

	// funcs
	// f.PrintResult()

	// operators
	// o.PrintOperator()

	// if else swith
	// logic.PrintLogicControl()

	// loop
	// loop.PrintLoop()

	// init
	// fmt.Println(innerInit.Init_var)

	// method
	// method.PrintMethod()

	// struct
	_struct.PrintStruct()
}
