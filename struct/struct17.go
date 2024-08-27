package main

import "fmt"

func main() {

	student:=struct {

		name string
		id int64
		age int32
	}{
		name: "Aniket",
		id: 112233,
		age:21,
	}
	fmt.Println(student)

}
