// count 0's and 1's in array

package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 0, 0, 1}

	var zeros, ones int
	fmt.Println("Before:")
	fmt.Println("Zero's:", zeros, "One's:", ones)

	// passing array as slice to a function would give an error
	zeros, ones = count(arr, len(arr))

	fmt.Println("Zero's:", zeros, "One's:", ones)
}

func count(arr []int, size int) (int, int) {

	zeros, ones := 0, 0
	for i := 0; i < size; i++ {
		if arr[i] != 0 && arr[i] != 1 {
			return -1, -1
		}

		if arr[i] == 0 {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}
