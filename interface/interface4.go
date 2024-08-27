package main

import "fmt"

func main() {

	var result = "Jay Ganesh"
	display(result)

	var result2 = 123
	display(result2)

}

func display(a interface{}) {
	value,ok := a.(string)
	fmt.Printf("Value : %v ---- Status is %v\n",value,ok)
}
