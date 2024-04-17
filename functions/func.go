package main

import "fmt"

func main() {
	fmt.Println("functions in go")

	//function with no parameters
	welcome() // function call with no arguments

	name := "hemani"

	//function with parameters
	invite(name) // function call with one argument

	val1 := 23
	val2 := 56

	res := add(val1, val2)
	fmt.Printf("%v + %v = %v\n", val1, val2, res)

	res1, res2 := addANDsub(val1, val2)
	fmt.Printf("%v + %v = %v\n%v - %v = %v", val1, val2, res1, val1, val2, res2)
}

func welcome() {
	fmt.Println("Hey Everyone! Welcome to the session...")
}

func invite(name string) {
	fmt.Println("Hey", name, "! I'm glad to see here...")
}

// function with return
func add(num1 int, num2 int) int {
	return num1 + num2
}

// function with naked return
func addN(num1 int, num2 int) (result int) {
	result = num1 + num2
	return
}

//function which returns multiple values
func addANDsub(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

//function which naked returns multiple values
func addANDsubN(num1 int, num2 int) (res1 int, res2 int) {
	res1 = num1 + num2
	res2 = num1 - num2
	return
}
