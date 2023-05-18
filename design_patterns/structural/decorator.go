// 1. allows behavior to be added to an individual object dynamically
// without affecting the behavior of other objects from the same class
// 2. easy to maintain because different it is componentlized
// 3. easy to reuse code because different it is componentlized

package structural

import "fmt"

// abstract
type Component interface {
	Operation() string
}

// struct which utlize Component interface
type ConcreteComponent struct{}

func (c *ConcreteComponent) Operation() string {
	return "ConcreteComponent"
}

// decorator is a wrapper func, which take component as input
// param and return a struct with additional type
// abstract
type Decorator func(Component) Component

// concrete
type DecoratorA struct {
	component Component
}

func (d DecoratorA) Operation() string {
	return fmt.Sprintf("ConcreteDecoratorA(%s)", d.component.Operation())
}

func ConcreteDecoratorA(c Component) Component {
	return &DecoratorA{component: c}
}

// concrete
type DecoratorB struct {
	component Component
}

func (d DecoratorB) Operation() string {
	defer fmt.Println("last bread")
	fmt.Println("top bread")
	return fmt.Sprintf("ConcreteDecoratorB(%s)", d.component.Operation())
}

func ConcreteDecoratorB(c Component) Component {
	return &DecoratorB{component: c}
}

func PrintDecorator() {
	component := &ConcreteComponent{}

	componentA := ConcreteDecoratorA(component)
	componentB := ConcreteDecoratorB(component)

	fmt.Println(component.Operation())
	fmt.Println(componentA.Operation())
	opB := componentB.Operation()
	fmt.Println(opB)
}
