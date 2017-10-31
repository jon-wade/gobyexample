package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	//calling the function normally
	f("direct")

	//calling the function in a new thread (goroutine) which runs asynchronously
	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	//Scanln requires we press a key before the program continues
	var input string
	//TODO: not sure why you have to dereference input here when using Scanln
	fmt.Scanln(&input)
	fmt.Println("done")
}
