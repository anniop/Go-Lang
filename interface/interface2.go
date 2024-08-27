package main

import "fmt"

func kind(i interface{}) {
	fmt.Printf("The type of %v is %T\n",i,i)
}
func main() {

	var i int = 100
	var s string = "Jay Ganesh"
	kind(i)
	kind(s)
}
