package main

import "fmt"

// Rectangle Interface 
type rectangle interface {
	Perimeter() float64
	Area() float64
}

// Calculate Struct 
type calculate struct {
	length float64
	width float64 
}

// Perimeter Method
func (c calculate) Perimeter() float64 {
	return 2 * (c.length + c.width)
}

func (c calculate) Area() float64 {
	return c.length * c.width
}

func main() {
	
	var len float64
	var wid float64
	fmt.Printf("Enter the length : ")
	fmt.Scanf("%f",&len)
	fmt.Printf("Enter the width : ")
	fmt.Scanf("%f",&wid)
	var r rectangle
	r = calculate{
		length: len,
		width: wid,
	}

	fmt.Println("Perimeter of the Rectangle is : ", r.Perimeter())

	fmt.Println("Area of the Rectangle is : ", r.Area())
	

}
