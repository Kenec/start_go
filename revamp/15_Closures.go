package main

import "fmt"

// Anonymous functions are useful when you want to define a function inline without having to name it
func intSqe() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSqe()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSqe()
	fmt.Println(newInts())
	fmt.Println(newInts())
}