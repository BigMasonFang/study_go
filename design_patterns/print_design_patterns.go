package design_patterns

import (
	"fmt"
	"study_go/design_patterns/creational_patterns"
	"study_go/design_patterns/structural_patterns"
)

func init() {
	designPatterns := []string{
		"singleton", "simple_factory", "abstract_factory", "builder",
		"configuration", "prototype", "adapter", "bridge", "object_tree",
		"decorator",
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
	case "1", "simple_factory":
		creational_patterns.PrintSimpleFactory()
	case "2", "abstract_factory":
		creational_patterns.PrintAbstractFactory()
	case "3", "builder":
		creational_patterns.PrintBuilder()
	case "4", "configuration":
		creational_patterns.PrintComplexConfig()
	case "5", "prototype":
		creational_patterns.PrintPrototype()
	case "6", "adapter":
		structural_patterns.PrintAdapter()
	case "7", "bridge":
		structural_patterns.PrintBridge()
		structural_patterns.PrintBridgeExec()
	case "8", "object_tree":
		structural_patterns.PrintObjectTree()
	case "9", "decorator":
		structural_patterns.PrintDecorator()
		structural_patterns.PrintDecoratorExercise()
	}
}
