package main

// string converter
import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// changes last name
func (p *Person) changeLastName(newName string) {
	if p.age <= 24 {
		return
	} else {
		p.lastName = newName
	}
}

func main() {
	// Init person using struct

	person1 := Person{firstName: "Sarah", lastName: "Joy", city: "Boston", gender: "genderfluid", age: 24}

	//Alternative
	person2 := Person{"Kara", "Gale", "Boston", "f", 20}

	fmt.Println(person1)
	person1.age++
	fmt.Println(person1.firstName)
	fmt.Println(person1)

	// calling the method
	person1.hasBirthday()
	person1.changeLastName("Lou")
	person2.changeLastName("Kou")

	fmt.Println(person2.greet())
}
