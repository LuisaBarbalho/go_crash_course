package main

import "fmt"

func main() {

	// Define map

	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["Sarah"] = "sarah@gmail.com"
	emails["Jack"] = "jack@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add key values
	// first type is for the keys and second is for the values
	ages := map[string]int{"Bob": 12, "Sarah": 23, "Jack": 45}
	fmt.Println(ages)
}
