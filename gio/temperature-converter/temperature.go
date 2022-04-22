package main

import (
	"strconv"

	"gioui.org/io/key"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/iwaltgen/7guis-go/domain/temperature"
)

// Temperature implements temperature converter widget.
type Temperature struct {
	temperature      *temperature.Temperature
	editorCelsius    *widget.Editor
	editorFahrenheit *widget.Editor
}

/// UpdateFrame implements update frame logic.
func (t *Temperature) UpdateFrame(gtx layout.Context) error {
	for _, e := range t.editorCelsius.Events() {
		if _, ok := e.(widget.ChangeEvent); ok {
			text := t.editorCelsius.Text()
			celsius, err := strconv.ParseInt(text, 10, 64)
			if err != nil || t.temperature.Celsius == celsius {
				continue
			}

			t.temperature.SetCelsius(celsius)
			t.editorFahrenheit.SetText(strconv.FormatInt(t.temperature.Fahrenheit, 10))
		}
	}

	for _, e := range t.editorFahrenheit.Events() {
		if _, ok := e.(widget.ChangeEvent); ok {
			text := t.editorFahrenheit.Text()
			fahrenheit, err := strconv.ParseInt(text, 10, 64)
			if err != nil || t.temperature.Fahrenheit == fahrenheit {
				continue
			}

			t.temperature.SetFahrenheit(fahrenheit)
			t.editorCelsius.SetText(strconv.FormatInt(t.temperature.Celsius, 10))
		}
	}
	return nil
}

/// UpdateEvent implements update event logic.
func (t *Temperature) UpdateEvent(evt interface{}) error {
	switch e := evt.(type) {
	case key.Event:
		if e.Name == key.NameTab {
			if t.editorCelsius.Focused() {
				t.editorFahrenheit.Focus()
			} else {
				t.editorCelsius.Focus()
			}
		}

	default:
		// log.Printf("window event: %#v\n", e)
	}
	return nil
}

// Render implements draw widget.
func (t *Temperature) Render(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	inset := unit.Dp(6)
	return layout.Flex{}.Layout(gtx,
		layout.Flexed(1, t.uniformInset(inset, material.Editor(theme, t.editorCelsius, "Celsius").Layout)),
		layout.Rigid(t.uniformInset(inset, material.Body1(theme, "Celsius = ").Layout)),
		layout.Flexed(1, t.uniformInset(inset, material.Editor(theme, t.editorFahrenheit, "Fahrenheit").Layout)),
		layout.Rigid(t.uniformInset(inset, material.Body1(theme, "Fahrenheit").Layout)),
	)
}

func (t *Temperature) uniformInset(inset unit.Value, widget layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(inset).Layout(gtx, widget)
	}
}

func newTemperature(iv int) *Temperature {
	temperature := &temperature.Temperature{
		Celsius:    0,
		Fahrenheit: 32,
	}
	editorCelsius := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}
	editorFahrenheit := &widget.Editor{
		SingleLine: true,
		Submit:     true,
	}

	editorCelsius.Focus()
	editorCelsius.SetText(strconv.FormatInt(temperature.Celsius, 10))
	editorFahrenheit.SetText(strconv.FormatInt(temperature.Fahrenheit, 10))
	return &Temperature{
		temperature:      temperature,
		editorCelsius:    editorCelsius,
		editorFahrenheit: editorFahrenheit,
	}
}
