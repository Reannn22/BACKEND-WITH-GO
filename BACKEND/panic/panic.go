package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("kata tidak boleh kosong")
	}
	return true, nil
}

func main() {
	var name string
	fmt.Print("Type your name: ")
	fmt.Scanln(&name)
	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		defer fmt.Println("tetap terbaca")
		panic(err.Error())
		fmt.Println("end")
	}
}
