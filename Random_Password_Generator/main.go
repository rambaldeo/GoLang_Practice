package main

import (
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new Fyne App and Window
	a := app.New()
	myWindow := a.NewWindow("Password Generator")
	myWindow.Resize(fyne.NewSize(400, 400))

	// Title text
	title := canvas.NewText("Password Generator", theme.ForegroundColor())

	// Input box for password length
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter Password Length")

	// Text label to display the generated password
	text := canvas.NewText("", theme.ForegroundColor())

	// Button to generate password
	generateButton := widget.NewButton("Generate", func() {
		// Convert input to integer
		passwordLength, err := strconv.Atoi(input.Text)
		if err != nil || passwordLength <= 0 {
			text.Text = "Please enter a valid number"
			text.Refresh()
			return
		}
		// Generate password
		text.Text = PasswordGen(passwordLength)
		text.Refresh()
	})

	// Setup window content
	myWindow.SetContent(
		container.NewVBox(
			title,
			input,
			generateButton,
			text,
		),
	)

	// Run the window
	myWindow.ShowAndRun()
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
