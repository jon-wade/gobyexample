package main

import "fmt"

func main() {

	//buffered channel, accepting up to 3 strings
	messages := make(chan string, 2)

	//2 sends without corresponding receives
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}