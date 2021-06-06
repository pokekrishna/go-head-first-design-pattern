package main

import "fmt"

type Dollar float64

type Beverage interface {
	Description() string
	Cost() Dollar
}

// TODO: `price` and `cost` are two words used as same meaning...
// TODO: ...Causing confusion

// DarkRoast is an Beverage
type DarkRoast struct {
	basePrice   Dollar
	description string
}

func NewDarkRoast() Beverage {
	return &DarkRoast{
		basePrice:   99.99,
		description: "The Most Excellent Dark Roast on the Planet!",
	}
}

func (d *DarkRoast) Description() string {
	return d.description
}

func (d *DarkRoast) Cost() Dollar {
	return d.basePrice
}

// Mocha is a decorator. Satisfies Beverage interface for Mirroring.
type Mocha struct {
	basePrice   Dollar
	description string
	beverage Beverage
}

func (m *Mocha) Description() string {
	return fmt.Sprintf("%s %s", m.beverage.Description(), m.description)
}

func (m *Mocha) Cost() Dollar {
	return m.beverage.Cost() + m.basePrice
}

// TODO: does not solve for 'with double mocha' as 'One' is hardcoded...
// TODO: ... in cost and description.
func AddMocha(b Beverage) Beverage {
	return &Mocha{
		basePrice:   7.99,
		description: "Charged with One Mocha.",
		beverage: b,
	}
}

// Whip is a decorator. Satisfies Beverage interface for Mirroring.
type Whip struct {
	basePrice   Dollar
	description string
	beverage Beverage
}

func (w *Whip) Description() string {
	return fmt.Sprintf("%s %s", w.beverage.Description(), w.description)
}

func (w *Whip) Cost() Dollar {
	return w.beverage.Cost() + w.basePrice
}

func AddWhip(b Beverage) Beverage {
	return &Whip{
		basePrice:   7.99,
		description: "Whipped, not whooped.",
		beverage: b,
	}
}

func main() {
	// trying wrapping / decorating
	dr := NewDarkRoast()
	fmt.Printf("Base Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding Mocha
	dr = AddMocha(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding whip
	dr = AddWhip(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())
}
