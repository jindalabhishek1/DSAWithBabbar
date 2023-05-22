// simple function

package main

import "fmt"

func printName() {
	var num int

	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	for i := 0; i < num; i++ {
		fmt.Println("Johny")
	}
}

func main() {
	printName()
}
