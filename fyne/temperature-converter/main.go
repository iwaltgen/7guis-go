package main

import (
	"math"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	inputC := widget.NewEntry()
	inputF := widget.NewEntry()

	inputC.OnChanged = func(text string) {
		cDeg, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return
		}

		fDeg := math.Round(float64(cDeg)*(9.0/5.0) + 32)
		inputF.Text = strconv.Itoa(int(fDeg))
		inputF.Refresh()
	}

	inputF.OnChanged = func(text string) {
		fDeg, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return
		}

		cDeg := math.Round((float64(fDeg) - 32) * (5.0 / 9.0))

		inputC.Text = strconv.Itoa(int(cDeg))
		inputC.Refresh()
	}

	w.SetContent(
		fyne.NewContainerWithLayout(layout.NewGridLayout(4),
			inputC,
			widget.NewLabel("Celsius ="),
			inputF,
			widget.NewLabel("Fahrenheit"),
		),
	)

	w.ShowAndRun()
}
