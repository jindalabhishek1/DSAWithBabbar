// hollow full pyramid

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	/*

			1
		   1 2
		  1   3
		 1     4
		1 2 3 4 5

	*/

	// rows
	for i := 0; i < size; i++ {

		// last row
		if i == size-1 {
			for j := 0; j < size; j++ {
				fmt.Printf("%v ", j+1)
			}
		} else {

			// left empty spaces
			for j := 0; j < size-i-1; j++ {
				fmt.Print(" ")
			}

			// pyramid
			for j := 0; j < i+1; j++ {
				if j == 0 || j >= i {
					fmt.Printf("%v ", j+1)
				} else {
					fmt.Print("  ")
				}
			}
		}
		// newline
		fmt.Println()
	}
}
