package data

import (
	"time"
)

type TemperatureInC float64
type HumidityInRH float64
type PressureInPa float64

// MeasurementChanged is called whenever the measurement reading from the station
// changes The custom code starts from here.
func MeasurementChanged(){
	// TODO: All the interested displays should be notified here.
	// notified through the API, not individually because tomorrow
	// there can be new displays we may not be aware
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