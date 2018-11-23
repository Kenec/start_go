package main

import (
	"math"
	"fmt"
)

// Go does not have classes
// A method is a function with a special receiver argument

func main() {
	v := Vertex{ 2, 3}
	fmt.Println(v.method())
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) method() float64 {
	return math.Sqrt(v.X * v.X + v.Y* v.Y)
}
