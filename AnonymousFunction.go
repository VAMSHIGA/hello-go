package main 
import (
	"fmt"
	"time"
)
func main(){
	go func(){
		fmt.Println("hello this goroutine")
	}()
	fmt.Println("start")
	time.Sleep(time.Second)
	fmt.Println("end")
}	