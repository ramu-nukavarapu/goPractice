package main

import "fmt"

func main() {
	fmt.Println("Creation of slices in go")

	//similar to array declaration
	var slice = []int{1, 2, 3}
	fmt.Printf("%T\n", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//slicing the array using : operator
	arr := [8]int{11, 12, 13, 14, 15, 16, 17, 18}

	arrSlice := arr[3:6]
	fmt.Printf("%v\n", arrSlice) //14,15,16
	fmt.Println(len(arrSlice))   //3
	fmt.Println(cap(arrSlice))   //5

	//using make()
	makeSlice := make([]int, 5, 8)
	fmt.Println(len(makeSlice)) //5
	fmt.Println(cap(makeSlice)) //8

	anSlice := make([]int, 5)
	fmt.Println(len(anSlice)) //5
	fmt.Println(cap(anSlice)) //5
}
