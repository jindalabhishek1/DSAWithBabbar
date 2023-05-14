package main

import "fmt"

func main() {

	// for with just condition, we don't need empty placeholders like c,C++ if we want to escape the initaization and updation
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// forever for loop, we have to return or break from this kind of loops
	for {
		fmt.Println("loop")
		break
	}
	// infinite loop
	// for {
	// 	if i >= 0 {
	// 		fmt.Println(i)
	// 		i--
	// 	}
	// }

	// how to skip one iteration, use continue
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for p := 5; p >= 0 && p < 10; p++ {
		fmt.Printf("P: %v\n", p)
	}
	var a int
	fmt.Println("Enter something: ")
	fmt.Println(fmt.Scan(&a))
	fmt.Println(a)
}
