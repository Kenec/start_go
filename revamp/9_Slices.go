package main

import "fmt"

func main(){

	// create an empty slice of len 3
	s := make([]string, 3)
	fmt.Println("emp:",s)

	// populate the empty slice
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:",s[2])

	// len return the length of the slice
	fmt.Println("len:", len(s))

	// append return slices containing new values
	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("apd", s)

	// slices can be copied.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices support "slice" operator. Slices from position 2 to 4. Take note that it excluded 5
	l := s[2:5]
	fmt.Println("sl1:", l)

	// slice up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// slice up from (an including) s[2]
	l = s[2:]
	fmt.Println("sl3", l)

	// slice can be declared in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:",t)

	// slices can be composed into multi-dimensional data structure.
	// The length of the inner slices can vary unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
