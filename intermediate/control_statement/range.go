package control

import (
	"fmt"
	"time"
)

func RangeScopeVar() {
	var a = [...]int{1, 2, 3, 4, 5}
	// array case
	for i, v := range a {
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println("in array for range i, v =", i, v)
			// opt i,v all 4, 5
			// Reason
			// 1. for range = do almost immediately
			//  but goroutine takes 3 secs to run, and
			// 2. the go routine uses outer scope var i and v, so
			// the i, v = reused in go routines inside the loop
		}()
		go func(i, v int) {
			// if u pass i,v as input param
			// will pass the scope problem and the opt will be allright
			time.Sleep(time.Second * 2)
			fmt.Println("in array for range i, v =", i, v)
		}(i, v)
	}
}

func RangeCopy() {
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int
	pToA, sliceToA := a, a
	rOfP, rOfSlice := r, r

	// range over array directly
	fmt.Println("init array a =", a)

	for i, v := range a {
		if i == 0 {
			a[1] = 22
			a[2] = 33
		}
		r[i] = v
	}
	fmt.Println("after range over to a directly")
	fmt.Println("a =", a)
	fmt.Println("r =", r)

	// range over pointer to array
	fmt.Println("init array pToA =", pToA)

	for i, v := range &pToA {
		if i == 0 {
			pToA[1] = 22
			pToA[2] = 33
		}
		rOfP[i] = v
	}
	fmt.Println("after range over to pointer of a")
	fmt.Println("pToA =", pToA)
	fmt.Println("rOfP =", rOfP)

	// range over slice
	fmt.Println("init array sliceToA =", sliceToA)

	for i, v := range sliceToA[:] {
		if i == 0 {
			sliceToA[1] = 22
			sliceToA[2] = 33
		}
		rOfSlice[i] = v
	}
	fmt.Println("after range over to slice of a")
	fmt.Println("sliceToA =", pToA)
	fmt.Println("rOfSlice =", rOfP)
}

func RangeLen() {
	// range over slice and dynamic append to it is useless
	// because it range will loop slice.Len times, which is
	// captured at the begining of the loop(fixed)
	s := []int{1, 2, 3, 4, 5}
	r := make([]int, 0)

	fmt.Println("init slice s =", s)

	for i, v := range s {
		if i == 0 {
			s = append(s, 6, 7)
		}
		r = append(r, v)
	}
	fmt.Println("after range over the slice when adding element to it inside loop")
	fmt.Println("s =", s)
	fmt.Println("r = ", r)
}

func RangeString() {
	s := "我爱你"
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}
}

func RangeMap() {
	// order is not garunteened
	// dynamic update is not garunteened

	for i := 0; i < 5; i++ {
		var m = map[string]int{
			"a": 1,
			"b": 2,
			"c": 3,
			"e": 5,
		}
		counter := 0
		for k, v := range m {
			if counter == 0 {
				delete(m, "a")
			}
			if counter == 3 {
				m["d"] = 4
			}
			counter++
			fmt.Println(k, v)
		}
		fmt.Println("counter is", counter)
	}
}

func RangeChan() {
	// recv from unbuffered channel
	var c = make(chan int)

	go func() {
		// beaware of the tricky order
		fmt.Println("sleep 1 seconds")
		time.Sleep(1 * time.Second)
		c <- 1

		fmt.Println("sleep another 1 seconds")
		time.Sleep(1 * time.Second)
		c <- 2

		fmt.Println("sleep final 1 seconds")
		time.Sleep(1 * time.Second)
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println("recv", v)
	}
}

func RangeBreak() {
	exit := make(chan interface{})

	go func() {
	loop: // label break's label
		for {
			select {
			case <-time.After(time.Second):
				fmt.Println("tick")
			// once recv exit signal, break
			case <-exit:
				fmt.Println("exiting...")
				// break // this break will only break select block
				break loop
				// this break [label] can determine which block to break
			}
		}
		fmt.Println("exit!")
	}()

	time.Sleep(3 * time.Second)
	// after 3 secs, send a exit signal to exit channel
	exit <- "exit"

	// wait child goroutine exit
	time.Sleep(3 * time.Second)
}

func PrintRange() {
	// RangeScopeVar()
	// RangeCopy()
	// RangeLen()
	// RangeString()
	// RangeMap()
	// RangeChan()
	RangeBreak()
}
