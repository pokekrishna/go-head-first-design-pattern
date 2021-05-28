// display package provides ease to implement displays.
package display

import (
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

// ListenerCallback gives a call to make when Notify has been called.
// This is an interface and not just a function or a type on function
// because if the display wants to operate on the object variables, it
// can do so by accessing the receiver parameter of the dynamic type of
// the interface.
type ListenerCallback interface {
	Callback(time time.Time, data interface{}) error
}

// TODO: Is it necessary that all display behaviors will need to be notified...
// TODO: ...of some event or the other? May wanna change the Name of the type?
type behavior struct {
	Observer *observer.Observer
	lc ListenerCallback
}

func (behavior *behavior) Notify(time time.Time, data interface{}) error{
	if err := behavior.lc.Callback(time, data); err != nil{
		return err
	}
	return nil
}

func (behavior *behavior) AddToListen () {
	behavior.Observer.Listeners = append(behavior.Observer.Listeners, behavior)
}

func NewBehavior(o *observer.Observer, lc ListenerCallback) observer.Listener{
	b := &behavior{
		Observer: o,
		lc: lc,
	}
	b.AddToListen()
	return b
}

// GenericDisplay exists because we aim to build a marketplace of
// "Displays" in future, so it acts as a contract of calling an object "Display"
type GenericDisplay struct{
	// Attributes common across all displays go here
	// ...
	// TODO: Design limitation: With a slice, adding new behaviors is easy but...
	// TODO: ...'get'-ing them is cumbersome - How do you know which is which?
	behaviors []observer.Listener
}

func (g *GenericDisplay) AddNewBehavior(behavior observer.Listener){
	g.behaviors = append(g.behaviors, behavior)
}



