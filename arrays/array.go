package main

import "fmt"

func main() {
	fmt.Println("Arrays in go")
	//using a number
	var arr [5]int
	fmt.Println(arr)

	//using constant as a size
	const SIZE = 5
	var arCons [SIZE]int
	fmt.Println(arCons)

	array := [...]int{}
	fmt.Println(len(array))
	fmt.Println(array)

	arrInfer := [...]int{1, 2}
	fmt.Println(arrInfer)
	fmt.Println(len(arrInfer))

	for index, value := range arr {
		fmt.Println(index, value)
	}
}

/* for (int x : arr)
{
	value
}
*/
