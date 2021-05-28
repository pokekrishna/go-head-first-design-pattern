package data

import (
	"fmt"
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
	err := measurementChangeObserver.NotifyAllListeners("measurement changed!")
	if err != nil {
		fmt.Println("error occurred in notifying listeners.", err)
	}
}

func SetMeasurementChangeObserver(o *observer.Observer) {
	measurementChangeObserver = o
}

func GetTemperature() (time.Time, TemperatureInC, error){
	return time.Time{}, *(new(TemperatureInC)), nil
}
func GetHumidity() (time.Time, HumidityInRH, error) {
	return time.Time{}, *(new(HumidityInRH)), nil
}
func GetPressure() (time.Time, PressureInPa, error) {
	return time.Time{}, *(new(PressureInPa)), nil
}