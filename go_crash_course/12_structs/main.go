package main

import (
	"fmt"
	"strconv"
)

// Person : Define person struct
type Person struct {
	// firstName string
	// lastname  string
	// city      string
	// gender    string
	// age       int

	firstName, lastname, city, gender string
	age                               int
}

// greet method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastname + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	}
	p.lastname = spouseLastName

}

func main() {
	// Init person using struct
	// person1 := Person{firstName: "Samantha", lastname: "Smith", city: "Boston", gender: "F", age: 25}
	// Alternative
	person1 := Person{"Samantha", "Smith", "Boston", "F", 25}
	person2 := Person{"Bob", "Johnson", "New York", "M", 30}
	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Williams")
	person2.getMarried("Thompson")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
