// Structs are most important, there is no class in Go,
// hence struct is MOST IMPORTANT data structure

package main

import (
	"fmt"
	"strconv"
)

// define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	// alternative way to declare attributes to make it short
	firstName, lastName, city, gender string
	age                               int
}

// define methods(value receivers) this will be attached to struct
// value receiver are basically getters, which can access existing values
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " of age: " + strconv.Itoa(p.age) // NOTE this conversion
}

// define methods(pointer receivers) this will be attached to struct
// pointer receiver are basically setters, which can change values
func (p *Person) hasBirthday() {
	p.age++
}

func main() {
	// init person using struct
	rob := Person{
		firstName: "Robert",
		lastName:  "Martin",
		city:      "London",
		gender:    "m",
		age:       30,
	}

	// alternative init
	bob := Person{"Bob", "Daylen", "WC", "m", 45}

	fmt.Println(rob)      // {Robert Martin London m 30}
	fmt.Println(bob.city) // WC

	// update property
	bob.age++
	fmt.Println(bob.age) // 46

	fmt.Println(bob.greet()) // Hello, my name is Bob of age: 46
	bob.hasBirthday()
	fmt.Println(bob.age) // 47
}
