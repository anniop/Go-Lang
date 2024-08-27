package main

import "fmt"

type student struct {
	name string
	marks int64
}

func main() {

	student1 := student{"Aniket",90}
	fmt.Println("Student1: ",student1)
	student2:=student1
	fmt.Println("Student2: ",student2)

	student2.name = "Ganesh"
	student2.marks = 80

	fmt.Println("Display student after changes")
	fmt.Println("Student1",student1)
	fmt.Println("Student2",student2)

}
