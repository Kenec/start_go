package main

import "fmt"

// The length of a slice is the number of elements it contains.
//  The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s)

func main() {
	// A slice literal is like an array literal without the length
	// Example
	// This is an array literal:
	//
	// [3]bool{true, true, false}
	// And this creates the same array as above, then builds a slice that references it:
	//
	// []bool{true, true, false}

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{1, true},
	}

	sli := q[2:]

	fmt.Println(s)
	fmt.Println(cap(sli))

	// The zero value of a slice is nil.

	var M []int
	fmt.Println(M)
	if M == nil {
		fmt.Println("Nil")
	}
}
