package study_go

import "fmt"

var ENV = "INIT_DEV"
var Init_var = "something"

func init() {
	fmt.Println("init inner mod env is", ENV)
}
