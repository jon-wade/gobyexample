package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0

	//range returns index and value, here we are ignoring the index with the _, notation
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//for maps, returns key and value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//can also iterate over just keys
	for k := range kvs {
		fmt.Println("key", k)
	}

	//range iterating over strings, i gives index of the Unicode rune and c the rune itself (Unicode value of char)
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}