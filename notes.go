package main 
import "fmt"

func main(){
	var amount int
	fmt.Print("enter a amount:")
	fmt.Scan(&amount)

	if amount % 100!=0{
		fmt.Println("invalid amount")
		return
	}
	if amount<0{
		fmt.Println("invalid amount")
		return
	}
	hundred:=amount/100
	twohundred:=0
	fivehundered:=0
	thousand:=0

	fmt.Println("\nATM Dispense Details")
	fmt.Println("----------------------")
	fmt.Println("₹100 Notes  :", hundred)
	fmt.Println("₹200 Notes  :", twohundred)
	fmt.Println("₹500 Notes  :", fivehundered)
	fmt.Println("₹1000 Notes :", thousand)
	fmt.Println("----------------------")
	fmt.Println("Total Notes :", hundred)
}