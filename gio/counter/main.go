package main

import (
	"log"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		window := app.NewWindow(
			app.Title("Counter"),
			app.MinSize(unit.Dp(140), unit.Dp(34)),
			app.Size(unit.Dp(140), unit.Dp(34)),
			app.MaxSize(unit.Dp(800), unit.Dp(400)),
		)

		if err := loop(window); err != nil {
			log.Fatal(err)
		}

		os.Exit(0)
	}()
	app.Main()
}

func loop(window *app.Window) error {
	var (
		ops   op.Ops
		theme = material.NewTheme(gofont.Collection())
	)

	for {
		evt := <-window.Events()
		switch e := evt.(type) {
		case system.DestroyEvent:
			return e.Err

		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			update(gtx)
			draw(gtx, theme)

			e.Frame(gtx.Ops)

		default:
			// log.Printf("window event: %#v\n", e)
		}
	}
}

var (
	count  = 0
	button = &widget.Clickable{}
)

func update(gtx layout.Context) {
	if button.Clicked() {
		count++
	}
}

func draw(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, material.Body1(th, strconv.Itoa(count)).Layout)
		}),
		layout.Flexed(1, material.Button(th, button, "Count").Layout),
	)
}
