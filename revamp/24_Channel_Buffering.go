package main

import "fmt"

func main() {
	message := make(chan string, 3)

	message <- "buffered"
	message <- "channel"
	message <- "Kigali"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}