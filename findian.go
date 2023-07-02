package main

import (
	"fmt"
	"strings"
)

func main() {
	var yourWord string
	fmt.Printf("Enter any String: ")
	fmt.Scan(&yourWord)

	if yourWord[0:1] == "i" && yourWord[len(yourWord)-1:] == "n" && strings.Contains(strings.ToLower(yourWord), "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
