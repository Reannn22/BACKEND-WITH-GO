package main

import (
	"fmt"
	"strings"
	"errors"
)

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("kata tidak boleh kosong")
	}
	return true, nil
}


func main() {
	defer catch()
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
	fmt.Println("halo", name)
	} else {
	panic(err.Error())
	fmt.Println("end")
	}
	}