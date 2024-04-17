package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("conversions using strconv package")
	//strconv -> builtin methods to convert from string to others and vice-versa

	str := strconv.Itoa(120)
	fmt.Printf("%v -> value \n %T -> type", str, str)

	num, _ := strconv.Atoi(str)
	fmt.Printf("\n%v -> value \n %T -> type", num, num)

	strNum := strconv.FormatInt(50, 10) //format -> O => S || parse -> S => O
	numStr, _ := strconv.ParseInt(strNum, 10, 64)

	fmt.Printf("\n%T,%T", strNum, numStr)
}
