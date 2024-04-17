package main

import "fmt"

func main() {
	fmt.Println("Modifying maps in go")

	numbers := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Printf("%v\n", numbers)

	// accessing the elements in the map
	value := numbers[2]
	fmt.Println("the value at key 2 is", value)

	// adding or updating elements in the map
	numbers[3] = 33 //updation
	numbers[4] = 40 //adding

	fmt.Printf("%v\n", numbers)

	// removing elements from the map
	delete(numbers, 3)
	fmt.Printf("%v\n", numbers)

	// checking the element is present or not
	val1, isContains1 := numbers[4]
	fmt.Printf("value -> %v and isContains -> %v\n", val1, isContains1)

	val2, isContains2 := numbers[3]
	fmt.Printf("value -> %v and isContains -> %v\n", val2, isContains2)

	//iterate over maps
	for key, valu := range numbers {
		fmt.Printf("key -> %v value -> %v\n", key, valu)
	}
}
