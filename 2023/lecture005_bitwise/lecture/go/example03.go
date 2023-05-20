// swtich case

package main

import "fmt"

func main() {
	var num int
	num = 10
	temp := 5

	switch num - 5 {

	// case values need not to be a value and interger in go
	case temp + 3:
		fmt.Println("first case")
		// in go switch case, we don't need break
	case temp + 5:
		fmt.Println("Second case")
	case temp + 5:
		fmt.Println("Same second case")
	default:
		fmt.Println("Default case")
	}

	// switch with no condition is same as switch true
	switch {
	case temp > num:
		fmt.Println("temp greater than num")
	case temp < 4:
		fmt.Println("temp less than 4")
	case temp == num:
		fmt.Println("temp equeal to num")
	default:
		fmt.Println("Default case")
	}

}
