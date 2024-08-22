package main

import "fmt"

func main() {

	Ganesh := func (x int, y int)int{
		result := x + y
		return result
	}

	fmt.Println(Ganesh(12,32))
}
