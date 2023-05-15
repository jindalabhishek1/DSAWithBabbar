// fancy pattern

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	/*

		1
		2*2
		3*3*3
		4*4*4*4
		4*4*4*4
		3*3*3
		2*2
		1

	*/

	// upper half
	// rows
	for i := 0; i < size; i++ {

		// columns
		for j := 0; j <= i; j++ {
			if j%2 == 0 {
				fmt.Printf("%v", i+1)
			} else {
				fmt.Print("*")
			}
		}

		// newline
		fmt.Println()
	}

	// lower half
	// rows
	for i := 0; i < size; i++ {

		// columns
		for j := 0; j < size-i; j++ {
			if j%2 == 0 {
				fmt.Printf("%v", size-i)
			} else {
				fmt.Print("*")
			}
		}

		// newline
		fmt.Println()
	}

}
