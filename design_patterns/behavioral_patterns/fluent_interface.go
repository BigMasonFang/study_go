package behavioral_patterns

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) *Person {
	p.name = name
	return p
}

func (p *Person) SetAge(age int) *Person {
	p.age = age
	return p
}

func (p *Person) Describe() {
	fmt.Printf("Name: %s, Age: %d\n", p.name, p.age)
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func PrintFluentInterface() {
	p := NewPerson("pete", 15)
	p.SetName("peter").SetAge(20).Describe()
}
