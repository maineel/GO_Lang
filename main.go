package main

import "fmt" // fmt is a package that contains functions for printing formatted output and scanning input

func main() {
	var name string = "Neel Sheth"
	fmt.Println(name)

	name = "Neel"
	fmt.Println(name)

	nametwo := "Sheth" // another way of declaring a variable (shorthand method). This method cannot be used outside a function
	fmt.Println(nametwo)

	ageOne := 21
	ageTwo := 22
	fmt.Println(ageOne, ageTwo)

	// bots and memory
	var numOne int8 = 127  // int8 is a type of integer that can store values from -128 to 127
	var numTwo uint8 = 255 // uint8 is a type of integer that can store values from 0 to 255
	fmt.Println(numOne, numTwo)

}
