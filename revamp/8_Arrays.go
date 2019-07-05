package main

import "fmt"

func main() {
	// initialize empty array of length 5
	var a [5]int
	fmt.Println("emp:", a)

	// set a value for an index
	a[4] = 200
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// length of an array
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// multi dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}