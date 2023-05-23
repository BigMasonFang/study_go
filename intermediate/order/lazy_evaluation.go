package Order

import (
	"fmt"
	"time"
)

func Expr(n int) int {
	fmt.Println("expr", n)
	return n
}

func PrintSwitchExpr() {
	switch Expr(2) {
	case Expr(1), Expr(2), Expr(3):
		fmt.Println("enter into case 1")
		fallthrough
	case Expr(4):
		fmt.Println("enter into case 2")
	}
}

func getAReadOnlyChannel() <-chan int {
	fmt.Println("invoke getAReadOnlyChannel")
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		c <- 1
	}()

	return c
}

func getASlice() *[5]int {
	fmt.Println("invoke getASlice")
	var a [5]int
	return &a
}

func getAWriteOnlyChannel() chan<- int {
	fmt.Println("invoke getAWriteOnlyChannel")
	return make(chan int)
}

func getAIntToChannel() int {
	fmt.Println("invoke getAIntTOChannel")
	return 2
}

func PrintSelectExpr() {

	select {
	case getASlice()[0] = <-getAReadOnlyChannel():
		fmt.Println("recv something from a readonly channel")
	case getAWriteOnlyChannel() <- getAIntToChannel():
		fmt.Println("recv something from a writeonly channel")
	}

	// invoke getAReadOnlyChannel
	// invoke getAWriteOnlyChannel
	// invoke getAIntTOChannel
	// invoke getASlice
	// recv something from a readonly channel
}
