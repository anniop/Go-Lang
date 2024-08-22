package main

import "fmt"

func main() {

	var a int = 20
	var b *int = &a

	fmt.Printf("Address of the variable : %x\n", &a)
	fmt.Printf("Address of the variable stored in p : %x\n",b)

	fmt.Printf("value in *b : %d\n",*b)

}
