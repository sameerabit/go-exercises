package main

import "fmt"

func main() {
	// sli := make([]int, 3)
	// for {
	// 	var newNum string
	// 	fmt.Print("Enter an integer: ")
	// 	b, _ := fmt.Scan(&newNum)
	// 	if b == 1 {
	// 		x, error := strconv.Atoi(newNum)
	// 		if error != nil {
	// 			if newNum == "X" {
	// 				break
	// 			}
	// 			continue
	// 		}
	// 		sli = append(sli, x)
	// 		sort.IntSlice(sli).Sort()
	// 		fmt.Println(sli)
	// 	}
	// }
	x := [...]int{1, 2, 3, 4, 5, 6}
	y := x[0:2]
	z := x[2:4]
	fmt.Println(len(y), len(z), cap(y), cap(z))
}
