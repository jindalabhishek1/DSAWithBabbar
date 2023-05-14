// solid rectangle pattern

package main

import "fmt"

func main() {
	var rows, columns int

	fmt.Print("Enter rows: ")
	fmt.Scan(&rows)

	fmt.Print("Enter columns: ")
	fmt.Scan(&columns)

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("Hint: Play with different kinds of input")
}
