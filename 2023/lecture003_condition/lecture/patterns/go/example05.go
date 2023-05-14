// inverted half pyramid

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	/*

	*****
	****
	***
	**
	*

	 */
	for i := 0; i < size; i++ {
		for j := size; j > i; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}

	/*

		*****
		 ****
		  ***
		   **
		    *

	*/
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j >= i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
