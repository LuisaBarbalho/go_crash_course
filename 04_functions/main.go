package main

import "fmt"

// func name_of_the_function(name_of_the_parameter type_of_the_parameter) type_of_the_return
func greeting(name string) string {
	return "Hello " + name
}

func soma(t1, t2 int) int {
	return t1 + t2
}

func main() {
	fmt.Println(greeting("Luisa"), soma(25, 26))
}
