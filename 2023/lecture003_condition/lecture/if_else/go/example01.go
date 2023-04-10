package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scanln(&num)

	if num > 20 {
		fmt.Println("Greater that 20")
	} else if num < 20 {
		fmt.Println("Less than 20")
	} else {
		fmt.Println("Equal to 20")
	}
}
