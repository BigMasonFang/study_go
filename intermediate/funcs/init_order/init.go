package init_order

import (
	"fmt"
	_ "study_go/intermediate/funcs/init_order/pkg1"
	_ "study_go/intermediate/funcs/init_order/pkg3"
)

var (
	_  = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

const (
	c1 = "c1"
	c2 = "c2"
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 init")
	}
	if c2 != "" {
		fmt.Println("main: const c2 init")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s init\n", name)
	return name
}

func init() {
	fmt.Println("main: init")
}

func PrintInit() {
	fmt.Println("do something")
}
