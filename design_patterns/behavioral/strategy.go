// 1. easy to expand
// 2. decoupling strategy's implementation and useage, which can
// allow switch strategies without modifying client's code
// 3. reuse, same as above, encapsulation strateies which can be
// used in multiple scenarios, witout implementation in client
// 4. enhance maintainability by making single duty strategies

package behavioral

import "fmt"

type Strategy interface {
	Excute(a, b int) int
	Name() string
}

type AddStrategy struct{}

func (add *AddStrategy) Excute(a, b int) int {
	return a + b
}

func (add *AddStrategy) Name() string {
	return "add"
}

type SubtractStrategy struct{}

func (sub *SubtractStrategy) Excute(a, b int) int {
	return a - b
}

func (sub *SubtractStrategy) Name() string {
	return "subtract"
}

// or called Strategier
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.strategy = s
}

func (c *Context) ExcuteStrategy(a, b int) int {
	return c.strategy.Excute(a, b)
}

func PrintStrategy() {
	strategier := &Context{}
	add := &AddStrategy{}
	subtract := &SubtractStrategy{}

	strategier.SetStrategy(add)
	result := strategier.ExcuteStrategy(5, 4)
	fmt.Printf("use <%s> strategy, result is %d\n", strategier.strategy.Name(), result)

	strategier.SetStrategy(subtract)
	result = strategier.ExcuteStrategy(5, 4)
	fmt.Printf("use <%s> strategy, result is %d\n", strategier.strategy.Name(), result)
}
