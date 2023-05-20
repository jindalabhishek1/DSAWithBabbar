/*

1
2*3
4*5*6
7*8*9*10
7*8*9*10
4*5*6
2*3
1

*/

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	for i := 0; i < size; i++ {
		for j := 0; j < 2*i+1; j++ {
			if j%2 == 0 {
				fmt.Printf("%v", i+j+1)
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	for i := 0; i < size; i++ {
		for j := 0; j < 2*size-2*i-1; j++ {
			if j%2 == 0 {
				fmt.Printf("%v", size+j-i)
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
