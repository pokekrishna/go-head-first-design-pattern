package main

import (
	_ "github.com/pokekrishna/observenotify/pkg/observer"
	"time"
)

type TemperatureInC float64
type HumidityInRH float64
type PressureInPa float64

// TODO: why do you need a interface?
type interface Display {

}

// main is acting like the agent that triggers the measurementChanged()
// and mimics the weather station
func main() {
	measurementChanged()
}

// measurementChanged is to be called whenever the measurement reading from the station changes
// The custom code starts from here.
func measurementChanged(){

}

func getTemperature() (time.Time, TemperatureInC, error){

}
func getHumidity() (time.Time, HumidityInRH, error) {

}
func getPressure() (time.Time, PressureInPa, error) {

}