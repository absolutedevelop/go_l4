package main 

import "fmt"


func addNumbers(num1 int, num2 int) int {
	var sum = num1 + num2

	return sum
}


func main(){

	var value1 int = 10
	var value2 int = 20
	var values_sum int = addNumbers(value1,value2)

	fmt.Println("The sum is", values_sum)

}