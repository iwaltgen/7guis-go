package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"

	"github.com/iwaltgen/7guis-go/gio/window"
)

func main() {
	counter := newCounter(0)
	wp := window.NewProcessor(counter)

	go func() {
		window := app.NewWindow(
			app.Title("Counter"),
			app.MinSize(unit.Dp(140), unit.Dp(34)),
			app.Size(unit.Dp(140), unit.Dp(34)),
			app.MaxSize(unit.Dp(800), unit.Dp(400)),
		)

		if err := wp.Run(window); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
