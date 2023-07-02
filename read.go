package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filename string
	fmt.Print("Enter the filename: ")
	fmt.Scan(&filename)

	type name struct {
		fname string
		lname string
	}

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	fScanner := bufio.NewScanner(readFile)
	fScanner.Split(bufio.ScanLines)

	sli := make([]name, 0, 100)

	for fScanner.Scan() {
		n := name{}
		flname := strings.Split(fScanner.Text(), " ")
		n.fname = TrimToSize(flname[0], 20)
		n.lname = TrimToSize(flname[1], 20)

		sli = append(sli, n)
	}
	readFile.Close()

	for _, sliname := range sli {
		fmt.Printf("First Name: %s and Last Name: %s \n", sliname.fname, sliname.lname)
	}
}

func TrimToSize(str string, size int) string {
	str = strings.TrimSpace(str)
	if len(str) > size {
		str = str[:size]
	}
	return str
}
