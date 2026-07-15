package main

import "fmt"

func main() {
    age := 25
    ptr := &age

    fmt.Println("Age:", age)
    fmt.Println("Address:", ptr)
    fmt.Println("Value:", *ptr)
}
