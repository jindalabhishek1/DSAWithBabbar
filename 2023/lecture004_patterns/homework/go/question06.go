// numeric palindrome equilateral pyramid
/*

        1
      1 2 1
    1 2 3 2 1
  1 2 3 4 3 2 1
1 2 3 4 5 4 3 2 1

*/

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			fmt.Print("  ")
		}

		for j := 0; j < i+1; j++ {
			fmt.Printf("%v ", j+1)
		}

		for j := 0; j < i; j++ {
			fmt.Printf("%v ", i-j)
		}
		fmt.Println()
	}
}
