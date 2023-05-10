package goroutine

import "fmt"

func squre(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func PrintPipeline() {
	out := squre(gen(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	for i := 0; i < 9; i++ {
		fmt.Println(<-out)
	}
	fmt.Println(<-out)
}
