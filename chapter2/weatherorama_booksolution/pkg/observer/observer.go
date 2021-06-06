package observer

// Subject not only caters to notifying all the observers it has, but also
// registers and deregisters (`RemoveObserver`) observers.
type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)

	NotifyObservers()
}

// Observer has only responsibility as contract, that it should have an method
// (`Update`) which Subject can callback when it deems to `Notify` the observer.
//
// Update does not take any parameter because, it has been said repeatedly in the
// book that ideally Update should not be prescriptive about what data should
// the observer need when notified. Instead Observers are supposed to pull data
// from the Subject as and when required.
type Observer interface {
	Update()
}
