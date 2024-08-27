package main

import "fmt"

//Student Structure
type student struct {
	name string
	age int
}

//Instant Method
func (s *student) Instant(){
	if s.name ==""{
		s.name = "Ganesh"
	}

	if s.age == 0 {
		s.age = 18
	}
}

func main() {

	student1 := student{}
	student1.Instant()
	fmt.Println(student1)

	student2 := student{name : "Mahadev"}
	student2.Instant()
	fmt.Println(student2)

	student3 := student{
		age: 34,
	}
	student3.Instant()
	fmt.Println(student3)
}
