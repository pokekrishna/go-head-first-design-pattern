package app

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/internal/data"
	"github.com/pokekrishna/weatherorama/pkg/display"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

type measurementFrame struct {
	timestamp time.Time
	temp data.TemperatureInC
	pres data.PressureInPa
	humi data.HumidityInRH
}

type statisticsDisplay struct {
	measurementStore []*measurementFrame
}

func NewStatisticsDisplay(measurementChangeObserver *observer.Subject) *display.GenericDisplay{
	mainBehavior := display.NewBehavior(measurementChangeObserver, &statisticsDisplay{})

	statisticsDisplay := &display.GenericDisplay{}
	statisticsDisplay.AddNewBehavior(mainBehavior)
	return statisticsDisplay
}

func (s *statisticsDisplay) Callback(timestamp time.Time, callData interface{}) error{
	fmt.Printf("Statistics Display Notified at %v with data '%v'\n", timestamp, callData)
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

	m := &measurementFrame{
		timestamp: timestamp,
		temp:      t,
		pres:      p,
		humi:      h,
	}
	s.measurementStore = append(s.measurementStore, m)
	s.ShowCompleteHistory()
	return nil
}

func (s *statisticsDisplay) ShowCompleteHistory() {
	for _, mf := range s.measurementStore{
		fmt.Printf("Measurement at '%v': %v, %v, %v\n", mf.timestamp, mf.temp, mf.pres, mf.humi)
	}
}