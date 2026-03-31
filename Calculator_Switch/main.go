package main

import "fmt"

func main() {
	var num1, num2 int
	var operator string

	fmt.Print("Enter the first number : ")
	fmt.Scan(&num1)
	fmt.Print("Enter th second number : ")
	fmt.Scan(&num2)
	fmt.Print("Enter the operator (+,-,*,/) : ")
	fmt.Scan(&operator)

	switch operator {
	case "+":
		fmt.Println("The addition is : ", num1+num2)
	case "-":
		fmt.Println("The subtraction is : ", num1-num2)
	case "*":
		fmt.Println("The multiplication is : ", num1*num2)
	case "/":
		fmt.Println("The division is : ", num1/num2)
	default:
		fmt.Println("Invalid operator")
	}
}
