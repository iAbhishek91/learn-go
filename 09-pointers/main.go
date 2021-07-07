package main

import "fmt"

func main() {
	// pointers: refer to memory address
	// pointers are mostly used when we want to pass large data
	// instead pass pointers, in will improve performance.
	a := 5  // of type int
	b := &a // of type *int

	fmt.Println(a, b) // 5 0xc0000b8010
	// access value using pointers
	fmt.Println(*b)  // 5
	fmt.Println(*&a) // 5

	// change value using poointers
	*b = 6
	fmt.Println(a)
}
