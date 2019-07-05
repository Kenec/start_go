package main

import "fmt"

// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

// struct are mutable
type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Kene", 29})
	fmt.Println(person{name: "Nzube", age: 27})
	fmt.Println(person{name: "Cat"})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 52
	fmt.Println(sp.age)
}

