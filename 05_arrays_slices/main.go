package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	//Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Grape"

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])

	// Declare and assign
	namesArr := [2]string{"Lis", "Joca"}
	fmt.Println(namesArr)
	fmt.Println(namesArr[1])

	// Slices
	kidsArr := []string{"Lis", "Joca", "Pedro"}
	fmt.Println(kidsArr)
	fmt.Println(len(kidsArr))
	fmt.Println(kidsArr[0:2])

}
