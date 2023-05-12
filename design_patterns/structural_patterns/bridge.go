// Bridge is a structural is used to separate an abstraction from its implementation
// so that both can be developed independently. In other words,
// it decouples an abstraction from its implementation details
// so that the two can vary independently.
package structural_patterns

import "fmt"

// Abstraction shape
type Shape interface {
	Draw()
}

// Implementation
type Color interface {
	ApplyColor() string
}

type RedColor struct{}

func (r *RedColor) ApplyColor() string {
	return "Red"
}

type BlueColor struct{}

func (r *BlueColor) ApplyColor() string {
	return "Blue"
}

// bridge part combine shape and color
type Circle struct {
	color Color
}

func (c *Circle) Draw() {
	fmt.Printf("Drawing a %s circle\n", c.color.ApplyColor())
}

type Square struct {
	color Color
}

func (c *Square) Draw() {
	fmt.Printf("Drawing a %s square\n", c.color.ApplyColor())
}

func PrintBridge() {
	red := &RedColor{}
	blue := &BlueColor{}

	circleRed := &Circle{red}
	circleBlue := &Circle{blue}

	squareRed := &Square{red}
	squareBlue := &Square{blue}

	circleRed.Draw()
	circleBlue.Draw()
	squareRed.Draw()
	squareBlue.Draw()

}
