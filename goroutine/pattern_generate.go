package goroutine

import "fmt"

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func PrintGen() {
	out := gen(1, 2, 3, 7, 8, 9)
	for i := 0; i < 8; i++ {
		fmt.Println(<-out) // will output 0 two times cuz channel is closed
	}
}
