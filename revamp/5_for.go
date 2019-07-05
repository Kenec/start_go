package main

import "fmt"

// for is Go's only looping construct

func main() {
	i := 1

	// most basic type with a single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classic initial/condition/after for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a consition will loop repeatedly until you break out of the loop or return from the enclosing fn
	for {
		fmt.Println("loop")
		break
	}

	// continue to the next iteration in for loop
	for n:= 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
