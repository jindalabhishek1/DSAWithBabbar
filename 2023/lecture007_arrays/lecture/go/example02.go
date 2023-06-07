// arrays

package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 5}
	fmt.Println("Before:", arr)
	for i := 0; i < 5; i++ {
		arr[i] = 1
	}
	fmt.Println("After:", arr)

	var arr1 [5]int
	fmt.Println(arr1)
}
