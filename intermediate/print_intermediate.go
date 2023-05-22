package intermediate

import (
	"fmt"
	_map "study_go/intermediate/map"
	"study_go/intermediate/slice"
)

func init() {
	tips := []string{
		"slice",
		"map",
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
	}
}
