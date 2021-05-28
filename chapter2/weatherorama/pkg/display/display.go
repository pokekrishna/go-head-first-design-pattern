// display package provides ease to implement displays.
package display

import (
	"fmt"
	"github.com/pokekrishna/weatherorama/pkg/observer"
	"time"
)

// TODO: Is it necessary that all display behaviors will need to be notified...
// TODO: ...of some event or the other? May wanna change the Name of the type?
type DisplayBehavior struct {
	Observer *observer.Observer

}

func (behavior *DisplayBehavior) Notify(time time.Time, data interface{}) error{
	// TODO: missing implementation
	fmt.Printf("notify called on %v, at %v, with '%v'\n", behavior, time, data)
	return nil
}

func (behavior *DisplayBehavior) AddToListen () {
	behavior.Observer.Listeners = append(behavior.Observer.Listeners, behavior)
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



