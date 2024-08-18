package main

import "fmt"

func main() {
	var student1 string = "Ganesh" // type is string
	var student2 = "Aniket"        // type is inferred
	x := 2                         // type is inferred
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
}

// inferred means compiler will decide the type of the variable.
