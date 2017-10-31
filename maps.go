package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	//ignore the value and check whether the key is present by placing a bool in prs
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//map literal
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}