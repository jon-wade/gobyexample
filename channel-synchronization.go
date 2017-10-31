package main

import "fmt"
import "time"

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 2)
	fmt.Println("done")

	// Send a value to notify that we're done.
	done <- true
}

func main() {

	//start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	//block until we receive a notification from the worker on the channel
	//without this receiver the program would exit before the worker had begun
	<-done
}