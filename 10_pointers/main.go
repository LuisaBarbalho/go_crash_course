package main

import "fmt"

func main() {
	a := 5
	// assigns b as a pointer of a (where a is stored in the memory)
	b := &a

	fmt.Println(a, b)

	// checking its type
	fmt.Printf("%T\n", b)
	//the '*' represents a pointer

	// Use * to read value from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change value with pointer
	*b = 10
	fmt.Println(a)
}
