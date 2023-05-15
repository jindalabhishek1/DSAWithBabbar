// flipped solid diamond

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	// upper half
	// rows
	for i := 0; i < size; i++ {

		// left stars
		for j := 0; j < size-i; j++ {
			fmt.Print("*")
		}

		// upper pyramid of spaces
		for j := 0; j < i+1; j++ {
			fmt.Print("  ")
		}

		// left stars
		for j := 0; j < size-i; j++ {
			fmt.Print("*")
		}

		// newline
		fmt.Println()
	}

	// lower part
	// rows
	for i := 0; i < size; i++ {

		// left stars
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		// lower pyramid of spaces
		for j := 0; j < size-i; j++ {
			fmt.Print("  ")
		}

		// right stars
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		// newline
		fmt.Println()
	}
}
