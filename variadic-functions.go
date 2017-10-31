package main

import "fmt"

//can take an arbitrary number of arguments
func sum(nums ...int)  {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum (1, 2, 3)

	nums := []int{1, 2, 3, 4}

	//if you already have multiple arguments in a slice, apply them to a variadic function as below
	sum(nums...)
}