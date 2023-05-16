// 1. Concurrency: Channels are an efficient way of implementing concurrency in Golang.
// It allows multiple goroutines to send and receive data asynchronously without any race conditions.

// 2. Synchronization: Channels enforce synchronization and ensure that data is sent
// and received in a synchronized and safe way.

// 3. Simplicity: The channel pattern is simple and easy to use, even for developers who are new to Golang.
// This is because channels are a built-in feature of the language, making it easy to use and maintain.

// 4. Memory Management: Channels help to manage memory efficiently by allowing goroutines to communicate
// and share data without creating unnecessary copies or leaks.

// 5. Scalability: Channels allow for scalable development by allowing developers to easily manage
// goroutines and scale the application as the number of users grows.

// Channel Closing Principle
// don't close a channel from the receiver side and don't close a channel if the channel has multiple concurrent senders.

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

func buildTimestamp(interval int) Timestamp {
	// build timestamp which is interval minutes from now
	return Timestamp(time.Now().Add(-time.Minute * time.Duration(interval)).Unix())
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

func filterMsgByResult(msg <-chan Message, result string) <-chan Message {
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

func chfilterUnexpiredMsg(ch chan Message, resultCh chan Message, expireTime Timestamp) {
	defer close(resultCh)
	for v := range ch {
		if v.time > expireTime {
			// fmt.Println(v)
			resultCh <- v
		}
	}
}

func chfilterByResult(ch chan Message, resultCh chan Message, result string) {
	defer close(resultCh)
	for v := range ch {
		if v.result == result {
			// fmt.Println(v)
			resultCh <- v
		}
	}
}

func Intersect(chs ...<-chan Message) <-chan Message {
	result := make(chan Message)
	visited := make(map[Message]int)

	for _, ch := range chs {
		go func(ch <-chan Message) {
			for v := range ch {
				visited[v]++
			}
		}(ch)
	}

	go func() {
		for message, v := range visited {
			if v == len(chs) {
				result <- message
			}
		}
		// close(result)
	}()

	return result
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

	expireTimestamp := buildTimestamp(expireMin)
	processBytime := filterUnexpiredMsg(process, expireTimestamp)
	processByfilter := filterMsgByResult(processBytime, "ERROR")

	return processByfilter
}

// concurrent
func getUnexpiredErrorMsgConcurrent(messages []Message, expireMin int) <-chan Message {
	// put it in channel
	process := make(chan Message)
	// wg.Add(1)

	go func() {
		defer close(process)
		for _, v := range messages {
			process <- v
		}
	}()

	expireTimestamp := buildTimestamp(expireMin)
	processBytime := make(chan Message)
	processBytimeAndResult := make(chan Message)

	go chfilterUnexpiredMsg(process, processBytime, expireTimestamp)
	go chfilterByResult(processBytime, processBytimeAndResult, "ERROR")

	// result := Intersect(&wg, resultCh1, resultCh2)
	return processBytimeAndResult
}

var testData []Message = []Message{
	{buildTimestamp(70), MessageBody{"test 1"}, "SUCCESS"},
	{buildTimestamp(65), MessageBody{"test 2"}, "ERROR"},
	{buildTimestamp(64), MessageBody{"test 3"}, "SUCCESS"},
	{buildTimestamp(59), MessageBody{"test 4"}, "ERROR"}, // will show
	{buildTimestamp(54), MessageBody{"test 5"}, "SUCCESS"},
	{buildTimestamp(30), MessageBody{"test 6"}, "ERROR"}, // will show
	{buildTimestamp(25), MessageBody{"test 7"}, "SUCCESS"},
	{buildTimestamp(24), MessageBody{"test 8"}, "SUCCESS"},
	{buildTimestamp(15), MessageBody{"test 9"}, "ERROR"}, // will show
	{buildTimestamp(5), MessageBody{"test 10"}, "SUCCESS"},
}

func PrintPipeline() {
	processedCh := getUnexpiredErrorMsg(testData, 60)
	for v := range processedCh {
		fmt.Println(v)
	}
}

func PrintFanInIntersect() {
	processedCh := getUnexpiredErrorMsgConcurrent(testData, 60)
	for v := range processedCh {
		fmt.Println(v)
	}
}
