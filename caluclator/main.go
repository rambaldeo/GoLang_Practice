package main

import (
	"fmt"
	"strings"
)

func main() {
	var firstnumber, secondnumber int
	var enterAgain, functionchoice string

	//Put into a for while loops so it continues to ask the user for input until they choose to leave
	for {
		if enterAgain == "no" {
			fmt.Println("Have a nice day")
			break
		}
		fmt.Println("Enter the function you want to perform (add, subtract, multiply, divide):")
		fmt.Scanln(&functionchoice)
		functionchoice = strings.ToLower((functionchoice))
		//make sure that the user enters a valid function choice
		if functionchoice != "add" && functionchoice != "subtract" && functionchoice != "multiply" && functionchoice != "divide" {
			fmt.Println("Invalid function choice. Please enter add, subtract, multiply, or divide.")
			continue
		}
		fmt.Println("Enter the first number:")
		fmt.Scanln(&firstnumber)
		fmt.Println("Enter the second number:")
		fmt.Scanln(&secondnumber)

		calculate(functionchoice, firstnumber, secondnumber)
		fmt.Println("Do you want to enter another calculation? (yes/no)")
		fmt.Scanln(&enterAgain)
	}

}

func calculate(choice string, num1 int, num2 int) {
	var result int
	switch choice {
	case "add":
		result = num1 + num2
		fmt.Printf("The result of %d + %d is: %d\n", num1, num2, result)
	case "subtract":
		result = num1 - num2
		fmt.Printf("The result of %d - %d is: %d\n", num1, num2, result)
	case "multiply":
		result = num1 * num2
		fmt.Printf("The result of %d * %d is: %d\n", num1, num2, result)
	case "divide":
		if num2 != 0 {
			result = num1 / num2
			fmt.Printf("The result of %d / %d is: %d\n", num1, num2, result)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	}
}
