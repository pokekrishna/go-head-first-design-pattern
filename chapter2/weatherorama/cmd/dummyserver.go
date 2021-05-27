package main

import (
	"github.com/pokekrishna/weatherorama/internal/data"
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
)

// main is acting like the agent that triggers the measurementChanged()
// and mimics the weather station
func main() {
	measurementChangeObserver := &observer.Observer{}
	data.SetMeasurementChangeObserver(measurementChangeObserver)

	// TODO: creating individual displays here! Any suggested pattern that can ...
	// TODO: ... create the displays and also cater to the need of future ...
	// TODO: ... displays through marketplace?
	currentConditionsDisplay := display.NewGenericDisplay()
	statisticsDisplay := display.NewGenericDisplay()
	foreCastDisplay := display.NewGenericDisplay()

	currentConditionsDisplay.AddToListen(measurementChangeObserver)
	statisticsDisplay.AddToListen(measurementChangeObserver)
	foreCastDisplay.AddToListen(measurementChangeObserver)

	data.MeasurementChanged()
}