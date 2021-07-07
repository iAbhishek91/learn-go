package main

import "fmt"

func main() {
	// Array: are of fixed length
	// and have  length
	var fruits [2]string

	// assign value
	fruits[0] = "mango"
	fruits[1] = "grapes"
	fmt.Println(fruits)    // [mango, grapes]
	fmt.Println(fruits[0]) // mango

	// declare and assign at the same time
	fruitsArr := [2]string{"Apple", "Oranges"}
	fmt.Println(fruitsArr)    // [mango, grapes]
	fmt.Println(fruitsArr[0]) // mango

	// slice
	// with no fixed size
	fruitsSlice := []string{"Apple", "Oranges", "grapes"}
	fmt.Println(fruitsSlice)      // [mango, Oranges, grapes]
	fmt.Println(fruitsSlice[0])   // mango
	fmt.Println(fruitsSlice[1:2]) // [Oranges]

	fmt.Println(len(fruitsSlice)) //2
}
