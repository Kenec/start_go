package main

import "fmt"

type rect struct {
	width, height int
}

type developer struct {
	name, location string
}

func (p *developer) getName() string {
	return p.name
}

func (p *developer) getLocation() string {
	return p.location
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width:20, height:5}
	developer := developer{"Kene", "Lagos"}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	fmt.Println("Developer: ", developer.name)
	fmt.Println("Location: ", developer.location)
}
