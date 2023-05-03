package goroutine

import (
	"fmt"
	"sync"
)

func receive(sendChannel <-chan int, outputChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range sendChannel {
		outputChannel <- n
	}
}

func merge(sendChannel ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(sendChannel))
	for _, c := range sendChannel {
		go receive(c, out, &wg)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func PrintFanin() {
	c1 := gen(1, 2, 3)
	c2 := gen(4, 5, 6)
	c3 := gen(7, 8, 9)
	out := merge(c1, c2, c3)
	// for i := 0; i < 9; i++ {
	// 	fmt.Println(<-out)
	// }
	for e := range out {
		fmt.Println(e)
	}
	fmt.Println(<-out)
}
