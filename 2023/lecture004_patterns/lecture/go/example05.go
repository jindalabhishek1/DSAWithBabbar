/*

A
A B A
A B C B A
A B C D C B A
A B C D E D C B A

*/

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	// option 1
	// rows
	for i := 0; i < size; i++ {
		k := 0
		// columns
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("%c ", 65+k)

			// till column is less than row, keep increasing the alphabet
			if j < i {
				k++
			} else {
				k--
			}
		}
		fmt.Println()
	}

	// option 2
	// rows
	for i := 0; i < size; i++ {

		// first part
		for j := 0; j < i+1; j++ {
			fmt.Printf("%c ", 65+j)
		}

		// second part
		for j := 0; j < i; j++ {
			fmt.Printf("%c ", 65+(i-j-1))
		}

		// newline
		fmt.Println()
	}
}
