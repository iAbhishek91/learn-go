package main

import "fmt"

func main() {
	x := 5
	y := 10

	// if else
	// we can use paranthesis, buts its common not to use one
	if x < y {
		fmt.Println("X < Y")
	} else {
		fmt.Println("X > Y")
	}

	//else if
	color := "red"
	if color == "red" && color != "" {
		fmt.Println("Color is red")
	} else if color == "blue" {
		fmt.Println("Color is blue")
	} else {
		fmt.Println("Color is not red or blue")
	}

	// switch
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")
	}
}
