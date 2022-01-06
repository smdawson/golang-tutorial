package main

import "fmt"

func main() {
	// Example of a Explicit Variable Definition
	var number uint16 = 260
	// Example of an Implicit Variable Definition
	var number1 = 260
	number = number + 5
	number3 := 3
	var bl bool
	fmt.Println("Hello World!", number, number1)
	fmt.Printf("%T", number3)
	fmt.Printf("%T", bl)
}
