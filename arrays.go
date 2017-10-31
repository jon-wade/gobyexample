package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//this is a built-in function which returns the length of the array
	fmt.Println("len:", len(a))

	//this is an array literal
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//arrays are only one dimensional in Go, but you can create a type which effectively creates a 2-dimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}