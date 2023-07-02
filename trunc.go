package main

import (
	"fmt"
)

func main() {
	var yourNum float64
	fmt.Printf("Enter a float number: ")
	fmt.Scan(&yourNum)
	fmt.Println(int64(yourNum))
	//fmt.Println(math.Trunc(yourNum))
}
