package main 
import(
	"fmt"
	"sync"
)
func display(wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("hello fished goroutine")
}
func main(){
	
	var wg sync.WaitGroup
	wg.Add(1)
	go display(&wg)
	wg.Wait()

}