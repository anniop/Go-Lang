package main

import ("fmt"
		"errors")

func main(){
	var iRet int = 0
	iRet = division(12, 0)
	fmt.Println(iRet)
}

func division(a int, b int) (int, error) {
    var err error
	if b == 0{
		err = errors.New("Cannot Divide by Zero")
		return 0, err
	}
	var result int = a / b
	return result

}
