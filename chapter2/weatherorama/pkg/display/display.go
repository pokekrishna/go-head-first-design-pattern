// display package provides ease to implement displays.
package display

import (
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

// GenericDisplay exists because we aim to build a marketplace of
// "Displays" in future, so it acts as a contract of calling an object "Display"
type GenericDisplay struct{
	// Attributes common across all displays go here
	// ...
	// TODO: Design limitation: With a slice, adding new behaviors is easy but...
	// TODO: ...'get'-ing them is cumbersome - How do you know which is which?
	behaviors []observer.Observer
}

func (g *GenericDisplay) AddNewBehavior(behavior observer.Observer){
	g.behaviors = append(g.behaviors, behavior)
}

// behavior implements the Observer interface and it depicts that
// behavior of an display that is Notify-able, which means there can
// be multiple behaviors in a display which can be Notify-ied by their
// respective subject.
//
// When Notify is called on the behavior, oc is in turn called.
type behavior struct {
	subject *observer.Subject
	oc      ObserverCallback
}

func (behavior *behavior) Notify(time time.Time, data interface{}) error{
	if err := behavior.oc.Callback(time, data); err != nil{
		return err
	}
	return nil
}

func (behavior *behavior) AddToListen() {
	behavior.subject.Observers = append(behavior.subject.Observers, behavior)
}

// NewBehavior returns a behavior that satisfies the observer.Observer interface
// Each behavior is tied to a observer.Subject, meaning the observer.Observer
// is registered to the observer.Subject.
func NewBehavior(s *observer.Subject, oc ObserverCallback) observer.Observer {
	b := &behavior{
		subject: s,
		oc:      oc,
	}
	b.AddToListen()
	return b
}

// ObserverCallback gives a call to make when Notify has been called.
// This is an interface and not just a function or a type on function
// because if the display wants to operate on the object variables, it
// can do so by accessing the receiver parameter of the dynamic type of
// the interface.
type ObserverCallback interface {
	Callback(time time.Time, data interface{}) error
}
