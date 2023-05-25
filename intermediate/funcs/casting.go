package funcs

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome\n")
}

func RunHTTPServer() {

	// check ListenAndServe, Handler, HandlerFunc source code

	// http.ListenAndServe
	// func ListenAndServe(addr string, handler Handler) error {
	// 	server := &Server{Addr: addr, Handler: handler}
	// 	return server.ListenAndServe()
	// }

	// http.Handler
	// type Handler interface {
	// 	ServeHTTP(ResponseWriter, *Request)
	// }

	// http.HandlerFunc()
	// type HandlerFunc func(ResponseWriter, *Request)

	// // ServeHTTP calls f(w, r).
	// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	// 	f(w, r)
	// }

	// in there http.HandlerFunc explicitly cast greeting to Handler type
	// the HandlerFunc wrap the greeting f with ServeHTTP func
	// so therefore it implement ServeHTTP and also became type Handler
	http.ListenAndServe(":8888", http.HandlerFunc(greeting))
}

// simpler example
type BinaryAdder interface {
	Add(x, y int) int
}

type MyAdderType func(int, int) int

func (f MyAdderType) Add(x, y int) int {
	return f(x, y)
}

func MyAdd(x, y int) int {
	return x + y
}

func PrintCasting() {
	var i BinaryAdder = MyAdderType(MyAdd)
	fmt.Println(i.Add(3, 4))
}
