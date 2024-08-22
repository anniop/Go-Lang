package main

import "fmt"

func main() {

	var a int = 10

	var b *int = &a
	var c **int = &b

	fmt.Printf("Value in b pointer is : %d\n",*b)
	fmt.Printf("Value in c pointer is : %d\n",**c)
	fmt.Printf("Value of b pointer is : %d\n",b)
	fmt.Printf("Value of c pointer is : %d\n",c)


}
