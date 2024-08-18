package main

import "fmt"

func main() {

	for i:=0; i<=11;i++{
		fmt.Println(i)

		if i == 3 {
			break
		}
	}
}
