package main

import (
	"fmt"
	"math"

	"github.com/iabhishek91/learn_go/02_packages/utility" // note here "-" are not allowed
)

func main() {
	fmt.Println(math.Floor(2.999))
	fmt.Println(utility.ReverseString("Give me in reverse"))
}
