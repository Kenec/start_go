package main

import "fmt"

// The type [n]T is an array of n values of type T

func main() {
	// An array's length is part of its type, so arrays cannot be resized.
	var a[3]string

	a[0] = "Hello "
	a[1] = "world"

	prime := [4]int{2,3,5,7}

	fmt.Println(a[0] + a[1])
	fmt.Println(prime)

	// A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	var s []int = prime[1:4]
	fmt.Println(s)

	//A slice does not store any data, it just describes a section of an underlying array.
	//Changing the elements of a slice modifies the corresponding elements of its underlying array.



}
