package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	// Prompt the user for the desired password length
	var inputFromUser int 
	fmt.Println("Enter the desired password length:")
	fmt.Scanln(&inputFromUser)

	GeneratedPassword := PasswordGen(inputFromUser)
	println("Generated Password:", GeneratedPassword)
}

// PasswordGen generates a random password of specified length
func PasswordGen(passwordLength int) string {
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{}|;:,.<>?/"

	password := ""
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < passwordLength; i++ {
		switch rng.Intn(4) {
		case 0:
			password += string(lowerCase[rng.Intn(len(lowerCase))])
		case 1:
			password += string(upperCase[rng.Intn(len(upperCase))])
		case 2:
			password += string(digits[rng.Intn(len(digits))])
		case 3:
			password += string(specialChars[rng.Intn(len(specialChars))])
		}
	}
	return password
}
