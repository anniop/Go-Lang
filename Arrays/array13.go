package main

import "fmt"

func main() {

	var array[2][2] int

	array[0][0] = 10
	array[0][1] = 20
	array[1][0] = 30
	array[1][1] = 40 

	for i:=0;i<len(array);i++ {
		for j:=0;j<len(array[i]);j++ {
			fmt.Printf("%d\n", array[i][j])
		}
	}
}
