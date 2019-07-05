package main

import "fmt"

// Go has built-in support for multiple return values
func vals() (int, int, int) {
	return 3, 7, 4
}

func main() {
	a, b, _ := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, _, d := vals()
	fmt.Println(d)
}
