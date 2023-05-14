// square pattern

package main

import "fmt"

func main() {

	var sides int

	fmt.Print("Sides: ")
	fmt.Scan(&sides)

	for i := 0; i < sides; i++ {
		for j := 0; j < sides; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
