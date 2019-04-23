package main

import "fmt"

//intSeq returns an anonymous function
func intSeq() func() int {
	i := 0
	//the anonymous function closes over the variable i to form a closure
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
