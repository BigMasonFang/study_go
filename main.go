package main

import (
	"fmt"
	// "sync"
	// "net/http"
	// basic "study_go/basic"
	// patterns "study_go/design_patterns"
	// tips "study_go/tips"
	intermediate "study_go/intermediate"
)

func init() {
	// initialization before main
	MAIN_ENV := "DEV"
	fmt.Println("initialization in env ", MAIN_ENV)
}

func main() {
	// basic
	// basic.PrintBasic()

	// patterns
	// patterns.PrintDesignPatterns()

	// http.ListenAndServe(":8888", nil)

	// tips.PrintTips()

	intermediate.PrintTopics()

}
