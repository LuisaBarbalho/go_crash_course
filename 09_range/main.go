package main

import "fmt"

func main() {
	ids := []int{22, 34, 55, 65, 34, 54, 2}

	//Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum of ids: %d\n", sum)

	// Range with maps
	ages := map[string]int{"Bob": 12, "Sarah": 23, "Jack": 45}

	for key, value := range ages {
		fmt.Printf("%s: %d\n", key, value)
	}
}
