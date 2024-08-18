package main

import "fmt"

func main() {
	var today int = 2

	switch  {
		case today==1:
			fmt.Println("Monday")
		case today==2:
			fmt.Println("Tuesday")
		case today==3:
			fmt.Println("Wednesday")
		default:
			fmt.Println("Invalid day")
	}
}
