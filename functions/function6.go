package main

import "fmt"

func swap1(x *int,y *int) {

	var tmp int

	tmp = *x

	*x = *y
	*y = tmp

}

func main() {

	var a int = 100
	var b int = 200


	fmt.Printf("Before swap , A = %d , B = %d\n",a,b)
	swap1(&a,&b)
	fmt.Printf("After swap , A = %d, B = %d\n",a,b)
}
