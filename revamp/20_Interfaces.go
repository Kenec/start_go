package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures.

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 3}

	// The circle and rectangle struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)
}


