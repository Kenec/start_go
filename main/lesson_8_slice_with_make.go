package main

import (
	"fmt"
	"strings"
)

// A slice does not store any data, it just describes a section of an underlying array.
//Changing the elements of a slice modifies the corresponding elements of its underlying array.
func main() {
	slice_make()
	slice_slice()
	slice_append()
}

// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays
func slice_make() {
	a := make([]int, 5)
	fmt.Println("a: ", a)

	b := make([]int, 0, 6)
	fmt.Println("b: ", b)
}

// Slices can contain any type, including other slices
func slice_slice() {

	// createa tic tac toe board
	board := [][]string{
		[]string {"-", "-","-"},
		[]string {"-", "-", "-"},
		[]string {"-", "-","-"},
		[]string {"-", "-", "-"},
	}

	board[0][0] = "X"
	board[2][2] = "0"
	board[1][0] = "X"
	board[0][1] = "0"

	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}
}

// appending to a slice
// func append(s []T, vs ...T) []T
// The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

func slice_append() {
	var s []int
	fmt.Println(s, "caps: ", cap(s))

	// append works on nil slices.
	s = append(s, 0)
	fmt.Println(s, "caps: ", cap(s))

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	fmt.Println(s, "caps: ", cap(s))

}