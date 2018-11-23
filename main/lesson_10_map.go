package main

import "fmt"

func main(){
	map_test()
	map_literal_test()
	map_again()
}

// map maps keys to values
// To create an empty map, use the builtin make: make(map[key-type]val-type)
func map_test() {
	m := make(map[string]int)

	m["k1"] = 10
	m["k2"] = 12

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := m["k2"]
	fmt.Println("prs:", prs)
}

// map literal are like struct only that they require key
func map_literal_test() {
	type Vertex struct {
		Lat, Long float64
	}

	var m = map[string]Vertex {
		"Andela Lagos": { 34.44, 334.44},
		"Andela Nairobi": { 40.89, 21.89},
	}

	fmt.Println(m)
}

// Do it your self
func map_again() {
	nigeria_population := make(map[string]int)
	nigeria_population = map[string] int {
		"Lagos": 23421,
		"Ondo": 23450,
	}
	nigeria_population["Enugu"] = 2345
	Lagos_population, ok := nigeria_population["Lagos"]
	fmt.Println(nigeria_population)
	delete(nigeria_population, "Lagos")
	fmt.Println(nigeria_population)
	fmt.Println(Lagos_population, ok)
}