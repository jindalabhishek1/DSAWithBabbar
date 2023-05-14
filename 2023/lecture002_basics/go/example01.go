package main

import "fmt"

func main() {

	// simple print
	fmt.Println("Hello")

	// string to int
	var num int
	num = 'c'
	fmt.Println(num)

	// type conversions
	fmt.Println(20.0 / 2)
	fmt.Println(float32(20.0) / 2)
	fmt.Println(10 / 3)
	fmt.Println(10.0 / 3)
	// https://golangbyexample.com/typed-untyped-constant-golang/
	// fmt.Println(int(10.06) / 2) // error
	fmt.Printf("Type of untyped constant: %T\n", 10.04)

	var somefloat float64
	somefloat = 10.04
	fmt.Printf("Type of somefloat: %T\n", somefloat)
	fmt.Printf("Type of int(somefloat): %T\n", int(somefloat))
	fmt.Println(int(somefloat) / 2)

	// operators
	fmt.Println(7 >> 1) // right shift
	fmt.Println(8 == 7) // equality
	fmt.Println((1 == 1) && (2 == 3) && (-1 == -1))
	fmt.Println((1 == 1) || (2 == 4) && (4 == 4))
}
