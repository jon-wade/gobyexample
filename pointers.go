package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// *int is a int pointer type, *iptr dereferences the pointer (finds its value) and amends it
func zeroptr(iptr *int)  {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	//zeroval gets a copy of the value of i, not a reference to the actual memory location of i, so it does not affect the value of i
	zeroval(i)
	fmt.Println("zeroval:", i)

	//&i passes the memory address of variable i (which is then dereferenced in the function and the value is altered)
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	//this prints the address in memory of i
	fmt.Println("pointer:", &i)
}