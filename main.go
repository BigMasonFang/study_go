package main

import (
	"fmt"
	basic "study_go/basic"
	_ "study_go/design_patterns/creational_patterns"
)

func init() {
	// initialization before main
	MAIN_ENV := "DEV"
	fmt.Println("initialization in env ", MAIN_ENV)
}

func main() {
	// basic
	basic.PrintBasic()
}
