package main

import "fmt"

func main() {

	value1:= 100
	value2:= 200

	p1 := &value1
	p2 := &value2

	p3 := &value1

	result1 := p1 == p2
	fmt.Println("Result1 : Is P1 equals P2", result1)

	result2 := p1 == p3
	fmt.Println("Result2 : Is P1 equals P3",result2)
     
	result3 := p2 == p3
    fmt.Println("Result3 : Is P2 equals P3",result3)

	result4 := &p1 == &p3
	fmt.Println("Result4 : Is P1 value equals P3 value",result4)

}
