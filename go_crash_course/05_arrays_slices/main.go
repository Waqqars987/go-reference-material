package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// Assing Values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and Assign
	fruitArr := [2]string{"Apple", "Mango"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Slices
	fruitSlice := []string{"Apple", "Mango", "Orange", "Grape"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
