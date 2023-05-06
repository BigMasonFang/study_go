package goroutine

import (
	"context"
	"fmt"
	"time"
)

type InChan <-chan time.Time

type Reciver struct {
	ctx     context.Context
	cancel  context.CancelFunc
	outChan chan string
	inChan  InChan
}

func (r *Reciver) receiveWithClose() {
	for {
		select {
		// 这个写法可以获取到inChan的数据
		case msgIn := <-r.inChan:
			// t = <-r.inChan
			r.outChan <- "msg + " + msgIn.String()
		// 这个没有获取
		case <-r.ctx.Done():
			r.outChan <- "close"
			close(r.outChan)
			return
		}
	}
}

func (r *Reciver) genWithClose(n int) {

	go r.receiveWithClose()

	go func() {
		fmt.Println("start to context")
		time.Sleep(time.Duration(n) * time.Second)
		r.cancel()
	}()
}

// 发送方发送完数据可以选择主动guanbi, 但如果有的时候需要
// 接收完数据后再关闭管道
// 使用context
func PrintClose() {
	ctx, cancel := context.WithCancel(context.Background())
	out := make(chan string)
	tick := time.Tick(time.Second)
	reciver := Reciver{ctx, cancel, out, tick}

	reciver.genWithClose(5)

	for v := range reciver.outChan {
		fmt.Println(v)
	}
}
