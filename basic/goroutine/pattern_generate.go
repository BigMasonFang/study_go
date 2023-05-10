package goroutine

import "fmt"

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("send %v to channel\n", n)
			out <- n
			fmt.Printf("%v sended\n", n)
		}
		close(out)
	}()
	return out
}

func PrintGen() {
	out := gen(1, 2, 3, 7, 8, 9)
	for i := 0; i < 8; i++ {
		fmt.Println("try to receive from channel")
		v := <-out
		fmt.Printf("read %v from channel\n", v) // will output 0 two times cuz channel is closed
	}
}

// it is key to understand
//  `Receive` operation on the channel is blocking and waits for a value to be available in the channel.

// order in line num, # is printed
// 20 -> 7# -> 22 -> 10 (23 is blocked cz no value in ch) -> 11 -> 12# -> 10# -> 23 (11 is blocked cz ch is full)
// -> 24#(read 1 from channel) -> 22#(try to receive from channel) -> 11!(23 is  blocked cz no value in ch) -> 23
// !!! channel
// -> 24#(read 2 from channel) -> 22# -> 12!# (2 sended)(23 is  blocked cz no value in ch) ->
// -> 10#(send 3 to channel) -> 11 -> 12# (3 sended) -> 10# (send 7 to channel) -> 23 -> 24# (read 3 from channel)
// -> 22#(try to receive from channel) -> 23 -> 24#(read 7 from channel) -> 22#(try to receive from channel)
// ->
