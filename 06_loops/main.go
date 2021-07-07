package main

import "fmt"

func main() {
	i := 1
	// long method
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// short method
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// fizzbuzz challenge - common interview questions
	for i := 1; i <= 20; i++ {
		if i%3 == 0 {
			fmt.Println("i = %d, fizz", i)
		}

	}
}
