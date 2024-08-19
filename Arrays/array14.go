// First Case :- If an array does not initialized explicitly then the default value of this array is 0.

package main

import "fmt"

func main() {
	var myarray[2] int

	fmt.Println(myarray)
}

// OutPut :- [0 0]
