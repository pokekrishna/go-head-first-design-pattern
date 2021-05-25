// This package needs to provide an interface that would let listeners to register
// a callback function which would be invoked when the a desired change is observed.
package observer

import "time"

type Notifier interface {
	Notify(time time.Time)
}

type Observer interface {
	Register()
}




