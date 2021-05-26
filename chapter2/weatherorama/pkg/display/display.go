// display package provides ease to implement displays.
package display

import "github.com/pokekrishna/weatherorama/pkg/observer"

// GenericDisplay interface exists because we aim to build a marketplace of
// "Displays" in future, so it acts as a contract of calling an object "Display"
//
// All displays should implement Notifier so that the observer package can
// notify() the display whenever the registered event occurs.
type GenericDisplay interface {
	observer.Notifier
}

