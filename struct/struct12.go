package main

import "fmt"

type author struct {
	name string 
	branch string
}

func (a *author)show1(abranch string){
	(*a).branch = abranch
}

func (a author) show2() {
	a.name = "Ganesh"
	fmt.Println("Before Changing")
	fmt.Println("Author's Name : ",a.name)
}

func main() {

	result := author {
		name: "Mahadev",
		branch: "CSE",
	}

	fmt.Println("Branch Name(Before):",result.branch)

	result.show1("IT")
	fmt.Println("Branch Name After Changing :",result.branch)

	(&result).show2()
	fmt.Println("Author's Name After Changing : ",result.name)
}
