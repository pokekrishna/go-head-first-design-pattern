package display

// DisplayElement is an interface which acts as a contract for all kinds of
// displays. Remember, we should expect developers to bring in their own displays
// through `marketplace.`
type DisplayElement interface {
	Display()
}

