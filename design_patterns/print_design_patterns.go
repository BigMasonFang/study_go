package design_patterns

import (
	"fmt"
	"study_go/design_patterns/creational_patterns"
)

func init() {
	designPatterns := []string{
		"singleton",
	}
	fmt.Println("u can select from these patterns")
	for i, v := range designPatterns {
		fmt.Println("input ", i, "or", v)
	}
}

func PrintDesignPatterns() {
	var input string
	fmt.Scanln(&input)

	switch input {
	case "0", "singleton":
		creational_patterns.PrintSingleton()
	}
}
