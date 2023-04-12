package study_go

import (
	"fmt"
)

type Passenger struct {
	Name string
}

type Bus struct {
	PLimit, PNum   int
	PassengerQueue []Passenger
}

func (b *Bus) pop() {
	if b.PNum == 0 {
		return
	}
	b.PassengerQueue = b.PassengerQueue[1:]
	b.PNum--
}

func (b *Bus) push(newP *Passenger) {
	if b.PNum >= b.PLimit {
		return
	}
	b.PassengerQueue = append(b.PassengerQueue, *newP)
	b.PNum++
}

func (b *Bus) showPassengers() {
	fmt.Println(b.PassengerQueue)
}

func (b *Bus) showSeats() {
	fmt.Println(b.PLimit-b.PNum, "seats remains")
}

func PrintStruct() {
	passenger1 := Passenger{"Jack"}
	passenger2 := Passenger{Name: "Dick"}
	bus := &Bus{PassengerQueue: []Passenger{{"Alice"}}, PLimit: 100, PNum: 1}

	bus.showPassengers()
	bus.pop()
	bus.push(&passenger1)
	bus.showPassengers()
	bus.push(&passenger2)
	bus.showPassengers()
	bus.showSeats()
	bus.pop()
	bus.showPassengers()
}
