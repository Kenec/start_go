package main

import "fmt"

// When ranging over a slice, two values are returned for each iteration.
// The first is the index, and the second is a copy of the element at that index.
func main() {
	range_slice()
	fmt.Println(1 >> uint(1))
	range_one()
}

func range_slice() {
	pow := []int{1,2,3,4,5,6,7}

	for i, v := range pow {
		fmt.Println(i, v)
	}
}

// You can skip the index or value by assigning to _.
func range_one() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
