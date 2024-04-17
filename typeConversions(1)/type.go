package main

import "fmt"

func main() {
	var val int = 97
	fmt.Printf("value -> %v type -> %T \n", val, val)
	var num float32 = float32(val)
	fmt.Printf("value -> %v type -> %T\n", num, num)
	number := string(val)
	fmt.Printf("value -> %v type -> %T\n", number, number)
}
