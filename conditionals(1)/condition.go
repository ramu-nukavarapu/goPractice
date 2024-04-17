package main

import "fmt"

func main() {
	fmt.Println("Conditionals in go")

	value := 5
	//simple - if
	if value%2 == 0 {
		fmt.Println(value, "is a even number")
	}

	//if - else
	if value%2 == 0 {
		fmt.Println(value, "is a even number")
	} else {
		fmt.Println(value, "is a odd number")
	}

	//else - if ladder
	food := "hot"
	if food == "hot" {
		fmt.Println("Food is very Hot")
	} else if food == "cold" {
		fmt.Println("Food is very Cold")
	} else {
		fmt.Println("Food is normal")
	}

	// nested-if
	clg := true
	class := false
	if clg {
		fmt.Println("Hi from college")
		if class {
			fmt.Println("Hi from class")
		} else {
			fmt.Println("Hi from outside the class")
		}
	} else {
		fmt.Println("Hi from Home")
	}
}
