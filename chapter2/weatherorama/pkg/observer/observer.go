// This package needs to provide an interface that would let Observers to register
// a callback function which would be invoked when the desired change is observed
// by the Subject.
package observer

import "time"

type Observer interface {
	Notify(time time.Time, data interface{}) error
	AddToListen()
}

type Subject struct {
	Observers []Observer
}

func (s *Subject) NotifyAllObservers(data interface{}) error {
	receiveTimeStamp := time.Now().UTC()
	for _, listener := range s.Observers {
		if err := listener.Notify(receiveTimeStamp, data); err != nil{
			return err
		}
	}
	return nil
}
