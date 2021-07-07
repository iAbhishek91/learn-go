package main

import "fmt"

func main() {
	// range are loop through map, slices or array
	// slices and array
	ids := []int{33, 44, 55, 66, 77, 88, 99}
	for i, id := range ids {
		fmt.Printf("%d is - %d\n", i, id) //0 is - 33 and so on ...
	}

	// underscore not to use any variable
	for _, id := range ids {
		fmt.Printf("id is - %d\n", id) //id is - 33 and so on ...
	}

	// map
	friends := map[string]string{"bob": "bob@gmail.com", "abhishek": "i.abhishek.dass@gmail.com"}

	for k, v := range friends {
		fmt.Printf("%s have email - %s\n", k, v) //bob have email - bob@gmail.com and so on ...
	}

	// can omit key
	for k := range friends {
		fmt.Println("Name: " + k) //Name: bob and so on ...
	}

}
