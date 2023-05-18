package design_patterns

import (
	"fmt"
	"study_go/design_patterns/behavioral"
	"study_go/design_patterns/creational"
	"study_go/design_patterns/structural"
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
		creational.PrintSingleton()
	case "1", "simple_factory":
		creational.PrintSimpleFactory()
	case "2", "abstract_factory":
		creational.PrintAbstractFactory()
	case "3", "builder":
		creational.PrintBuilder()
	case "4", "configuration":
		creational.PrintComplexConfig()
	case "5", "prototype":
		creational.PrintPrototype()
	case "6", "adapter":
		structural.PrintAdapter()
	case "7", "bridge":
		structural.PrintBridge()
		structural.PrintBridgeExec()
	case "8", "object_tree":
		structural.PrintObjectTree()
	case "9", "decorator":
		structural.PrintDecorator()
		structural.PrintDecoratorExercise()
	case "10", "pipeline":
		structural.PrintPipeline()
		structural.PrintFanInIntersect()
	case "11", "plugin":
		structural.PrintPlugin()
	case "12", "fluent_interface":
		behavioral.PrintFluentInterface()
	case "13", "chain_of_responsibility":
		behavioral.PrintChainOfResponsibility()
	case "14", "observer":
		behavioral.PrintObserver()
	case "15", "cache_proxy":
		behavioral.PrintCacheProxy()
	case "16", "strategy":
		behavioral.PrintStrategy()
	case "17", "templates":
		behavioral.PrintTemplates()
	}
}
