package main

import (
	"log"
	"math"
	"os"
	"strconv"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
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
			app.Title("Temperature Converter"),
			app.MinSize(unit.Dp(360), unit.Dp(34)),
			app.Size(unit.Dp(360), unit.Dp(34)),
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

		case key.Event:
			if e.Name == key.NameTab {
				if cEditor.Focused() {
					fEditor.Focus()
				} else {
					cEditor.Focus()
				}
			}

		default:
			// log.Printf("window event: %#v\n", e)
		}
	}
}

var (
	cDeg    int64
	cEditor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	fDeg    int64
	fEditor = &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
)

func update(gtx layout.Context) {
	for _, e := range cEditor.Events() {
		if _, ok := e.(widget.ChangeEvent); ok {
			text := cEditor.Text()
			deg, err := strconv.ParseInt(text, 10, 64)
			if err != nil || cDeg == deg {
				continue
			}

			cDeg = deg
			deg = int64(math.Round(float64(cDeg)*(9.0/5.0) + 32))
			if fDeg == deg {
				continue
			}

			fDeg = deg
			fEditor.SetText(strconv.FormatInt(fDeg, 10))
		}
	}

	for _, e := range fEditor.Events() {
		if _, ok := e.(widget.ChangeEvent); ok {
			text := fEditor.Text()
			deg, err := strconv.ParseInt(text, 10, 64)
			if err != nil || fDeg == deg {
				continue
			}

			fDeg = deg
			deg = int64(math.Round(float64(fDeg-32) * (5.0 / 9.0)))
			if cDeg == deg {
				continue
			}

			cDeg = deg
			cEditor.SetText(strconv.FormatInt(cDeg, 10))
		}
	}
}

func draw(gtx layout.Context, th *material.Theme) layout.Dimensions {
	insetUnit := unit.Dp(6)
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, uniformInset(insetUnit, material.Editor(th, cEditor, "Celsius").Layout)),
		layout.Rigid(uniformInset(insetUnit, material.Body1(th, "Celsius = ").Layout)),
		layout.Flexed(1, uniformInset(insetUnit, material.Editor(th, fEditor, "Fahrenheit").Layout)),
		layout.Rigid(uniformInset(insetUnit, material.Body1(th, "Fahrenheit").Layout)),
	)
}

func uniformInset(inset unit.Value, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(inset).Layout(gtx, widget)
	}
}
