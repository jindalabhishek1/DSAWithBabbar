// solid half diamond

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

	for i := 0; i < size-1; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
