package data

import (
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

var (
	// TODO: Do away with global var ASAP
	measurementChangeObserver *observer.Observer
)

type TemperatureInC float64
type HumidityInRH float64
type PressureInPa float64

// MeasurementChanged is called whenever the measurement reading from the station
// changes The custom code starts from here.
func MeasurementChanged(){
	measurementChangeObserver.NotifyAllListeners("measurement changed!")
}

func SetMeasurementChangeObserver(o *observer.Observer) {
	measurementChangeObserver = o
}

func getTemperature() (time.Time, TemperatureInC, error){
	return time.Time{}, *(new(TemperatureInC)), nil
}
func getHumidity() (time.Time, HumidityInRH, error) {
	return time.Time{}, *(new(HumidityInRH)), nil
}
func getPressure() (time.Time, PressureInPa, error) {
	return time.Time{}, *(new(PressureInPa)), nil
}