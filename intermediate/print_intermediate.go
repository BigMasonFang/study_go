package intermediate

import (
	"fmt"
	"study_go/intermediate/slice"
)

func init() {
	tips := []string{
		"slice",
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
	}
}
