package main

import "fmt"

func main() {
	
	var result = "Jay Ganesh"
	display(result)

}

func display(a interface{}){

	value := a.(string)
	fmt.Println("Value : ",value)
}
