package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	count := 0
	value := widget.NewLabel("0")
	w.SetContent(container.NewGridWithColumns(2,
		value,
		widget.NewButton("Count", func() {
			count++
			value.SetText(strconv.Itoa(count))
		}),
	))
	w.ShowAndRun()
}
