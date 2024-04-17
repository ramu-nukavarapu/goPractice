package main

import "fmt"

func main() {
	// variable declaration using datatype when only declaration done
	var name string
	var rollNo int

	// variable type is inferred only when
	//declaration and initialization done at once
	var dept = "CSE"
	fav := "sai"

	//initialization
	name = "ramu"
	rollNo = 12

	fmt.Println(name)
	fmt.Println(rollNo)
	fmt.Println(dept)
	fmt.Println(fav)

	// multiple declarations
	a, b := 6, "hello"
	var c, d int = 1, 2
	fmt.Println(a, b)
	fmt.Println(c, d)
}
