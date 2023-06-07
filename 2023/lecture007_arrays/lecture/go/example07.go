// union of two arrays

package main

import "fmt"

func main() {
	var firstSize, secondSize int

	fmt.Printf("Size of first array: ")
	fmt.Scan(&firstSize)

	firstArr := make([]int, firstSize)

	for i := 0; i < firstSize; i++ {
		fmt.Scan(&firstArr[i])
	}

	fmt.Printf("Size of second array: ")
	fmt.Scan(&secondSize)

	secondArr := make([]int, secondSize)

	for i := 0; i < secondSize; i++ {
		fmt.Scan(&secondArr[i])
	}

}
