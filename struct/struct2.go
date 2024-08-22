package main

import "fmt"

type address struct {
	state string
	city string
	zipcode int
}

func main() {

	var a address // var keyword to create instance of struct
	fmt.Println(a)

}
