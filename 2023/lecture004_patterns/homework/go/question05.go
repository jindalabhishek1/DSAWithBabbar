// numeric hollow inverted half pyramid
/*

1 2 3 4 5
2     5
3   5
4 5
5

*/

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == 0 || j == 0 || j == size-i-1 {
				fmt.Printf("%v ", j+i+1)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
