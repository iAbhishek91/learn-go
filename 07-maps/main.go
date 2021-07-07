package main

import "fmt"

func main() {
	// maps are key value pair
	email := make(map[string]string)

	// assign values
	email["Abhsiehk"] = "i.abhishek.dass@gmail.com"
	email["bob"] = "bob.martin@gmail.com"
	fmt.Println(email)        // map[Abhsiehk:i.abhishek.dass@gmail.com bob:bob.martin@gmail.com]
	fmt.Println(len(email))   // 2
	fmt.Println(email["bob"]) // bob.martin@gmail.com

	// delete
	delete(email, "bob")
	fmt.Println(email)      // map[Abhsiehk:i.abhishek.dass@gmail.com]
	fmt.Println(len(email)) // 1

	// declare and assign at same time
	friend := map[string]string{"bob": "bob@gmail.com", "abhishek": "i.abhishek.dass@gmail.com"}
	fmt.Println(friend) // map[abhishek:i.abhishek.dass@gmail.com bob:bob@gmail.com]
}
