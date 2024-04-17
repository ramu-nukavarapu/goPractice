package main

import "fmt"

func main() {
	var num8 int8 = 127 // -128 to +127
	var num16 int16 = 255
	var num32 int32 = 1234
	var char rune = 'a'
	var num int = 234567

	var unum8 uint8 = 255 //0 to 255
	var byt byte = 123
	var unum16 uint16 = 1213
	var unum32 uint32 = 1122
	var unum uint = 12334456778

	var dec float32 = 2112.342654747282
	var deci float64 = 2112.342654747282

	var str string = "hello world!"

	fmt.Println(num8)
	fmt.Println(num16)
	fmt.Println(num32)
	fmt.Println(char)
	fmt.Println(num)

	fmt.Println(unum8)
	fmt.Println(byt)
	fmt.Println(unum16)
	fmt.Println(unum32)
	fmt.Println(unum)

	fmt.Println(dec)
	fmt.Println(deci)

	fmt.Println(str)
}
