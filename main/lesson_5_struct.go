package main

import "fmt"


// A struct is a collection of fields.

type Cartesian struct {
	x int
	y int
	z int
}

func main() {
	v := Cartesian{2,4,2}
	fmt.Println(v)

	// struct values are accessible through dot
	fmt.Println(v.x)

}