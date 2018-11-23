package main

import "fmt"

// A defer statement defers the execution of a function until the surrounding function returns
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
	stacking_defer()
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.
func stacking_defer() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}