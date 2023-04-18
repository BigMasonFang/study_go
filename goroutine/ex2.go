package goroutine

import (
	"fmt"
	"math"
	"sync"

	"github.com/davecgh/go-spew/spew"
)

// exercise 2
// concurrent cal the sum of a array
func conSum(s []int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel
}

func ConcurrentCalSum(s []int, i int) int {
	// i is the num of intervals of the sub array
	// and the routine num to open
	intervalLength := int(math.Round(float64(len(s)) / float64(i)))
	fmt.Println(intervalLength)

	c := make(chan int, i)
	sum := 0
	// create a wait group to wait for the goroutine
	var wg sync.WaitGroup

	for j := 0; j < i; j++ {
		wg.Add(1)
		if j != i-1 {
			go conSum(s[j*intervalLength:(j+1)*intervalLength], c, &wg)
		} else {
			// near end get all the rest
			go conSum(s[j*intervalLength:], c, &wg)
		}
		// fmt.Printf("indexs are [%v %v]\n", j*intervalLength, (j+1)*intervalLength)
		fmt.Println("open goroutine", j)
	}

	wg.Wait()

	spew.Dump(c)
	close(c) // close the channel to loop read from it
	for v := range c {
		fmt.Println(v)
		sum += v
	}

	return sum
}

func PrintConcurrentCalSum() {

	testArray := []int{4, 5, 19, 0, 57, 89, 6, 11, 12}
	sum := ConcurrentCalSum(testArray, 3)
	fmt.Println(sum)
}
