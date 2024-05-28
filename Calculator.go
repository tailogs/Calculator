package main

import (
	"errors"
	"strconv"
	"strings"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2"
	"github.com/Knetic/govaluate"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	// Установить иконку для окна
	resource, err := fyne.LoadResourceFromPath("Calculator.ico")
	if err != nil {
		myApp.SetIcon(theme.FyneLogo())
	}
	myWindow.SetIcon(resource)

	// Установить размеры окна в два раза больше
	myWindow.Resize(fyne.NewSize(300, 250))

	output := widget.NewLabel("")
	input := ""

	updateOutput := func() {
		output.SetText(input)
	}

	invertSign := func() {
		if input == "" {
			return
		}
		if strings.HasPrefix(input, "-") {
			input = strings.TrimPrefix(input, "-")
		} else {
			input = "-" + input
		}
		updateOutput()
	}

	buttons := []struct {
		label  string
		action func()
	}{
		{"7", func() { input += "7"; updateOutput() }},
		{"8", func() { input += "8"; updateOutput() }},
		{"9", func() { input += "9"; updateOutput() }},
		{"/", func() { input += "/"; updateOutput() }},
		{"4", func() { input += "4"; updateOutput() }},
		{"5", func() { input += "5"; updateOutput() }},
		{"6", func() { input += "6"; updateOutput() }},
		{"*", func() { input += "*"; updateOutput() }},
		{"1", func() { input += "1"; updateOutput() }},
		{"2", func() { input += "2"; updateOutput() }},
		{"3", func() { input += "3"; updateOutput() }},
		{"-", func() { input += "-"; updateOutput() }},
		{"0", func() { input += "0"; updateOutput() }},
		{".", func() { input += "."; updateOutput() }},
		{"=", func() {
			result, err := eval(input)
			if err != nil {
				input = "Error"
			} else {
				input = strconv.FormatFloat(result, 'f', -1, 64)
			}
			updateOutput()
		}},
		{"+", func() { input += "+"; updateOutput() }},
		{"+/-", func() { invertSign() }},
		{"C", func() { input = ""; updateOutput() }},
		{"<-", func() {
			if len(input) > 0 {
				input = input[:len(input)-1]
				updateOutput()
			}
		}},
	}

	grid := container.NewGridWithColumns(4)
	for _, btn := range buttons {
		label := btn.label
		action := btn.action
		grid.Add(widget.NewButton(label, action))
	}

	content := container.NewVBox(
		output,
		grid,
	)

	myWindow.SetContent(content)

	// Обработчик событий клавиатуры
	myWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
		switch key.Name {
		case fyne.Key0:
			input += "0"
		case fyne.Key1:
			input += "1"
		case fyne.Key2:
			input += "2"
		case fyne.Key3:
			input += "3"
		case fyne.Key4:
			input += "4"
		case fyne.Key5:
			input += "5"
		case fyne.Key6:
			input += "6"
		case fyne.Key7:
			input += "7"
		case fyne.Key8:
			input += "8"
		case fyne.Key9:
			input += "9"
		case fyne.KeyPlus:
			input += "+"
		case fyne.KeyMinus:
			input += "-"
		case fyne.KeyAsterisk:
			input += "*"
		case fyne.KeySlash:
			input += "/"
		case fyne.KeyPeriod:
			input += "."
		case fyne.KeyReturn, fyne.KeyEnter:
			result, err := eval(input)
			if err != nil {
				input = "Error"
			} else {
				input = strconv.FormatFloat(result, 'f', -1, 64)
			}
		case fyne.KeyDelete:
			input = ""
		case fyne.KeyBackspace:
			if len(input) > 0 {
				input = input[:len(input)-1]
			}
		case fyne.KeyF9:
			invertSign()
		}
		updateOutput()
	})

	myWindow.ShowAndRun()
}

func eval(expression string) (float64, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	result, err := expr.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	switch result.(type) {
	case float64:
		return result.(float64), nil
	case int:
		return float64(result.(int)), nil
	default:
		return 0, errors.New("invalid result type")
	}
}
