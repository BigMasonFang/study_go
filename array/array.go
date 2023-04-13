package array

import (
	"fmt"
)

type Person struct {
	Name  string
	Money int
}

func printPersonArray(persons *[5]Person) {
	for _, p := range *persons {
		fmt.Printf("Name: %v Money: %v\n", p.Name, p.Money)
	}
}

func calcPersonArrayMoney(persons *[5]Person) int {
	moneySum := 0
	for i := range *persons {
		moneySum += persons[i].Money
	}
	return moneySum
}

func PrintArray() {
	nameArray := [...]string{"a", "b", "c", "d"}
	moneyArray := [4]int{100, 200, 300, 400}
	personArray := [5]Person{}
	for i := 0; i < len(nameArray); i++ {
		personArray[i] = Person{nameArray[i], moneyArray[i]}
	}
	fmt.Println(personArray)

	personArray[len(personArray)-1] = Person{"e", 50}
	fmt.Println(personArray)

	printPersonArray(&personArray)
	fmt.Printf("Money sum is %v\n", calcPersonArrayMoney(&personArray))

	personArray[3].Money = 300
	personArray[4].Money = 100
	printPersonArray(&personArray)
}
