package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) string {
	reverse := ""

	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}

	return reverse
}

func main() {
	fmt.Println("Welcome to the palindrome checker made with Go!")

	for {
		var userInput string
		fmt.Print("Please enter a string: ")
		fmt.Scanln(&userInput)

		reverse := reverseString(userInput)

		fmt.Println("The reverse of the string is:", reverse)

		if userInput == reverse {
			fmt.Println("The string is a palindrome.")
		} else {
			fmt.Println("The string is not a palindrome.")
		}

		var enterAgain string
		fmt.Print("Do you want to enter another string? (yes/no): ")
		fmt.Scanln(&enterAgain)

		if strings.ToLower(enterAgain) != "yes" {
			fmt.Println("Have a nice day!")
			break
		}
	}
}
