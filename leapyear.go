package main
import (
	"fmt"
	"log"
	
)
func checkleapyear(year int) {
	if (year%400==0) || !(year%100==0) && (year%4==0){
	fmt.Println("The year is leap")
	}else{
	fmt.Println("The year is not leap")
	}
}
func main(){
	log.Println("Welcome to check leap year program")
	fmt.Println("Enter the year:")
	var year int
	fmt.Scanf("%d",&year)
	checkleapyear(year)

	
}