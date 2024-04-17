package main

import "fmt"

// declaring a structure
type student struct {
	name   string
	rollNo int
	branch string
	CGPA   float32
}

func main() {
	fmt.Println("structs in go")

	// creating an instance for the structure
	var stud1 student

	// adding the information of a student
	stud1.name = "ramu"
	stud1.rollNo = 18
	stud1.branch = "CSE"
	stud1.CGPA = 9.06

	fmt.Println("Name of the student:", stud1.name)
	fmt.Println("Roll No:", stud1.rollNo)
	fmt.Println("Branch:", stud1.branch)
	fmt.Println("CGPA:", stud1.CGPA)
}
