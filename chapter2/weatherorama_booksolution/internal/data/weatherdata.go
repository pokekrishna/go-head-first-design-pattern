package data

import (
	"fmt"
	"github.com/pokekrishna/weatherorama_booksolution/pkg/observer"
)

type TemperatureInC float64

func (t TemperatureInC) String() string{
	return fmt.Sprintf("%v°C", float64(t))
}

type HumidityInRH float64

func (h HumidityInRH) String() string {
	return fmt.Sprintf("%vRH", float64(h))
}

type PressureInPa float64

func (p PressureInPa) String() string {
	return fmt.Sprintf("%vPa", float64(p))
}

// WeatherData implements the Subject interface. In other words,
// WeatherData is a Subject.
type WeatherData struct {
	observers []observer.Observer
	temp      TemperatureInC
	humid     HumidityInRH
	press     PressureInPa
}

func (w *WeatherData) RegisterObserver(o observer.Observer) {
	w.observers = append(w.observers, o)
}

// Go does not provide a method to remove element from a slice.
// This implementation of removing observer from the slice is not
// at all performance optimized or fail proof.
func (w *WeatherData) RemoveObserver(o observer.Observer) {
	for i, observer := range w.observers {
		if observer == o {
			w.observers[i] = nil
		}
	}
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update()
	}
}

func (w *WeatherData) MeasurementChanged(){
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(t TemperatureInC, h HumidityInRH, p PressureInPa){
	w.temp = t
	w.humid = h
	w.press = p
	w.MeasurementChanged()
}

func (w *WeatherData) GetTemperature() TemperatureInC {
	return w.temp
}

func (w *WeatherData) GetHumidity() HumidityInRH {
	return w.humid
}

// POSSIBLE DESIGN FLAW: Returning a WeatherData reference and not an observer.Subject
// Type WeatherData is required because the reference is passed to access the methods
// only available on WeatherData and not on observer.Subject
func NewWeatherData() *WeatherData {
	return &WeatherData{}
}
