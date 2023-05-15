// solid diamond

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	// upper half
	//rows
	for i := 0; i < size; i++ {

		// space
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}

		// upper pyramid
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}

		// newline
		fmt.Println()
	}

	// lower half
	// rows
	for i := 0; i < size; i++ {

		// spaces
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		// lower pyramid
		for j := 0; j < size-i; j++ {
			fmt.Print("* ")
		}

		// newline
		fmt.Println()
	}
}
