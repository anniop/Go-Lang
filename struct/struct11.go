// Method with a pointer reciever

package main

import "fmt"

type author struct {

	name string 
	branch string
}

func (a *author)show(abranch string) {
	(*a).branch = abranch

}


func main() {

	result := author{
		name:"Aniket",
		branch: "CSE",
	}
	fmt.Println("Before Changing")
	fmt.Println("Author's Name :",result.name)
	fmt.Println("Author's Branch :" ,result.branch)

	p:= &result
	p.show("ECE")

	fmt.Println("After Changing")
	fmt.Println("Author's Name :",result.name)
	fmt.Println("Author's Branch :" ,result.branch)
	

}
