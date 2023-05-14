// inverted full pyramid

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < size-i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
