package design_patterns

import "fmt"

type Prototype interface {
	SetAge(age int)
	SetName(name string)
	Clone() Prototype
}

type ConcretePrototype struct {
	Name string
	Age  int
}

func (cP *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{
		Name: cP.Name,
		Age:  cP.Age,
	}
}

func (cP *ConcretePrototype) SetName(name string) {
	cP.Name = name
}

func (cP *ConcretePrototype) SetAge(age int) {
	cP.Age = age
}

func NewPrototype(name string, age int) Prototype {
	return &ConcretePrototype{
		Name: name,
		Age:  age,
	}
}

func PrintPrototype() {
	proto := NewPrototype("tesla proto", 80)
	clone1 := proto.Clone()
	clone2 := proto.Clone()

	clone1.SetAge(10)
	clone1.SetName("tesla No.2")
	clone2.SetName("tesla No.3")

	fmt.Println(proto, clone1, clone2)
}
