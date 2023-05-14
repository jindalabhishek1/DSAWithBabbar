// inverted numeric pattern

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	/*

		12345
		 1234
		  123
		   12
		    1

	*/
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j >= i {
				fmt.Printf("%v ", j+1-i)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	// option 2
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			temp := j - i + 1
			if temp > 0 {
				fmt.Printf("%v ", temp)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	/*

		12345
		1234
		123
		12
		1

	*/
	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			fmt.Printf("%v ", j+1)
		}
		fmt.Println()
	}
}
