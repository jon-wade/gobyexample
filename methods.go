package main

import "fmt"

//methods can be declared on structs (which are like OO classes)
type rect struct {
	width, height int
}

//this area method has a pointer receiver on the rect struct
func (r *rect) area() int {
	return r.width * r.height
}

//this perim method has a value receiver on the rect struct
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

//IMPORTANT: Go automatically handles conversion between values and pointers for method calls
func main() {
	r := rect{width: 10, height: 5}

	//calling a method on the struct r with a value type, accessing it like a property using dot notation
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	//calling methods on the struct r with a pointer type
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
