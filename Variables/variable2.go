package main 

import "fmt"

func main() {
	var i float32 = 34.23
	var b float32 = 12.45
	c := i - b

	fmt.Printf("i = %f\n", i)
	fmt.Printf("b = %f\n", b)
	fmt.Printf("c = %f\n", c)
	fmt.Printf("%f - %f = %f\n", i, b, c)
	fmt.Printf("type of c is %T\n", c)

	var num1 int = 0
	fmt.Printf("type of num1 is %T\n", num1)
	
}
