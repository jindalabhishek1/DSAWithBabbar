// sort 0's and 1's

package main

import "fmt"

func main() {
	var arraySize int

	var array []int

	fmt.Printf("Size of array: ")
	fmt.Scan(&arraySize)

	array = make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		fmt.Scan(&array[i])
		fmt.Printf("Type of array[i]: %T\tType of 0: %T\tType of 1: %T\n", array[i], 0, 1)
		// have a look at the condition
		if array[i] != 0 || array[i] != 1 {
			i--
			fmt.Println("Valid values are 0 or 1, please enter again")
		}
	}
	fmt.Println(array)
}
