package main

import "fmt"

func main() {

	var employee = make(map[string]int)

	employee["Aniket"] = 100
	employee["Varun"] = 200
	employee["Rutvij"] = 300
	employee["Mansi"] = 400

	employeeList := make(map[string]int)

	fmt.Println(len(employee))
	fmt.Println(len(employeeList))
}
