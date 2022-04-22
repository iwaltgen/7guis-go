package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Temperature Converter")

	temperature := newTemperature()

	w.SetContent(temperature.Render())
	w.ShowAndRun()
}
