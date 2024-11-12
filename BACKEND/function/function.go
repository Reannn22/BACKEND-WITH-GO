package main

import (
	"fmt"
	"strings"
)

func main() {
	var names = []string{"Fauzi", "Edspert"}
	fungsi("halo", names)
	fungsi("lain", names)
}

func fungsi(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
