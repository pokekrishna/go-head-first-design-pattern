package app

import (
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
)

func NewCurrentConditionsDisplay(measurementChangeObserver *observer.Observer) *display.GenericDisplay{
	mainBehavior := &display.DisplayBehavior{Observer: measurementChangeObserver}
	mainBehavior.AddToListen()
	currentConditionsDisplay := &display.GenericDisplay{}
	currentConditionsDisplay.AddNewBehavior(mainBehavior)
	return currentConditionsDisplay
}