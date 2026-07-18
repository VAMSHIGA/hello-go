package main 
import (
	"fmt"
	"time"
)
func hello(){
	fmt.Println("hello this goroutine")
}
func main(){
	fmt.Println("start")
	go hello()
	time.Sleep(time.Second)
	fmt.Println("end")
}