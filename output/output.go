package main

import "fmt"

func main() {

	var greet, who string = "hello", "world"
	fmt.Print(greet, who)
	fmt.Print(greet, " ", who)

	var i, j int = 5, 6
	fmt.Print(i, j)
	fmt.Print("\n")
	fmt.Println(greet, who)
	fmt.Println(i, j)

	val := 59
	fmt.Printf("%d\n", val)
	fmt.Printf("%+d\n", val)
	fmt.Printf("%b\n", val)
	fmt.Printf("%o\n", val)
	fmt.Printf("%x\n", val)
	fmt.Printf("%X\n", val)

	value := 3.141
	fmt.Printf("%f\n", value)
	fmt.Printf("%.2f\n", value)
	fmt.Printf("%g\n", value)
	fmt.Printf("%e\n", value)

	msg := "hi hemani"
	fmt.Printf("%s\n", msg)
	fmt.Printf("%q\n", msg)
}
