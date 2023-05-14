// numeric half pyramid

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	/*

		1
		12
		123
		1234
		12345

	*/
	for i := 0; i < size; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v ", j+1)
		}
		fmt.Println()
	}

	/*

		    1
		   12
		  123
		 1234
		12345

	*/
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if j >= size-i-1 {
				fmt.Printf("%v ", i-size+j+2)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
