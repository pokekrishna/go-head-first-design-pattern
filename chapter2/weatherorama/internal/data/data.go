package data

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

var (
	// TODO: Do away with global var ASAP
	measurementChangeSubject *observer.Subject
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

// MeasurementChanged is called whenever the measurement reading from the station
// changes The custom code starts from here.
func MeasurementChanged(){
	err := measurementChangeSubject.NotifyAllListeners("measurement changed!")
	if err != nil {
		fmt.Println("error occurred in notifying listeners.", err)
	}
}

func SetMeasurementChangeSubject(o *observer.Subject) {
	measurementChangeSubject = o
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