// display package provides ease to implement displays.
package display

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

// genericDisplay exists because we aim to build a marketplace of
// "Displays" in future, so it acts as a contract of calling an object "Display"
//
// genericDisplay should implement Listener so that the observer package can
// notify() the display whenever the registered event occurs.
type genericDisplay struct{
	// Attributes common across all displays go here
	// ...

}

// TODO: Design limitation: A display can listen to multiple changes ...
// TODO: ...via AddToListen()s, but the Notify() is common across all ...
// TODO: ...observations. What if display wants to act differently ...
// TODO: ... based on different type of observations?
func (g *genericDisplay) Notify(time time.Time, data interface{}) error{
	// TODO: missing implementation
	fmt.Printf("notify called on %T, at %v, with data: %v\n", *g, time, data)
	return nil
}

func (g *genericDisplay) AddToListen (o *observer.Observer) {
	o.Listeners = append(o.Listeners, g)
}

func NewGenericDisplay() observer.Listener {
	return &genericDisplay{}
}


