// Adapater pattern is used to enable the interaction between two incompatible interfaces.
// The Adapter pattern acts as a bridge between two interfaces to help them work together.
// This pattern allows existing code to work with new code by ensuring that
// the interfaces are kept separate while providing a way for them to communicate with each other.
package structural_patterns

import "fmt"

// Adaptee is the interface needs to be adapted
type Adaptee interface {
	SpecificRequest() string
}

// AdapteeImpl is the concrete implement of Adaptee
type AdapteeImpl struct{}

func (a *AdapteeImpl) SpecificRequest() string {
	return "specific request"
}

// Target is the target interface we want to adapt to
type Target interface {
	Request() string
}

// Adapter is the adapter that adapts Adaptee to Target
type Adapter struct {
	Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

// client is the consumer of Target interface
func Client(t Target) string {
	return t.Request()
}

func PrintAdapter() {
	adaptee := &AdapteeImpl{}
	target := &Adapter{Adaptee: adaptee}
	output := Client(target)
	fmt.Println(output)

}
