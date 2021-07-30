package main

import (
	"math"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	bc := binding.NewString()
	bf := binding.NewString()

	inputC := widget.NewEntryWithData(bc)
	inputF := widget.NewEntryWithData(bf)

	inputC.OnChanged = func(text string) {
		cDeg, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return
		}

		fDeg := math.Round(float64(cDeg)*(9.0/5.0) + 32)
		_ = bf.Set(strconv.Itoa(int(fDeg)))
	}

	inputF.OnChanged = func(text string) {
		fDeg, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return
		}

		cDeg := math.Round((float64(fDeg) - 32) * (5.0 / 9.0))
		_ = bc.Set(strconv.Itoa(int(cDeg)))
	}

	w.SetContent(container.NewGridWithColumns(4,
		inputC,
		widget.NewLabel("Celsius ="),
		inputF,
		widget.NewLabel("Fahrenheit"),
	))

	w.ShowAndRun()
}
