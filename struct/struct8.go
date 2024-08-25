// Nested Structure
package main

import "fmt"

type author struct {

	name string
	branch string
	year int
}

type HR struct {

	details author

}

func main() {

	result := HR{
		details: author{
			name: "Aniket",
			branch: "CSE",
			year: 4,
		},
	}

	fmt.Println("Details of the Author")
	fmt.Println(result)

}
