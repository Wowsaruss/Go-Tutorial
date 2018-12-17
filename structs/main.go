package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	phone   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// contact := contactInfo{
	// 	email:   "russ@sibipro.com",
	// 	phone:   "6236933758",
	// 	zipCode: 85203,
	// }
	alex := person{
		lastName:  "Anderson",
		firstName: "Alex",
		contactInfo: contactInfo{
			email:   "russ@sibipro.com",
			phone:   "6236933758",
			zipCode: 85203,
		},
	}
	alex.print()

	// alexPointer := &alex
	// alexPointer.updateName("Russell", "Hayes")
	alex.updateName("Russell", "Hayes")
	alex.print()
}

func (pointerToPerson *person) updateName(firstName string, lastName string) {
	(*pointerToPerson).firstName = firstName
	(*pointerToPerson).lastName = lastName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v", p)
}
