// solid square

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
