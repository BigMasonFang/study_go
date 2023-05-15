package structural_patterns

import (
	"fmt"
	"time"
)

type Timestamp int64

type MessageBody struct {
	// example
	data string
}

type Message struct {
	time   Timestamp // time stamp
	data   MessageBody
	result string
}

func filterUnexpiredMsg(msg <-chan Message, expireTime Timestamp) <-chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		for v := range msg {
			// fmt.Println(v.time)
			// fmt.Println(expireTime)
			if v.time > expireTime {
				out <- v
			}
		}
	}()
	return out
}

func filterMsgbyResult(msg <-chan Message, result string) <-chan Message {
	out := make(chan Message)
	go func() {
		defer close(out)
		for v := range msg {
			if v.result == result {
				out <- v
			}
		}
	}()
	return out
}

// func use pipeline pattern
func getUnexpiredErrorMsg(messages []Message, expireMin int) <-chan Message {
	// put it in channel
	process := make(chan Message)
	go func() {
		defer close(process)
		for _, v := range messages {
			process <- v
		}
	}()
	// close(process)

	expireTimestamp := Timestamp(time.Now().Add(-time.Minute * time.Duration(expireMin)).Unix())
	processBytime := filterUnexpiredMsg(process, expireTimestamp)
	processByfilter := filterMsgbyResult(processBytime, "ERROR")

	return processByfilter
	// return processBytime
	// return process
}

func PrintPipeline() {
	data := []Message{
		{1684141567, MessageBody{"test 1"}, "SUCCESS"},
		{1684141568, MessageBody{"test 2"}, "SUCCESS"},
		{1684141569, MessageBody{"test 3"}, "SUCCESS"},
		{1684141570, MessageBody{"test 4"}, "ERROR"}, // will show
		{1684141571, MessageBody{"test 5"}, "SUCCESS"},
		{1684141571, MessageBody{"test 6"}, "ERROR"}, // will show
		{1684141573, MessageBody{"test 7"}, "SUCCESS"},
		{1684147967, MessageBody{"test 8"}, "SUCCESS"},
		{1684147968, MessageBody{"test 9"}, "ERROR"},
		{1684147969, MessageBody{"test 10"}, "SUCCESS"},
	}
	// var wg sync.WaitGroup
	// wg.Add(1)
	processedCh := getUnexpiredErrorMsg(data, 60)
	// wg.Wait()
	for v := range processedCh {
		fmt.Println(v)
	}
}
