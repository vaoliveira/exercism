// Package weather forecasts the weather for various cities.
package weather

// CurrentCondition holds the current condition of a city.
var CurrentCondition string

// CurrentLocation holds the current location.
var CurrentLocation string

// Forecast function that receives a city and its condition, and it returns a string with this data.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
