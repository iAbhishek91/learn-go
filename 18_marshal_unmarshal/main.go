// Converting Go Object tinto JSON is called marshalling, and vice-versa

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string
	Age     int
	Address string
}

func main() {
	abhi := Employee{Name: "Abhishek", Age: 30, Address: "UK"}

	// convert struct go object to json
	// the value returned is byte data
	jsonData, e := json.Marshal(abhi)
	if e != nil {
		panic(e)
	}

	// convert byte data to string
	jsonString := string(jsonData)
	fmt.Println(jsonString)

	// convert string to byte
	jsonByteData := []byte(jsonString)

	// convert byteData to go object
	var emp Employee
	json.Unmarshal(jsonByteData, &emp)
	fmt.Println(emp.Name)
	fmt.Println(emp.Address)
	fmt.Println(emp.Age)
}
