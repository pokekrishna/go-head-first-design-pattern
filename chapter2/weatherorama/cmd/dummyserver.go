package main

import (
	"github.com/pokekrishna/weatherorama/internal/app"
	"github.com/pokekrishna/weatherorama/internal/data"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

// main is acting like the agent that triggers the measurementChanged()
// and mimics the weather station
func main() {
	measurementChangeSubject := &observer.Subject{}
	data.SetMeasurementChangeSubject(measurementChangeSubject)

	// TODO: creating individual displays here! Any suggested pattern that can ...
	// TODO: ... create the displays and also cater to the need of future ...
	// TODO: ... displays through marketplace?
	_ = app.NewCurrentConditionsDisplay(measurementChangeSubject)
	_ = app.NewStatisticsDisplay(measurementChangeSubject)
	_ = app.NewForecastDisplay(measurementChangeSubject)

	data.MeasurementChanged()
	// simulating another call
	time.Sleep(2 * time.Second)
	data.MeasurementChanged()
}