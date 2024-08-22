package main

import "fmt"

func addition(x int, y int)int {

	result := x + y
	return result
}

func main() {

	var iRet int = 0
	iRet = addition(12,43)
	fmt.Println("The Addition is : ",iRet)
}
