package main

import "fmt"

func main() {
	/*

		MAIN TYPES

		- string
		- bool
		- int
		- int  int8 int16 int32 int64
		- uint uint8 uint16 uint32 uint64 uintptr
		- byte - alias for uint8
		- rune - alias for uint32
		- float32 float64
		- complex64 complex128

	*/

	// Using var

	var name string = "Luisa"
	var age int = 24
	fmt.Println(name, age)

	//or

	var nome = "Luisa"
	var idade = 24

	fmt.Println(nome, idade)
	fmt.Printf("%T\n", name)

	// Shorthand variable (easier way)
	nomeGlobal := "Luisa Barbalho"

	t1, t2 := "Babi", "Ynaye"

	fmt.Println(nomeGlobal, t1, t2)

}
