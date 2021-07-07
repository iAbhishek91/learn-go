// MAIN DATA TYPES
// string
// bool
// int
// int, int8, ini16, int32, int 64
// uint, uint8, uint16, uint32, uint 64 uintptr
// byte - alias for uint8
// rune - alias of int32
// float32, float64
// complex64, complex128

package main

import "fmt"

func main() {
	// VARIABLES
	// RULE-1: mentioning a datatype is optional, it can infer from the data
	// the below line can be rewritten as
	// var name = "abhishek"
	// RULE-2 unused variable are not allowed
	var name string = "abhishek"
	fmt.Println(name)

	// RULE-3 if requried we can overrite the default datatype
	// var age float32 = 32
	var age = 31
	// RULE-4, no type() function, %T is used to find the type of the variable
	fmt.Printf("%T\n", age)

	// RULE-5 shorthand declaration, can be used ONLY within a function
	// can declare multiple values
	// city := "London"
	// borough := "Tower of Hamlet"
	city, borough := "London", "Tower of Hamlet"
	fmt.Printf("%s lives in %s, %s\n", name, city, borough)

	// CONSTANTS
	// RULE-1: cant redefine consts
	const is_go_easy = true
	fmt.Printf("Is go programming language great? %t\n", is_go_easy)
}
