// hollow inverted half pyramid
/*

* * * * *
*     *
*   *
* *
*

 */

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 || j == 0 || j == size-i-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
