package app

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"github.com/pokekrishna/weatherorama/internal/data"
	"time"
)

type currentDisplay struct {}

func NewCurrentConditionsDisplay(measurementChangeObserver *observer.Observer) *display.GenericDisplay{
	mainBehavior := display.NewBehavior(measurementChangeObserver, &currentDisplay{})

	currentConditionsDisplay := &display.GenericDisplay{}
	currentConditionsDisplay.AddNewBehavior(mainBehavior)
	return currentConditionsDisplay
}


func (c *currentDisplay) Callback(timestamp time.Time, callData interface{}) error{
	fmt.Printf("Current Display Notified at %v with data '%v'\n", timestamp, callData)
	_, t, err := data.GetTemperature()
	if err != nil {
		return err
	}

	_, h, err := data.GetHumidity()
	if err != nil {
		return err
	}

	_, p, err := data.GetPressure()
	if err != nil {
		return err
	}

	fmt.Printf("Current measurement: %v, %v, %v\n", t, p, h)
	return nil
}