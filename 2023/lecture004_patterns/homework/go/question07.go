/*

********1********
*******2*2*******
******3*3*3******
*****4*4*4*4*****
****5*5*5*5*5****

 */

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i+3; j++ {
			fmt.Print("*")
		}

		for j := 0; j < 2*i+1; j++ {
			if j%2 == 0 {
				fmt.Printf("%v", i+1)
			} else {
				fmt.Print("*")
			}
		}

		for j := 0; j < size-i+3; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
