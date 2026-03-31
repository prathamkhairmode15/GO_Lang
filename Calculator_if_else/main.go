package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scan(&b)

	var operator string
	fmt.Print("Enter the operator (+, -,* ,/) : ")
	fmt.Scan(&operator)

	if operator == "+" {
		fmt.Println("The addition is: ", a+b)
	} else if operator == "-" {
		fmt.Println("The subtraction is: ", a-b)
	} else if operator == "*" {
		fmt.Println("The multiplication is: ", a*b)
	} else if operator == "/" {
		fmt.Println("The division is: ", a/b)
	} else {
		fmt.Println("Invalid operator")
	}

}
