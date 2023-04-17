package goroutine

import (
	"fmt"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

var _time time.Time

// exercise 1
func Printsome() {
	_time = time.Now()
	fmt.Println("current time:", _time, "PrintChannelWrite is going")
}

func ChannelWrite(ch chan int, wg *sync.WaitGroup, x int) {
	defer wg.Done()
	for x > 0 {
		_time = time.Now()
		if len(ch) == cap(ch) {
			fmt.Println("current time:", _time, "beyond buffer size, can not send to channel")
		}
		// may jamed in here
		ch <- x
		_time = time.Now()
		fmt.Printf("current time: %v , receive %v, plz input after 2 secondes\n", _time, x)
		time.Sleep(time.Second * 2)
		x--
	}
}

func PrintChannelWrite() {
	// create a buffered channel, can store 2 int
	ch := make(chan int, 2)

	// create a wait group to wait for the goroutine
	var wg sync.WaitGroup
	// add 1 to the wait group counter
	// in this case, the next go routine in line 34 is represented
	wg.Add(1)

	// create the go routine
	go ChannelWrite(ch, &wg, 4)

	time.Sleep(time.Second * 1)

	for i := 0; i < 3; i++ {
		go Printsome()
	}

	time.Sleep(time.Second * 5)

	for i := 0; i < 3; i++ {
		fmt.Println("PrintChannelWrite is going")
	}
	fmt.Printf("current time is %v and get value %v from ch\n", time.Now(), <-ch)
	time.Sleep(time.Second * 1)

	for i := 0; i < 3; i++ {
		fmt.Println("PrintChannelWrite is going")
	}
	fmt.Printf("current time is %v and get value %v from ch\n", time.Now(), <-ch)
	time.Sleep(time.Second * 1)

	for i := 0; i < 3; i++ {
		fmt.Println("PrintChannelWrite is going")
	}
	fmt.Printf("current time is %v and get value %v from ch\n", time.Now(), <-ch)

	// wait group wait untile all counter to 0
	wg.Wait()
	fmt.Println("current time is", time.Now())
	spew.Dump(ch) // still remains 1

	// how this run?
	/* main() -> PrintChannelWrite() ->

	{} means concurrently
	{
		routine 1 ChannelWrite() in line 41 wirite 4 to in ch, starts to sleep for 2 sleep
		main routine sleep 1 second in line 43
	}

	1 seconds passed ->
	{
		main routine goese to line 45 loop and -> routine 2 Printsome() in line 44 finish immediately
		routine 1 ChannelWrite() sleept for 1 second, 1 second to go
	}

	->
	main routine starts starts sleep for 5 seconds in line 49

	1 seconds passed ->
	{
		ch receive 3, and starts to sleep for 2 seconds
		main routine 4 seconds to sleep
	}

	2 seconds passed->
	{
		routine 1 because channel ([4,3]) is full, jammed, print beyond buffer warning
	    main routine 2 seconds to sleep
	}

	2 seconds passed ->
	{
		main routine sleep over ->
		print PrintChannelWrite is going in line 51, finish immediately ->
		get 1 value (4) from ch[4,3] in line 54 and print it ->
		{
			channel not jammed ([3]), routine 1 continue, receive (2), channel full ([3,2]) ->
			main routine starts to sleep for 1 seconds in line 55
			routine 1 starts to sleep for 2 seconds
		}
	}

	1 seconds passed ->
	{
		main routine sleep over ->
		print PrintChannelWrite is going in line 57, finish immediately ->
		get 1 value (3) from ch[3,2] in line 60 and print it ->
		{
			channel not jammed ([2]), routine 1 still 1 seconds to sleep
			main routine starts to sleep for 1 seconds in line 61
		}
	}

	1 seconds passed ->
	{
		routine 1 sleep over, ch not jamed, receive 1, ch ([2,1])
		routine 1 start to sleep for 2 seconds
	 	main routine sleep over ->
		print PrintChannelWrite is going in line 63, finish immediately ->
		get 1 value (2) from ch[2，1] in line 66 and print it ->
	}

	2 seconds passed
	routine 1 over
	main routine over
	*/
}
