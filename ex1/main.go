package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter Your First Name: ")
	var first string
	fmt.Scanln(&first)

	fmt.Println("Enter Your Middle Name: ")
	var middle string
	fmt.Scanln(&middle)

	fmt.Println("Enter Your Last Name: ")
	var last string
	fmt.Scanln(&last)

	fmt.Println("Enter Your Country Code ")
	var country string
	fmt.Scanln(&country)

	if country != "VN" {
		tmp := first
		first = last
		last = tmp
	}

	if middle != "" {
		fmt.Println(first + " " + middle + " " + last)
	} else {
		fmt.Println(first + " " + last)
	}

}
