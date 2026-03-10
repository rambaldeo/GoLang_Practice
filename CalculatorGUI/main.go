package main

//using https://github.com/sd4483/GO-GUI-App-Simple-Calculator as reference
import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func convertToInt(s string) int {
	i, e := strconv.Atoi(string(s))
	_ = e
	return i
}
func main() {
	a := app.New()
	w := a.NewWindow("Calculator")

	num_1 := widget.NewEntry()
	num_1.SetPlaceHolder("Enter the First number")

	num_2 := widget.NewEntry()
	num_2.SetPlaceHolder("Enter the second number")

	content_1 := binding.NewString()
	content_1.Set("")

	content_2 := binding.NewString()
	content_2.Set("")

	result := binding.NewString()
	result.Set("")

	text_1 := widget.NewLabelWithData(content_1)
	text_2 := widget.NewLabelWithData(content_2)

	result_text := widget.NewLabelWithData(result)
	result_text.TextStyle.Bold = true

	//Adding the buttons for the calculator
	add_button := widget.NewButton("Add", func() {
		content_1.Set("The first number is: " + num_1.Text)
		content_2.Set("the second number is: " + num_2.Text)

		n1 := convertToInt(num_1.Text)
		n2 := convertToInt((num_2.Text))
		add := n1 + n2
		result.Set("The reuslt is: " + strconv.Itoa(add))
	})
	layout_1 := container.NewVBox(num_1, num_2)
	layout_2 := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), add_button, layout.NewSpacer())
	layout_3 := container.NewVBox(text_1, text_2, result_text)

	w.SetContent(container.New(layout.NewVBoxLayout(), layout_1, layout_2, layout_3))
	w.Show()
	a.Run()

}
