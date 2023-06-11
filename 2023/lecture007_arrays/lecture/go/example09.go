// pair whose sum is equal to a sum (brute force)

package main

import "fmt"

func main() {
	var sum, arraySize int
	var array []int

	fmt.Printf("Size of Array: ")
	fmt.Scan(&arraySize)

	array = make([]int, arraySize)

	for i := 0; i < arraySize; i++ {
		fmt.Scan(&array[i])
	}

	fmt.Printf("Sum to find: ")
	fmt.Scan(&sum)

	fmt.Println(array)

	findPair(&array, arraySize, sum)

}

func findPair(array *[]int, size, sum int) {
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			firstNum := (*array)[i]
			secondNum := (*array)[j]
			if firstNum+secondNum == sum {
				fmt.Println(firstNum, secondNum)
				return
			}
		}
	}
	fmt.Printf("No pair found for which sum is %v.\n", sum)
}
