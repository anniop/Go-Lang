package main

import "fmt"

func main() {

	display(1234)
	display("Jay Ganesh")
	display(123.4356)
	display(true)

}

func display(a interface{}) {

	switch a.(type) {

	case string:
		fmt.Printf("Type: %T ---- Value : %v\n",a,a.(string))
	case int:
		fmt.Printf("Type: %T ---- Value : %v\n",a,a.(int))
	case float64:
		fmt.Printf("Type: %T ---- Value : %v\n",a,a.(float64))
	default:
		fmt.Println("Type Not Found")
	}
}
