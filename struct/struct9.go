package main

import "fmt"

type author struct {

	name string
	branch string
	salary int
}

func (a author) show() {

	fmt.Println("Author's Name :", a.name)
	fmt.Println("Branch Name :", a.branch)
	fmt.Println("Salary : ", a.salary)
}

func main() {

	result := author {
		name: "Aniket",
		branch: "CSE",
		salary: 50000,
		
	}

	result.show()
}
