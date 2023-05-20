/*

*
* 1 *
* 1 2 1 *
* 1 2 3 2 1 *
* 1 2 1 *
* 1 *
*

 */

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size+1; i++ {
		k := 0
		for j := 0; j < 2*i+1; j++ {
			if j == 0 || j == 2*i {
				fmt.Print("* ")
			} else {
				if j > i {
					k--
				} else {
					k++
				}
				fmt.Printf("%v ", k)
			}
		}
		fmt.Println()
	}

	for i := 0; i < size; i++ {
		k := 0
		for j := 0; j < 2*size-2*i-1; j++ {
			if j == 0 || j == 2*size-2*i-2 {
				fmt.Print("* ")
			} else {
				if j >= size {
					k--
				} else {
					k++
				}
				fmt.Printf("%v ", k)
			}
		}
		fmt.Println()
	}
}
