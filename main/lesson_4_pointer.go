package main

import "fmt"

// Pointer holds memory address of a value

// This is known as "dereferencing" or "indirecting".

func main()  {
	i, j := 42, 1270

	p := &i // point to i
	fmt.Println(p) // read i though pointer
	*p = 21 // set i through pointer
	fmt.Println(i) // see the new value of i

	p = &j // point to j
	*p = *p / 37 // divide j through the pointer
	fmt.Println(j)

}
