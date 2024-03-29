package main

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/iwaltgen/7guis-go/domain/counter"
)

// Counter implements counter widget.
type Counter struct {
	value  *counter.Counter
	button *widget.Clickable
}

/// UpdateFrame implements update frame logic.
func (c *Counter) UpdateFrame(gtx layout.Context) error {
	if c.button.Clicked() {
		c.value.Increase()
	}
	return nil
}

/// UpdateEvent implements update event logic.
func (c *Counter) UpdateEvent(evt interface{}) error {
	return nil
}

// Render implements draw widget.
func (c *Counter) Render(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, material.Body1(theme, c.value.String()).Layout)
		}),
		layout.Flexed(1, material.Button(theme, c.button, "Count").Layout),
	)
}

func newCounter(iv int) *Counter {
	value := counter.Counter(iv)
	return &Counter{
		value:  &value,
		button: &widget.Clickable{},
	}
}
