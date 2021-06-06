package app

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

type forecastDisplay struct {}

func NewForecastDisplay(measurementChangeObserver *observer.Subject) *display.GenericDisplay{
	mainBehavior := display.NewBehavior(measurementChangeObserver, &forecastDisplay{})

	forecastD := &display.GenericDisplay{}
	forecastD.AddNewBehavior(mainBehavior)
	return forecastD
}

func (f *forecastDisplay) Callback(timestamp time.Time, callData interface{}) error{
	fmt.Printf("Forecast Display Notified at %v with data '%v'\n", timestamp, callData)
	return nil
}

