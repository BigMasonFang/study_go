package creational

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Product interface {
	GetName() string
}

type ProductA struct {
	price float64
}

type ProductB struct {
	price float64
}

func (p *ProductA) GetName() string {
	return "ProductA"
}

func (p *ProductB) GetName() string {
	return "ProductB"
}

// factory
type SimpleProductFactory struct{}

func (f *SimpleProductFactory) CreateProduct(t string, price float64) Product {
	switch t {
	case "A":
		return &ProductA{price}
	case "B":
		return &ProductB{price}
	default:
		return nil
	}
}

func PrintSimpleFactory() {
	simplePorductFactory := &SimpleProductFactory{}
	productA := simplePorductFactory.CreateProduct("A", 64.0)
	productB := simplePorductFactory.CreateProduct("B", 57.0)
	fmt.Println(productA)
	spew.Dump(productB)
}
