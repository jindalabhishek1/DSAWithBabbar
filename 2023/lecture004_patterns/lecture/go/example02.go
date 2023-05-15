// hollow diamond

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	// upper half
	// rows
	for i := 0; i < size; i++ {

		// spaces
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}

		// upper hollow pyramid
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
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

		// lower hollow pyramid
		for j := 0; j < size-i; j++ {
			if j == 0 || j == size-i-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}

		// newline
		fmt.Println()
	}
}
