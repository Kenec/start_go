package main

import "fmt"

// Go allows you to pass references to values and records within your program

func zeroval(ival int) {
	ival = 0
}

// zeroptr has an *int parameter, meaning that it takes an int pointer
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	j := &i
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	fmt.Println("memory address: ", &i)

	fmt.Println("copy of i: ", *j)
	fmt.Println("original i: ", i)

	*j = 12

	fmt.Println("modified through *j of i: ", *j)
	fmt.Println("modified throgh *j i: ", i)
}

