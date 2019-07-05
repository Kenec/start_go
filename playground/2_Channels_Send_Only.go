package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTImeRequest(r chan<- int32 ) {
	// simulate heavy load
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main()  {
	rand.Seed(time.Now().UnixNano())
	r := make(chan int32, 2)
	go longTImeRequest(r)
	go longTImeRequest(r)

	fmt.Println(sumSquares(<-r, <-r))
}