package main

import "fmt"

func main() {
	var value interface{} = "Jay Shrii Ganesh !!"
    
	switch a:= value.(type){
	case int64:
	     fmt.Println("value is int64",a)
	case float64:
	fmt.Println("value is float64",a)

	case string:
	fmt.Println("value is string",a)

	default:
	 fmt.Println("Type is unknown")
	
	}
}
