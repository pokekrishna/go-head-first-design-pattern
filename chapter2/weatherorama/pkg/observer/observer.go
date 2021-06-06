// This package needs to provide an interface that would let Observers to register
// a callback function which would be invoked when the desired change is observed
// by the Subject.
package observer

import "time"

// TODO: Just like AddToListen Registers the Observer, need ...
// TODO: a DeRegister method.
type Observer interface {
	Notify(time time.Time, data interface{}) error
	AddToListen()
}

type Subject struct {
	Observers []Observer
}

func (s *Subject) NotifyAllObservers(data interface{}) error {
	notifyTimeStamp := time.Now().UTC()
	for _, observer := range s.Observers {
		if err := observer.Notify(notifyTimeStamp, data); err != nil{
			return err
		}
	}
	return nil
}
