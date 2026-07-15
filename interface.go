package main
import "fmt"
type animal interface{
	speak()
}
type dog struct{}
func(d dog)speak(){
	fmt.Println("woof")
}
func main(){
	var a animal=dog{}
	a.speak()
}