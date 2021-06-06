package main

import (
	"github.com/pokekrishna/weatherorama_booksolution/internal/app"
	"github.com/pokekrishna/weatherorama_booksolution/internal/data"
)

func main() {
	// Create a new subject
	weatherData := data.NewWeatherData()
	_ = app.NewCurrentConditionsDisplay(weatherData)
	// ... other displays instantiation creation here

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.9)
}
