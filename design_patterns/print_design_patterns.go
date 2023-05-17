package design_patterns

import (
	"fmt"
	"study_go/design_patterns/behavioral_patterns"
	"study_go/design_patterns/creational_patterns"
	"study_go/design_patterns/structural_patterns"
)

func init() {
	designPatterns := []string{
		"singleton", "simple_factory", "abstract_factory", "builder",
		"configuration", "prototype", "adapter", "bridge", "object_tree",
		"decorator", "pipeline", "plugin", "fluent_interface", "chain_of_responsibility",
		"observer", "cache_proxy", "strategy", "templates",
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
	case "10", "pipeline":
		structural_patterns.PrintPipeline()
		structural_patterns.PrintFanInIntersect()
	case "11", "plugin":
		structural_patterns.PrintPlugin()
	case "12", "fluent_interface":
		behavioral_patterns.PrintFluentInterface()
	case "13", "chain_of_responsibility":
		behavioral_patterns.PrintChainOfResponsibility()
	case "14", "observer":
		behavioral_patterns.PrintObserver()
	case "15", "cache_proxy":
		behavioral_patterns.PrintCacheProxy()
	case "16", "strategy":
		behavioral_patterns.PrintStrategy()
	case "17", "templates":
		behavioral_patterns.PrintTemplates()
	}
}
