package app

import (
	"fmt"
	"github.com/pokekrishna/weatherorama_booksolution/internal/data"
)

// CurrentConditionsDisplay implements both DisplayElement and Observer
type CurrentConditionsDisplay struct {
	currentTemp data.TemperatureInC
	currentHumid data.HumidityInRH
	
	// The subject to subscribe to. IMPORTANT thing to note here is that weatherData
	// is not of type observer.Subject but of type data.WeatherData.
	// This is required because we are calling methods on WeatherData which are not
	// present on the interface. For instance, GetTemperature.
	//
	// Would this make testing hard?
	weatherData *data.WeatherData
}

func (c *CurrentConditionsDisplay) Display() {
	fmt.Println("Current Display Notified!")
	fmt.Printf("Current measurement: %v, %v\n", c.currentTemp, c.currentHumid)
}

func (c *CurrentConditionsDisplay) Update() {
	c.currentTemp = c.weatherData.GetTemperature()
	c.currentHumid = c.weatherData.GetHumidity()
	c.Display()
}

func NewCurrentConditionsDisplay(w *data.WeatherData) *CurrentConditionsDisplay {
	currentConditionsD := &CurrentConditionsDisplay{
		weatherData: w,
	}
	w.RegisterObserver(currentConditionsD)
	return currentConditionsD
}