package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	num1, num2 := 0,1
	fmt.Println(num1)
	fmt.Println(num2)
	return func() int {
		fib := num1 + num2
		num1 = num2
		num2 = fib
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
