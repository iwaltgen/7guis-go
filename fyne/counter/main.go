package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	count := 0
	value := widget.NewLabel("0")
	w.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		value,
		widget.NewButton("Count", func() {
			count++
			value.SetText(strconv.Itoa(count))
		}),
	))
	w.ShowAndRun()
}
