package main

import "fmt"

func main()  {
	// range itertates over elements
	nums := []int {2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	// range on slice and array provides index and value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on maps iterates over the key/value pairs
	kvs := map[string]string{"a": "apple", "b": "bannana", "c": "cassave"}
	for k, v := range kvs {
		fmt.Printf("%s ==> %s\n", k, v)
	}

	// range can iterate just over the key
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range over string iterates over unicode code (ASCIII) point
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
