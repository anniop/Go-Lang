package main

import "fmt"

func main() {
	arr2 :=[2][2] int {
		{1,2},
		{3,4},
	}
	fmt.Println("value of array is")

	for i:=0;i<len(arr2);i++ {
		for j:=0;j<len(arr2[i]);j++{
			fmt.Printf("%d\n",arr2[i][j])
		}
	}

}
