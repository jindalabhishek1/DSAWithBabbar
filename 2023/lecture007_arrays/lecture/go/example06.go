// union/unique element (every element occurs twice in array except one element, that is union element)

package main

import "fmt"

func main() {

	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println(arr)
	// fmt.Printf("Size of arr: %v\n", len(arr))
	// fmt.Printf("Capacity of arr: %v\n", cap(arr))

	// arr = append(arr, 4, 5, 5, 6, 3, 9, 0)
	// fmt.Println(arr)
	// fmt.Printf("Size of arr: %v\n", len(arr))
	// fmt.Printf("Capacity of arr: %v\n", cap(arr))
	// arr := []int{1, 2, 4, 2, 1, 3, 6, 5, 5, 6, 4}
	// fmt.Println("Hello")

	answer := uniqueElement(&arr, len(arr))
	fmt.Printf("Unique element is: %v\n", answer)

}

// this solution only work if there is only one unique element
func uniqueElement(arr *[]int, size int) int {
	var ans int = 0
	for i := 0; i < size; i++ {

		fmt.Printf("Bits of ans: %b\n", ans)
		fmt.Printf("Bits of (*arr)[i]: %b\n", (*arr)[i])
		ans = ans ^ (*arr)[i]
		fmt.Printf("Bits of ans ^ (*arr)[i]: %b\n", ans)
	}
	return ans
}
