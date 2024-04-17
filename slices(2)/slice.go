package main

import "fmt"

func main() {
	fmt.Println("Modifying the slices in go")

	// slice
	slice := []int{1, 2, 3}
	fmt.Printf("%v\n", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	// adding elements to the slice
	slice = append(slice, 4, 5, 6)
	fmt.Printf("%v\n", slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	//another slice
	anSlice := []int{7, 8, 9}
	fmt.Printf("%v\n", anSlice)
	fmt.Println(len(anSlice))
	fmt.Println(cap(anSlice))

	//adding slice to other slice
	anSlice = append(anSlice, slice...)
	fmt.Printf("%v\n", anSlice)
	fmt.Println(len(anSlice))
	fmt.Println(cap(anSlice))

	//demonstrating capacity
	arr := [6]int{1, 2, 3, 4, 5, 6}

	arSlc := arr[2:4]
	fmt.Printf("%v\n", arSlc) // {3,4}
	fmt.Println(len(arSlc))   // 2
	fmt.Println(cap(arSlc))   // 4

	arSlc = append(arSlc, 5, 6)
	fmt.Printf("%v\n", arSlc) // {3,4,5,6}
	fmt.Println(len(arSlc))   // 4
	fmt.Println(cap(arSlc))   // 4

	arSlc = append(arSlc, 7, 8, 9, 10, 11)
	fmt.Printf("%v\n", arSlc) // {3,4,5,6,7}
	fmt.Println(len(arSlc))   // 9
	fmt.Println(cap(arSlc))   // 8  -> max(cap,len(arr))
	//if the new slice exceeds the capacity it takes the actual length +2
 
}
