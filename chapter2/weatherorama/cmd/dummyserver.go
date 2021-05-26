package main

import "github.com/pokekrishna/weatherorama/internal/data"

// main is acting like the agent that triggers the measurementChanged()
// and mimics the weather station
func main() {
	data.MeasurementChanged()
}