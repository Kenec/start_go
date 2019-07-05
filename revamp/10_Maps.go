package main

import "fmt"

func main() {
	// Maps are Go's built-in associative data type - also called hashes or dicts in other languages
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get a value for key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// builtin len returns the number of key/value pairs when called on map.
	fmt.Println("len:", len(m))

	// built-in delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// the optional second return value when getting a value from a map indicates if the key was present in the map
	_, prs := m["k1"]
	fmt.Println("prs",prs,)

	// map can be declared and initialized in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:",n)
}