package main

import "fmt"

func greeting(name string) string {
	return "Hello " + name
}

// if both input params are of the same type, mentioning only once would suffice
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Waqqar"))
	fmt.Println(getSum(3, 4))
}
