package main

import "fmt"

type student struct {
	string
	int

}

func main() {

	student1 := student{"Aniket", 90}
	fmt.Println("Name of the student :", student1.string)
	fmt.Println("Marks of the student : ", student1.int)

}
