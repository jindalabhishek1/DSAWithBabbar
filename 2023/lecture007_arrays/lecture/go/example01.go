// simple array

package main

import (
	"fmt"
	"reflect"
)

func main() {

	var arr [5]string

	// default value of array is 0 for int and nill(nothing) for string
	// array has fixed size
	fmt.Println("Default value of array in go:", arr)

	// can be declared like this
	// arrays length is part of it's type, so arrays cannot be resized
	primes := [5]int{2, 3, 5, 7, 11}
	primes3 := [...]int{2, 3, 4, 2}
	fmt.Println("declared and initialized at same time array:", primes)
	fmt.Println("Desclared and initialized array with compiler automatically calculating the array element:", primes3)
	fmt.Println("Type of above array:", reflect.TypeOf(primes3))

	// var primes2 [5]int
	// can't do the following
	// primes2 := {2, 3, 5, 7, 11}
	// fmt.Println(primes2)

	/*
		Slice is a more convinient way to interact with sequence than arrays

		Slice is define by the type only with no fixed size defined

		value of un initialized slice is nill and size 0

		slice is dynamically sized, view into elements of array
	*/
	slice1 := []int{1, 2, 3}

	fmt.Println("Basic Slice:", slice1)
	fmt.Println("lenth of above slice:", len(slice1))
	fmt.Println("Capacity of above slice:", cap(slice1))

	var slice2 []int
	fmt.Println("Default slice:", slice2, "Size:", len(slice2))

	slice3 := make([]int, 3)
	fmt.Println("Un-initialized slice with a non zero size:", slice3)

	slice4 := primes[1:4]
	fmt.Println("slice from array, from index including 1 till 4 (excluding):", slice4)

	slice1 = append(slice1, slice4...)
	fmt.Println("slice1 = append(slice1, slice4...):", slice1)
	// two D slice
	// the length of inner slices ca vary, unlike with multidimentional array
	var twoD [][]int
	twoD = make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLength := i + 1

		twoD[i] = make([]int, innerLength)

		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

}
