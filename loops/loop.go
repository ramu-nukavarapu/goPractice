package main

import "fmt"

func main() {
	fmt.Println("loops in go")

	// for loop
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

}
