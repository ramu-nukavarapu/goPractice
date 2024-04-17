package main

import "fmt"

func main() {
	fmt.Println("Constants in go!")

	// Untyped constant
	const PI = 3.14 //float

	var radius = 5 //int

	// type conversion int -> float
	area := PI * (float64)(radius*radius) // float

	fmt.Println("area of circle with radius ", radius, " is ", area)
}
