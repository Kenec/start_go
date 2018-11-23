package main

import "fmt"


// A struct is a collection of fields.

type Cartesian struct {
	x int
	y int
	z int
}

type AndelaFellows struct {
	name string
	age int
	stack []string
}

var (
	v1 = Cartesian{3, 4, 1}
	v2 = Cartesian{ x:1, y:7 }
	v3 = Cartesian{}
	lo = &Cartesian{3, 9, 10}

)

func main() {
	v := Cartesian{2,4,2}
	fmt.Println(v)

	// struct values are accessible through dot
	fmt.Println(v.x)

	// struct pointer
	// instead of *p.X use p.X

	p := &v
	fmt.Println(p.y)

	fmt.Println(v1, lo, v2, v3)
	andelans()

}

func andelans() { 
	Andela_36 := AndelaFellows {
		name: "Nnamani Kenechukwu",
		age: 28,
		stack: []string {
			"Javascript",
			"DevOps",
			"Golang",
		},
	}
	fmt.Println(Andela_36)
}