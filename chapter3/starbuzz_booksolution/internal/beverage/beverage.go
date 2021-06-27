package beverage

import i18n "github.com/pokekrishna/starbuzz_booksolution/internal/internationalization"

type Beverage interface{
	// Interfaces in Go do not support variables (data), therefore Getter like
	// getDescription() as suggested in the book does not make sense.
	//
	// Also, formal design of the Decorator Pattern does not dictate to
	// have variables (data) but does dictate to have behavior(s) as contract.
	Description() string
	Cost() i18n.Dollar
}

// TODO: Do these concrete types need to have struct fields? ...
// TODO... The way book's solution has it?

// Espresso implements Beverage
type Espresso struct {}

func (e *Espresso) Description() string {
	return "Espresso"
}

func (e *Espresso) Cost() i18n.Dollar{
	return 1.99
}


// DarkRoast implements Beverage
type DarkRoast struct {}

func (d *DarkRoast) Description() string {
	return "DarkRoast"
}

func (d *DarkRoast) Cost() i18n.Dollar{
	return 2.99
}