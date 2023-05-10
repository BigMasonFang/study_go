package main

import (
	// inner import
	// t "study_go/basic/type"
	// v "study_go/basic/variables"
	// s "study_go/basic/string_rune"
	// f "study_go/basic/funcs"
	// o "study_go/basic/operator"
	// logic "study_go/basic/if_else_switch"
	// loop "study_go/basic/for_loop"
	// innerInit "study_go/basic/init"
	// method "study_go/basic/method"
	// _struct "study_go/basic/struct"
	// array "study_go/basic/array"
	// s "study_go/basic/slice"
	// m "study_go/basic/map"
	i "study_go/basic/interface"
	// e "study_go/basic/error"
	// g "study_go/basic/generics"
	// routine "study_go/basic/goroutine"
	// r "study_go/basic/reflection"

	"fmt"

	// web import
	arraylist "github.com/emirpasic/gods/lists/arraylist" // alias + import

	// needed but not in main
	_ "github.com/emirpasic/gods/utils"
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
	// f.PrintVariadicFunc()

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
	// _struct.PrintStruct()

	// array
	// array.PrintArray()

	// slice
	// s.PrintSlice()

	// map
	// m.PrintMap()

	// interface
	i.PrintInterface()

	// routine
	// routine.PrintChannelWrite()
	// routine.PrintConcurrentCalSum()
	// routine.PrintErrorMsg()
	// routine.PrintCounter()
	// routine.PrintOrderPrint()
	// routine.PrintAtomic()
	// routine.PrintGen()
	// routine.PrintPipeline()
	// routine.PrintFanin()
	// routine.PrintClose()

	// error
	// e.PrintErrorAs()

	// generics
	// g.PrintGenric()

	// reflect
	// r.PrintReflect()
}
