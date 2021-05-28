package main

import (
	"github.com/pokekrishna/weatherorama/internal/app"
	"github.com/pokekrishna/weatherorama/internal/data"
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
	_ = app.NewCurrentConditionsDisplay(measurementChangeObserver)
	data.MeasurementChanged()
}