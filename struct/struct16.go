package main

import "fmt"

type person struct {
	name string
	id int64
}

func main() {

	person1 := person{"Ganesh", 112233}
	fmt.Println("Person1:",person1)
	person2 := &person1

	person2.name = "Mahadev"
	person2.id = 445566
	fmt.Println("Person2 : ",person2)

}
