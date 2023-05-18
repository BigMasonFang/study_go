// 1. decoupling, observer can change independently
// 2. broadcast
// 3. the fundation of event drive design, decouple event
// and event handling, bring flexibility and modulization

package behavioral

import "fmt"

type Observer interface {
	Update(string)
}

type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(msg string) {
	fmt.Printf("%s recevied : %s\n", co.name, msg)
}

type Subject interface {
	Register(Observer)
	Unregister(Observer)
	Notify(string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (cs *ConcreteSubject) Register(observer Observer) {
	cs.observers = append(cs.observers, observer)
}

func (cs *ConcreteSubject) Unregister(observer Observer) {
	for i, obs := range cs.observers {
		if obs == observer {
			cs.observers = append(cs.observers[:i], cs.observers[i+1:]...)
			break
		}
	}
}

func (cs *ConcreteSubject) Notify(msg string) {
	for _, obs := range cs.observers {
		obs.Update(msg)
	}
}

func PrintObserver() {
	obs1 := &ConcreteObserver{"ob1"}
	obs2 := &ConcreteObserver{"ob2"}

	subject := &ConcreteSubject{}
	subject.Register(obs1)
	subject.Register(obs2)

	subject.Notify("msg1")

	subject.Unregister(obs1)

	subject.Notify(("msg2"))
}
