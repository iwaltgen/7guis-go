package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/iwaltgen/7guis-go/domain/counter"
)

// Counter implements counter widget.
type Counter struct {
	value *counter.Counter
	label *widget.Label
}

// Render creates fyne.CanvasObject.
func (c *Counter) Render() fyne.CanvasObject {
	return container.NewGridWithColumns(2,
		c.label,
		widget.NewButton("Count", func() {
			c.value.Increase()
			c.label.SetText(c.value.String())
		}),
	)
}

func newCounter(iv int) *Counter {
	value := counter.Counter(iv)
	return &Counter{
		value: &value,
		label: widget.NewLabel(value.String()),
	}
}
