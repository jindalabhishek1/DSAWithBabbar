// numeric hollow half pyramid
/*

1
1 2
1   3
1     4
1 2 3 4 5

*/

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i || i == size-1 {
				fmt.Printf("%v ", j+1)
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
