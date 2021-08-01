package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var (
	duration        = binding.NewFloat()
	elapsedPercents = binding.NewFloat()
	elapsedString   = binding.NewString()
)

func main() {
	a := app.New()
	w := a.NewWindow("Timer")

	_ = duration.Set(float64(15 * time.Second))

	resetCh := frame(16 * time.Millisecond)

	w.SetContent(container.NewVBox(
		container.NewGridWithColumns(2,
			container.NewGridWithRows(3,
				widget.NewLabel("Elapsed Time:"),
				widget.NewLabel(""),
				widget.NewLabel("Duration:"),
			),
			container.NewGridWithRows(3,
				widget.NewProgressBarWithData(elapsedPercents),
				widget.NewLabelWithData(elapsedString),
				widget.NewSliderWithData(0, float64(30*time.Second), duration),
			),
		),
		widget.NewButton("Reset Timer", func() {
			resetCh <- struct{}{}
		}),
	))
	w.ShowAndRun()
}

func frame(sleep time.Duration) chan<- struct{} {
	resetCh := make(chan struct{}, 1)

	started := time.Now()
	go func() {
		defer close(resetCh)

		for {
			select {
			default:
			case <-resetCh:
				started = time.Now()
				continue
			}

			elapsed := time.Since(started)
			duration, _ := duration.Get()
			_ = elapsedPercents.Set(float64(elapsed) / duration)

			dmillis := time.Duration(time.Duration(duration).Milliseconds()/100*100) * time.Millisecond
			if dmillis < elapsed {
				_ = elapsedString.Set(fmt.Sprintf("%v", dmillis))
			} else {
				_ = elapsedString.Set(fmt.Sprintf("%v", time.Duration(elapsed.Milliseconds()/100*100)*time.Millisecond))
			}

			time.Sleep(sleep)
		}
	}()

	return resetCh
}
