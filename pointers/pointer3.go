package main

import "fmt"

func main() {

	var a = 458
	var p = &a
	fmt.Println("Value stored in the a variable is : ",a)
	fmt.Println("Address of a variable : ",&a)
	fmt.Println("Value stored in the p pointer is : ",p)
	fmt.Println("Value in p pointer before changing : ", *p)

	*p = 500
	fmt.Println("Value in p pointer after changing : ", *p)

}
