package goroutine

import (
	"fmt"
	"sync"
)

type routinePrintFunc func(wg *sync.WaitGroup, i int, s string)

// should place wg inside routine so that it can be done
func routinePrint(wg *sync.WaitGroup, i int, s string) {
	defer wg.Done()
	fmt.Printf("routine %v start\n", i)
	fmt.Printf("routine %v print something %v\n", i, s)
	fmt.Printf("routine %v end\n", i)
}

func OrderPrint(fn routinePrintFunc, n int) {
	var wg sync.WaitGroup
	// defer wg.Wait()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go routinePrint(&wg, i, "print something")
	}
	wg.Wait()
}

func PrintOrderPrint() {
	OrderPrint(routinePrint, 5)
	fmt.Println("All go routine finished")
}
