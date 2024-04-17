package main

import "fmt"

func main() {

	fmt.Println("Creation of Maps in go")

	// using var keyword
	var exMap = map[int]int{1: 100, 2: 200, 3: 300}
	fmt.Printf("%T\n", exMap)

	// := operator
	maping := map[int]string{1: "mon", 2: "tue", 3: "wed"}
	fmt.Printf("%T\n", maping)

	// using make function
	makeMap := make(map[int]string) // using make() creates an empty

	//inserting elements in the map using key
	makeMap[1] = "A"
	makeMap[2] = "B"
	makeMap[3] = "C"

	for key, value := range makeMap {
		fmt.Printf("%v -> key %v -> value \n", key, value)
	}
}
