// half pyramid

/*

*
**
***
****

 */

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	/*

		    *
		   **
		  ***
		 ****
		*****

	*/

	// first approach
	for i := size; i > 0; i-- {
		for j := 0; j < size; j++ {
			if j >= i-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	// second approach
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j >= size-i-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
