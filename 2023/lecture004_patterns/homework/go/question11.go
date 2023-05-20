// floid's triangle
/*

1
2 3
4 5 6
7 8 9 10
11 12 13 14 15
16 17 18 19 20 21
22 23 24 25 26 27 28

*/

package main

import "fmt"

func main() {
	var size int

	fmt.Print("Size: ")
	fmt.Scan(&size)

	temp := 0
	for i := 0; i < size; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%v ", temp+1)
			temp++
		}
		fmt.Println()
	}
}
