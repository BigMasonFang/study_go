/*
The Singleton pattern is a creational design pattern that ensures that only one instance of a class is created during the program's execution.
The purpose of this pattern is to provide a global point of access to the object,
which can be accessed by all other objects in the program.The Singleton pattern involves a single class,
which is responsible for creating and managing its own instance. The class ensures that only one instance exists by providing a specific method,
usually called getInstance(), which returns the instance of the class. This method checks if an instance of the class has already been created,
and if not, it creates one, stores it in the class, and returns it.

threads pool, db connection pool, global log, config
*/
package creational_patterns

import (
	"fmt"
	"sync"
)

type Singleton struct {
	count int
}

func (s *Singleton) AddOne() int {
	s.count++
	return s.count
}

var instance *Singleton
var once sync.Once

// global
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{0}
		// config.init(filename)
	})
	instance.AddOne()
	return instance
}

func PrintSingleton() {
	s1 := GetInstance()
	s2 := GetInstance()
	// %p is return the pointer
	// it is not right to judge by following statement
	// fmt.Printf("instance's address is %v count of s1 is %v\n", &s1, s1.count)
	// fmt.Printf("instance's address is %v count of s2 is %v\n", &s2, s1.count)
	// because it is two different pointers
	fmt.Println("is pointer &s1 and &s2 same? ", &s1 == &s2)

	fmt.Printf("instance's address is %p count of s1 is %v\n", s1, s1.count)
	fmt.Printf("instance's address is %p count of s2 is %v\n", s2, s2.count)
	fmt.Println("Is instance s1 and s2 the same", s1 == s2)
}
