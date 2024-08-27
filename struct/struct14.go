package main

import "fmt"

type student struct {
	name string 
	age int
}


func main() {
	
		student1 := student{
			name: "Ganesh",
			age: 23,
	}
	student2 := student {
		name: "Ganesh",
		age: 24,
	}

	if student1 ==student2 {
		fmt.Println("Variable student 1 is equal to variable student 2")
	} else{
		fmt.Println("Variable student 1 is not equal to variable student 2")
	}

}
