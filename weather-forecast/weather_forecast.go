// Package weather provides a way to forecast the weather.
package weather

// CurrentCondition of the city.
var CurrentCondition string

// CurrentLocation we are forecasting.
var CurrentLocation string

// Forecast provides current weather information about a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
