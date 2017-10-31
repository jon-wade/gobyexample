package main

import "time"
import "fmt"

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	//select proceeds with the first case that is ready
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	//allowing a longer timeout of 3s, then the receive
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	//TODO: why does this time.After behave like a channel?
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
