package main

import "time"
import "fmt"

func main() {

	timer1 := time.NewTimer(time.Second * 2)

	// <-timer1.C blocks on the timer's channel `C` until it sends a value indicating that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	//if you just wanted to wait, use time.Sleep

	//stopping a timer
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
