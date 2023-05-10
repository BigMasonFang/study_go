package goroutine

import (
	"fmt"
	"math/rand"
	"time"
)

var testStringArray []string = []string{"xx", "xx", "error"}

type Info struct {
	id   int
	msg  string
	data map[string]int
}

func sendInfo(i int, c chan Info) {
	// i is routine index
	for {
		c <- Info{rand.Intn(1001), testStringArray[rand.Intn(len(testStringArray))], map[string]int{"a": 1}}
		time.Sleep(time.Second)
	}
}

func PrintErrorMsg() {

	c1 := make(chan Info)
	c2 := make(chan Info)
	c3 := make(chan Info)

	go sendInfo(1, c1)
	go sendInfo(2, c2)
	go sendInfo(3, c3)

	for {
		select {
		case info1 := <-c1:
			if info1.msg == "error" {
				fmt.Printf("erro info from c1 with id %v\n", info1.id)
			}
		case info2 := <-c2:
			if info2.msg == "error" {
				fmt.Printf("erro info from c2 with id %v\n", info2.id)
			}
		case info3 := <-c3:
			if info3.msg == "error" {
				fmt.Printf("erro info from c3 with id %v\n", info3.id)
			}
		}
	}
}
