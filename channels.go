package main

import "fmt"

func main() {

	//channels are typed by the values they convey and are unbuffered by default
	messages := make(chan string)

	//send value to a channel from a goroutine (will only accept if there is a corresponding receive)
	go func() { messages <- "ping" }()

	//receive a value from a channel into the main routine
	msg := <-messages

	//channels block until both the sender and receiver are ready,
	//so this print wont execute until the communication has happened.
	fmt.Println(msg)
}
