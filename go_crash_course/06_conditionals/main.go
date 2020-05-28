package main

import "fmt"

func main() {
	x := 10
	y := 10

	// If Else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Else if
	color := "red"
	if color == "red" {
		fmt.Println("Color is RED")
	} else if color == "blue" {
		fmt.Println("Color is BLUE")

	} else {
		fmt.Println("Color is not RED or BLUE")
	}

	// Switch
	switch color {
	case "red":
		fmt.Println("Color is RED")
	case "blue":
		fmt.Println("Color is BLUE")
	default:
		fmt.Println("Color is not RED or BLUE")
	}
}
