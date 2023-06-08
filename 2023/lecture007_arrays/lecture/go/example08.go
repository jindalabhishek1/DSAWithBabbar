// intersection of two arrays

package main

import "fmt"

func main() {
	var firstSize, secondSize int

	fmt.Println("Size of first array: ")
	fmt.Scan(&firstSize)

	firstArr := make([]int, firstSize)

	for i := 0; i < firstSize; i++ {
		fmt.Scan(&firstArr[i])
	}

	fmt.Println("Size of second array: ")
	fmt.Scan(&secondSize)

	secondArr := make([]int, secondSize)

	for i := 0; i < secondSize; i++ {
		fmt.Scan(&secondArr[i])
	}

	intersectionElement(firstArr, firstSize, secondArr, secondSize)
}

func intersectionElement(array1 []int, size1 int, array2 []int, size2 int) {
	// for i, elementArray1 := range array1 {
	// 	fmt.Printf("index: %v\n", i)
	// 	for elementArray2 := range array2 {
	// 		if elementArray2 == elementArray1 {
	// 			fmt.Println(elementArray1)
	// 			elementArray2 = -1
	// 		}
	// 	}
	// }
	fmt.Println("Finding intersections ...")

	fmt.Printf("First array: %v\n", array1)
	fmt.Printf("Second array: %v\n", array2)
	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			if array1[i] == array2[j] {
				fmt.Println(array1[i])
				array2[j] = -1
			}
		}
	}
}
