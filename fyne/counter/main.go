package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Counter")

	counter := newCounter(0)

	w.SetContent(counter.Render())
	w.ShowAndRun()
}
