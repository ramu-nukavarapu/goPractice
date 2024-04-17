package main

import "fmt"

func main() {
	fmt.Print("Operators in go\n")
	op1, op2 := 6, 2

	fmt.Println(op1 + op2)
	fmt.Println(op1 - op2)
	fmt.Println(op1 * op2)
	fmt.Println(op1 / op2)
	fmt.Println(op1 % op2)

	op1 = 4
	op2 = 2
	op1 += op2
	fmt.Println(op1, op2)

	op2 *= op2
	fmt.Println(op1, op2)

	op1 /= op2
	fmt.Println(op1, op2)

	// op1 - 1 0 0
	// op2 - 0 1 0
	// &   - 0 0 0 -> 0
	// |   - 1 1 0 -> 6
	// ^   - 1 1 0 -> 6
	fmt.Println(op1 & op2)
	fmt.Println(op1 | op2)
	fmt.Println(op1 ^ op2)
}
