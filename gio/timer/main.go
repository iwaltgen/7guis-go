package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

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
			app.Title("Timer"),
			app.MinSize(unit.Dp(380), unit.Dp(200)),
			app.Size(unit.Dp(380), unit.Dp(200)),
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

const (
	sliderMin = 0
	sliderMax = 300
)

var (
	started  = time.Now()
	elapsed  = 0 * time.Millisecond
	duration = 15 * time.Second
	slider   = &widget.Float{Value: sliderMax / 2}
	button   = &widget.Clickable{}
)

func update(gtx layout.Context) {
	if button.Clicked() {
		started = time.Now()
	}
	elapsed = time.Since(started)

	if slider.Changed() {
		duration = time.Duration(math.Round(float64(slider.Value))*100) * time.Millisecond
	}
	op.InvalidateOp{}.Add(gtx.Ops)
}

func progress() float32 {
	if duration == 0 {
		return 100
	}
	return float32(elapsed) / float32(duration)
}

func elapsedSeconds() string {
	if duration < elapsed {
		return fmt.Sprintf("%v", duration)
	}

	elapsedMillis := time.Duration(elapsed.Milliseconds() / 100 * 100)
	return fmt.Sprintf("%v", elapsedMillis*time.Millisecond)
}

func draw(gtx layout.Context, th *material.Theme) layout.Dimensions {
	insetUnit := unit.Dp(6)
	return layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceBetween,
	}.Layout(gtx,
		layout.Flexed(3, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{}.Layout(gtx,
				layout.Flexed(1, uniformInset(insetUnit, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceAround,
					}.Layout(gtx,
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.W.Layout(gtx, material.Body1(th, "Elapsed Time:").Layout)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.W.Layout(gtx, material.Body1(th, "").Layout)
						}),
						layout.Rigid(func(gtx layout.Context) layout.Dimensions {
							return layout.W.Layout(gtx, material.Body1(th, "Duration:").Layout)
						}),
					)
				})),
				layout.Flexed(2, uniformInset(insetUnit, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Axis:    layout.Vertical,
						Spacing: layout.SpaceAround,
					}.Layout(gtx,
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.W.Layout(gtx, material.ProgressBar(th, progress()).Layout)
						}),
						layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
							return layout.W.Layout(gtx, material.Body1(th, elapsedSeconds()).Layout)
						}),
						layout.Flexed(1, material.Slider(th, slider, sliderMin, sliderMax).Layout),
					)
				})),
			)
		}),
		layout.Flexed(1, uniformInset(insetUnit, material.Button(th, button, "Reset Timer").Layout)),
	)
}

func uniformInset(inset unit.Value, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(inset).Layout(gtx, widget)
	}
}
