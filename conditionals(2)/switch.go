package main

import "fmt"

func main() {
	fmt.Println("Switch in go")
	food := "cold"

	//single-case
	switch food {
	case "hot":
		fmt.Println("Food is Hot")
	case "cold":
		fmt.Println("Food is Cold")
	case "normal":
		fmt.Println("Food is Normal")
	default:
		fmt.Println("There is no food")
	}

	food = "very hot"
	//multi-case
	switch food {
	case "hot", "slightly hot", "very hot":
		fmt.Println("Food is Hot")
	case "cold", "slightly cold", "very cold":
		fmt.Println("Food is Cold")
	case "normal":
		fmt.Println("Food is Normal")
	default:
		fmt.Println("There is no food")
	}
}
