// numeric full pyramid

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	/*

		     1
			121
		   12321
		  1234321
		 123454321
		12345654321

	*/

	// rows
	for i := 0; i < size; i++ {
		// print spaces
		for j := 0; j < size-i-1; j++ {
			fmt.Print(" ")
		}

		// print left half pyramid
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v", j+1)
		}

		// Print right half pyramid
		for j := 0; j < i; j++ {
			fmt.Printf("%v", i-j)
		}

		// new line
		fmt.Println()
	}

	/*

		     1
			232
		   34543
		  4567654
		 567898765
		678910119876

	*/
	// how to handle after 9, see that last row
}
