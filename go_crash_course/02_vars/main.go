package main

import "fmt"

// Global Variable
// var name = "Waqqar"

func main() {
	// Using var
	// var name = "Waqqar"
	// Shorthand Syntax, can only be used inside the function body
	// name := "Waqqar"
	// email := "waqqar@gmail.com"

	// assigning two varaibles at once
	name, email := "Waqqar", "waqqar@gmail.com"

	var age int32 = 37
	var size float32 = 2.3
	// size := 1.3
	const isCool = true
	// isCool = false

	fmt.Println(name, email, age, isCool)
	fmt.Printf("%T\n", size)
}
