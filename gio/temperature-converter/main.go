package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"

	"github.com/iwaltgen/7guis-go/gio/window"
)

func main() {
	temperature := newTemperature(0)
	wp := window.NewProcessor(temperature)

	go func() {
		window := app.NewWindow(
			app.Title("Temperature Converter"),
			app.MinSize(unit.Dp(360), unit.Dp(34)),
			app.Size(unit.Dp(360), unit.Dp(34)),
			app.MaxSize(unit.Dp(800), unit.Dp(400)),
		)

		if err := wp.Run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
