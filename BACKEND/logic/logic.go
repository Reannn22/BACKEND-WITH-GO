package main

import "fmt"

func main()  {
	var value = true && true
	var value1 = true || false
	var value2 = !false
	fmt.Println(value, value1, value2)	
}