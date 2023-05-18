package creational

import "fmt"

// The builder pattern is a creational design pattern that provides a solution to the complex object creation process
// by separating the construction of the object from its representation.
// In Go, the builder pattern involves creating a separate builder struct that handles the construction process of the desired object.
// The builder struct contains specific setter methods that can be used to set required attributes for the object being constructed.
// The builder struct also contains a `Build()` method, which returns the constructed object.
type Car struct {
	make  string
	model string
	year  string
}

type CarBuilder interface {
	SetMake(make string)
	SetModel(model string)
	SetYear(yr string)
	Build() *Car
}

type ConcreteCarBuilder struct {
	car *Car
}

func (carB *ConcreteCarBuilder) SetMake(make string) {
	carB.car.make = make
}

func (carB *ConcreteCarBuilder) SetModel(model string) {
	carB.car.model = model
}

func (carB *ConcreteCarBuilder) SetYear(yr string) {
	carB.car.year = yr
}

func (carB *ConcreteCarBuilder) Build() *Car {
	return carB.car
}

type Director struct {
	builder CarBuilder
}

func (d *Director) SetBuilder(builder CarBuilder) {
	d.builder = builder
}

func (d *Director) Construct() {
	d.builder.SetModel("Y")
	d.builder.SetMake("MK2")
	d.builder.SetYear("2019")
}

func PrintBuilder() {
	ConcreteCarBuilder := &ConcreteCarBuilder{car: &Car{}}
	director := &Director{builder: ConcreteCarBuilder}
	director.Construct()
	car := ConcreteCarBuilder.Build()
	fmt.Println(car.make, car.model, car.year)
}
