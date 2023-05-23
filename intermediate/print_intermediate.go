package intermediate

import (
	"fmt"
	_map "study_go/intermediate/map"
	Order "study_go/intermediate/order"
	"study_go/intermediate/scope"
	"study_go/intermediate/slice"
	_string "study_go/intermediate/string"
)

func init() {
	tips := []string{
		"slice",
		"map",
		"string",
		"evaluation_order",
		"scope",
	}
	fmt.Println("u can select from these intermediate topics")
	for i, v := range tips {
		fmt.Println("input ", i, "or", v)
	}
}

func PrintTopics() {
	var input string
	fmt.Scanln(&input)

	switch input {
	case "0", "slice":
		slice.PrintDynamicGrow()
	case "1", "map":
		_map.PrintMap()
	case "2", "string":
		_string.PrintString()
	case "3", "evaluation":
		Order.PrintEvaluationOrder()
		Order.PrintUsualOrder()
		Order.PrintSwitchExpr()
		Order.PrintSelectExpr()
	case "4", "scope":
		scope.PrintScope()
	}
}
