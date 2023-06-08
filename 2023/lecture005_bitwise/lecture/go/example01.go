// doc: https://go.dev/ref/spec#Operators
package main

import "fmt"

func main() {

	// bitwise operators can work with integers only
	a := 10
	b := -1
	fmt.Printf("a: %v\tbits: %b\n", a, a)
	fmt.Printf("b: %v\tbits: %b\n", b, b)
	fmt.Printf("a | b: %v\tbits: %b\n", a|b, a|b)
	fmt.Printf("a & b: %v\tbits: %b\n", a&b, a&b)
	fmt.Printf("a ^ b: %v\tbits: %b\n", a^b, a^b)
	fmt.Printf("a &^ b: %v\tbits: %b\n", a&^b, a&^b)
	fmt.Printf("a << 1: %v\tbits: %b\n", a<<1, a<<1)
	fmt.Printf("b >> 1: %v\tbits: %b\n", b>>1, b>>1)
	fmt.Println("Byee!!!!")
}
