package app

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"github.com/pokekrishna/weatherorama/internal/data"
	"time"
)

type currentConditionsDisplay struct {}

func NewCurrentConditionsDisplay(measurementChangeObserver *observer.Observer) *display.GenericDisplay{
	mainBehavior := display.NewBehavior(measurementChangeObserver, &currentConditionsDisplay{})

	currentConditionsDisplay := &display.GenericDisplay{}
	currentConditionsDisplay.AddNewBehavior(mainBehavior)
	return currentConditionsDisplay
}


func (c *currentConditionsDisplay) Callback(time time.Time, callData interface{}) error{
	fmt.Printf("Current Display Notified at %v with data '%v'\n", time, callData)
	_, t, err := data.GetTemperature()
	if err != nil {
		return err
	}
	fmt.Printf("Current temp: %v\n", t)

	_, h, err := data.GetHumidity()
	if err != nil {
		return err
	}
	fmt.Printf("Current humidity: %v\n", h)

	_, p, err := data.GetPressure()
	if err != nil {
		return err
	}
	fmt.Printf("Current temp: %v\n", p)

	return nil
}
