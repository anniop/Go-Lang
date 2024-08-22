package main
import "fmt"
func swap(x,y int) int {

	var tmp int

	tmp = x
	x = y
	y = tmp
	return tmp
}

func main() {

	var a int = 100
	var b int = 200


	fmt.Printf("Before swap , A = %d , B = %d\n",a,b)
	swap(a,b)
	fmt.Printf("After swap , A = %d, B = %d\n",a,b)

}
