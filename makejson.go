package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Println("Please enter your name :")
	fmt.Scan(&name)

	fmt.Println("Please enter your address :")
	fmt.Scan(&address)

	student := make(map[string]string)
	student["name"] = name
	student["address"] = address

	studenJson, err := json.Marshal(student)
	if err == nil {
		fmt.Println(string(studenJson))
	}

}
