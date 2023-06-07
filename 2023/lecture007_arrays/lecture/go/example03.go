// passing array to a function

package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 2, 0}
	size := len(arr)
	fmt.Println("Before:")
	fmt.Println("Size:", size)
	fmt.Println("array:", arr)
	double(arr, size)
	fmt.Println("After:")
	fmt.Println("Size:", size)
	fmt.Println("array:", arr)
}

func double(a []int, size int) {
	for i := 0; i < size; i++ {
		a[i] = a[1] * 2
	}
	fmt.Println("Inside function:")
	fmt.Println("Size:", size)
	fmt.Println("array:", a)
}
