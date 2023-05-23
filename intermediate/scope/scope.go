package scope

import (
	"fmt"
	"sync"
	"time"
)

func ScopeExample() {
	// block 1
	type bar struct{} // type scope start here
	a, b := 0, 0

	{ // block 2
		a, b := 1, 1 // var a scope start here
		if a == 1 {
			// block 3
			a, b := 2, 2 // var
			fmt.Println("a in block 3 is", a)
			fmt.Println("b in block 3 is", b)
		}
		fmt.Println("a in block 2 is", a)
		fmt.Println("b in block 2 is", b)

	}
	fmt.Println("a in block 1 is", a)
	fmt.Println("b in block 1 is", b)
	// type scope start here
}

func IfBlockScope() {
	if a := 1; false {
		// nothing
	} else if b := 2; false {
		// nothing
	} else if c := 3; false {
		// nothing
	} else {
		fmt.Println(a, b, c)
	}
}

func ForBlockScope() {
}

func SwitchBlockScope() {
	switch x, y := 1, 2; x + y {
	case 3:
		a := 1
		fmt.Println("case 1: a = ", a)
		fallthrough
	case 10:
		a := 5
		fmt.Println("case 2: a = ", a)
		fallthrough
	default:
		a := 7
		fmt.Println("default case: a = ", a)
	}
}

func SelectBlockScope(step int) {
	c1 := make(chan int, 5)
	c2 := make(chan int, 1)
	var wg sync.WaitGroup

	switch step {
	case 0:
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				time.Sleep(1 * time.Second)
				fmt.Println("try to recv msg to c1", <-c1)
			}()
		}
	case 1:
		c2 <- 1
	default:
		fmt.Println("do nothing and wait for default case")
	}

	for i := 0; i < 5; i++ {
		select {
		case c1 <- i:
			fmt.Println("SendStmt case has been choosen")
		case i := <-c2:
			fmt.Println("RecvStmt case has been choosen")
			fmt.Println("Recv msg", i)
		default:
			fmt.Println("default case hase been choosen")
		}
	}

	wg.Wait()

	// close(c1)
}

func PrintScope() {
	// ScopeExample()
	SelectBlockScope(0)
}
