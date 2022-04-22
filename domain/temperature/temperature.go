package temperature

import (
	"math"
)

// Temperature implements convert Celsius to Fahrenheit.
type Temperature struct {
	Celsius    int64
	Fahrenheit int64
}

// SetCelsius sets Celsius value and calculate Fahrenheit.
func (v *Temperature) SetCelsius(value int64) {
	v.Celsius = value
	v.Fahrenheit = int64(math.Round(float64(value)*(9.0/5.0) + 32))
}

// SetFahrenheit sets Fahrenheit value and calculate Celsius.
func (v *Temperature) SetFahrenheit(value int64) {
	v.Fahrenheit = value
	v.Celsius = int64(math.Round((float64(value) - 32) * (5.0 / 9.0)))
}
