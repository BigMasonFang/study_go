// 1. reduce coupling
// 2. flexiable and easy to maintain (dynamic add rm arrange handles)
// 3. easy to expand
// 4. if too long may be too complex

package behavioral

import "fmt"

type Handler interface {
	SetNext(Handler) Handler
	HandleRequest(string)
}

type ConcreteHandler1 struct {
	next Handler
}

func (h *ConcreteHandler1) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *ConcreteHandler1) HandleRequest(request string) {
	if request == "GET" {
		fmt.Printf("ConcreteHandler 1 handle %s request\n", request)
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

type ConcreteHandler2 struct {
	next Handler
}

func (h *ConcreteHandler2) SetNext(next Handler) Handler {
	h.next = next
	return next
}

func (h *ConcreteHandler2) HandleRequest(request string) {
	if request == "POST" {
		fmt.Printf("ConcreteHandler 2 handle %s request\n", request)
	} else if h.next != nil {
		h.next.HandleRequest(request)
	}
}

func PrintChainOfResponsibility() {
	h1 := &ConcreteHandler1{}
	h2 := &ConcreteHandler2{}

	h1.SetNext(h2)

	h1.HandleRequest("POST")
	h1.HandleRequest("GET")
}
