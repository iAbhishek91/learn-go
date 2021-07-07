package main

import "fmt"

func greeting(name string) string {
	return "hello " + name
}

// since both are same we can mention only one int in the definition
func add(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Abhishek"))
	fmt.Println(add(5, 5))
}
