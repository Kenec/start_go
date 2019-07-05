package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		// simulate load
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSqure(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := longRequest(), longRequest()

	fmt.Println(sumSqure(<-a, <-b))
}

