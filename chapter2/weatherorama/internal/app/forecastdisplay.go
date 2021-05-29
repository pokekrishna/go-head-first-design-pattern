package app

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

type forecastDisplay struct {}

func NewForecastDisplay(measurementChangeObserver *observer.Observer) *display.GenericDisplay{
	mainBehavior := display.NewBehavior(measurementChangeObserver, &forecastDisplay{})

	forecastDisplay := &display.GenericDisplay{}
	forecastDisplay.AddNewBehavior(mainBehavior)
	return forecastDisplay
}

func (f *forecastDisplay) Callback(timestap time.Time, callData interface{}) error{
	fmt.Printf("Forecast Display Notified at %v with data '%v'\n", timestap, callData)
	return nil
}

