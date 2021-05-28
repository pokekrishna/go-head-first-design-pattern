// This package needs to provide an interface that would let listeners to register
// a callback function which would be invoked when the desired change is observed.
package observer

import "time"

type Listener interface {
	Notify(time time.Time, data interface{}) error
	AddToListen()
}

type Observer struct {
	Listeners []Listener
}

func (o *Observer) NotifyAllListeners(data interface{}) error {
	receiveTimeStamp := time.Now().UTC()
	for _, listener := range o.Listeners{
		if err := listener.Notify(receiveTimeStamp, data); err != nil{
			return err
		}
	}
	return nil
}
