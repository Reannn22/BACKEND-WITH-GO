package main

import "fmt"

func main() {
	var firstName string = "my"
	const lastName = "skill"

	firstName = "aku"
	
	//var bilanganBulat uint8 = 20

	//var bilanganDesimal = 2.0

	//var varBool = true

	var numberA int = 4
	var numberB	*int = &numberA
	*numberB = 8

	fmt.Println("halo ", firstName, lastName, numberA, "!\n")
	
}
