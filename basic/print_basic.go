package basic

import (

	// inner import
	array "study_go/basic/array"
	e "study_go/basic/error"
	loop "study_go/basic/for_loop"
	f "study_go/basic/funcs"
	g "study_go/basic/generics"
	routine "study_go/basic/goroutine"
	logic "study_go/basic/if_else_switch"
	innerInit "study_go/basic/init"
	i "study_go/basic/interface"
	m "study_go/basic/map"
	method "study_go/basic/method"
	o "study_go/basic/operator"
	r "study_go/basic/reflection"
	s "study_go/basic/slice"
	str "study_go/basic/string_rune"
	_struct "study_go/basic/struct"
	t "study_go/basic/type"
	v "study_go/basic/variables"

	"fmt"

	// web import
	_ "github.com/emirpasic/gods/lists/arraylist" // alias + import
	// needed but not in main
	_ "github.com/emirpasic/gods/utils"
)

func init() {
	chapters := []string{
		"type", "variables", "string_rune", "operator", "if_else_switch",
		"for_loop", "init", "method", "struct", "array", "slice", "map",
		"interface", "error", "generics", "goroutine", "reflection", "funcs"}
	fmt.Println("u can select from these chapters")
	for i, v := range chapters {
		fmt.Println("input ", i, "or", v)
	}
}

func PrintBasic() {
	var input string
	fmt.Scanln(&input)
	switch input {
	case "type", "0":
		t.PrintType()
	case "variables", "1":
		v.PrintVars()
		v.PrintResult()
	case "string_rune", "2":
		str.PrintStringRune()
	case "operator", "3":
		o.PrintOperator()
	case "if_else_switch", "4":
		logic.PrintLogicControl()
	case "for_loop", "5":
		loop.PrintLoop()
	case "init", "6":
		innerInit.Sleep1s()
		// see when will the init print
	case "method", "7":
		method.PrintMethod()
	case "struct", "8":
		_struct.PrintStruct()
	case "array", "9":
		array.PrintArray()
	case "slice", "10":
		s.PrintSlice()
	case "map", "11":
		m.PrintMap()
	case "interface", "12":
		i.PrintInterface()
	case "error", "13":
		e.PrintErrorAs()
		e.PrintErrorIs()
	case "generics", "14":
		g.PrintGenric()
	case "goroutine", "15":
		routine.PrintAtomic()
		routine.PrintChannelWrite()
		routine.PrintClose()
		routine.PrintConcurrentCalSum()
		routine.PrintCounter()
		routine.PrintErrorMsg()
		routine.PrintFanin()
		routine.PrintGen()
		routine.PrintOrderPrint()
		routine.PrintPipeline()
		routine.Printsome()
	case "funcs", "16":
		f.PrintResult()
		f.PrintVariadicFunc()
	case "reflection", "17":
		r.PrintReflect()
	}
}
