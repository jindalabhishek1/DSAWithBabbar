// variable scoping

package main

import "fmt"

func main() {
	var temp int

	temp = 1
	fmt.Println("Value of temp in main func scope:", temp)
	if true {
		var temp float32

		temp = 4.5

		fmt.Println("Value of temp in if scope:", temp)
	}
}
