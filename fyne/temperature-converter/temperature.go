package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"

	"github.com/iwaltgen/7guis-go/domain/temperature"
)

// Temperature implements temperature converter widget.
type Temperature struct {
	temperature       *temperature.Temperature
	labelCelsius      *widget.Label
	labelFahrenheit   *widget.Label
	bindingCelsius    binding.String
	bindingFahrenheit binding.String
	inputCelsius      *widget.Entry
	inputFahrenheit   *widget.Entry
}

// Render creates fyne.CanvasObject.
func (t *Temperature) Render() fyne.CanvasObject {
	t.inputCelsius.OnChanged = func(text string) {
		celsius, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return
		}

		t.temperature.SetCelsius(celsius)
		_ = t.bindingFahrenheit.Set(strconv.FormatInt(t.temperature.Fahrenheit, 10))
	}

	t.inputFahrenheit.OnChanged = func(text string) {
		fahrenheit, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return
		}

		t.temperature.SetFahrenheit(fahrenheit)
		_ = t.bindingCelsius.Set(strconv.FormatInt(t.temperature.Celsius, 10))
	}

	return container.NewGridWithColumns(4,
		t.inputCelsius,
		t.labelCelsius,
		t.inputFahrenheit,
		t.labelFahrenheit,
	)
}

func newTemperature() *Temperature {
	temperature := &temperature.Temperature{
		Celsius:    0,
		Fahrenheit: 32,
	}

	bindingCelsius := binding.NewString()
	bindingFahrenheit := binding.NewString()

	_ = bindingCelsius.Set(strconv.FormatInt(temperature.Celsius, 10))
	_ = bindingFahrenheit.Set(strconv.FormatInt(temperature.Fahrenheit, 10))

	return &Temperature{
		temperature:       temperature,
		labelCelsius:      widget.NewLabel("Celsius ="),
		labelFahrenheit:   widget.NewLabel("Fahrenheit"),
		bindingCelsius:    bindingCelsius,
		bindingFahrenheit: bindingFahrenheit,
		inputCelsius:      widget.NewEntryWithData(bindingCelsius),
		inputFahrenheit:   widget.NewEntryWithData(bindingFahrenheit),
	}
}
