package main

import "fmt"

type addition struct {

	a int
	b int

}

func (a addition) add() int {
	return a.a + a.b
}

func main() {

	var m addition
	m.a = 10
	m.b = 21

	fmt.Println("Addition is :",m.add())

}
