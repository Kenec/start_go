package main

import (
	"fmt"
	"time"
)

// channels are the pipes that connect concurrent goroutine.
//  You can send values into channels from one goroutine and receive those values into another go routine
func main() {
	done := make(chan bool)

	go func(){
		fmt.Println("Series 1")
		time.Sleep(2 * time.Second)
		fmt.Println("Done go routine")
		done <- true
	}()

	fmt.Println("Series 2")
	<-done
	fmt.Println("Done with the whole application")
}