package main
import (
	"time"
	"fmt")
func main() {

	fmt.Println("Start Main Function ........")
	ch := make(chan int)
	go send(ch)

	go Recieve(ch)
	time.Sleep(time.Microsecond*100)
	
	fmt.Println("End of Main Function..........")
	
}

func send(ch chan int){
	ch <- 50
}

func Recieve(ch chan int){
	fmt.Println(200 + <-ch)
}
