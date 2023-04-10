// factorial using recursion

package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return num
	}
	return num * factorial(num-1)
}

func main() {
	var input, output int
	fmt.Print("Enter the number for which you want to find factorial: ")
	fmt.Scan(&input)

	output = factorial(input)
	fmt.Printf("Factorial of %v is %v\n", input, output)
}
