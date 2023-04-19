package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func ConcurrentAdd(n int, counter *int64, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < n; j++ {
		atomic.AddInt64(counter, 1)
	}
}

func ConcurrentMultiply(x, y int) int64 {
	var counter int64 = 0
	var wg sync.WaitGroup
	for i := 0; i < x; i++ {
		wg.Add(1)
		go ConcurrentAdd(y, &counter, &wg)
		// fmt.Println(atomic.LoadInt64(&counter))
	}
	wg.Wait()
	return atomic.LoadInt64(&counter)
}

func PrintAtomic() {
	fmt.Println("concurrent cal 17 * 95", ConcurrentMultiply(17, 95))
}
