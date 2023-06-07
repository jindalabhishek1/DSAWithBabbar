// linear search

package main

import "fmt"

func main() {
	arr := []int{3, 3, 6, 1, 0, -6}

	var res, num int

	fmt.Print("Enter the number to find: ")
	fmt.Scan(&num)

	res = linear_search(arr, len(arr), num)
	fmt.Println(res)

}

func linear_search(arr []int, size, num int) int {
	for i := 0; i < size; i++ {
		if arr[i] == num {
			return i
		}
	}
	return -1
}
