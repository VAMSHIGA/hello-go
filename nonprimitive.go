// array
// package main 
// import "fmt"
// func main(){
// 	var numbers [3]int=[3]int{1,2,3}
// 	fmt.Println(numbers)
// 	fmt.Println(numbers[0])

// }


// slice
// package main 
// import "fmt"
// func main(){
// 	numbers:=[]int{10,20,30}
// 	numbers=append(numbers,40)
// 	fmt.Println(numbers)
// }


// map 

// package main 
// import "fmt"
// func main(){
// 	person:=map[string]int{
// 		"Age":25,

// 	}
// 	fmt.Println(person)
// 	fmt.Println(person["Age"])
// }


// struct 



// package main 
// import "fmt"

// type student struct{
// 	Name string
// 	Age int
// }
// func main(){
// 	s:=student{
// 		Name:"vamshi",
// 		Age:25,
// 	}
// 	fmt.Println(s.Name)
// 	fmt.Println(s.Age)
// }




package main 
import "fmt"
func main(){
	age:=25
	ptr=&age 

	fmt.Println("age:",age)
	fmt.Println("ptr:",*ptr)
	fmt.Println("ptr:",ptr)
}

