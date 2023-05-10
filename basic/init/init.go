package init

import (
	"fmt"
	"time"
)

var ENV = "INIT_DEV"
var Init_var = "something"

func init() {
	fmt.Println("init inner mod env is", ENV)
}

func Sleep1s() {
	fmt.Println("start to sleep")
	time.Sleep(time.Second * 1)
	fmt.Println("sleep over")
}
