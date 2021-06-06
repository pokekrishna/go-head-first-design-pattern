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

type Condiment interface {
	AddToBeverage(Beverage) Beverage
}

// Mocha is a Condiment but satisfies the Beverage interface too
type Mocha struct {
	basePrice   Dollar
	description string

	newDescription string
	newPrice Dollar
}

// TODO: does not solve for 'with double mocha' as 'One' is hardcoded...
// TODO: ... in cost and description.
func (m *Mocha) AddToBeverage(beverage Beverage) Beverage {
	// change the price
	// change description
	m.newDescription = fmt.Sprintf("%s %s", beverage.Description(), m.description)
	m.newPrice = beverage.Cost() + m.basePrice
	return m
}

func (m *Mocha) Description() string {
	return m.newDescription
}

func (m *Mocha) Cost() Dollar {
	return m.newPrice
}

func NewMocha() Condiment {
	return &Mocha{
		basePrice:   7.99,
		description: "Charged with One Mocha.",
	}
}

// Whip is a Condiment but satisfies the Beverage interface too
type Whip struct {
	basePrice   Dollar
	description string

	newDescription string
	newPrice Dollar
}

func (w *Whip) Description() string {
	return w.newDescription
}

func (w *Whip) Cost() Dollar {
	return w.newPrice
}

func (w *Whip) AddToBeverage(beverage Beverage) Beverage {
	w.newDescription = fmt.Sprintf("%s %s", beverage.Description(), w.description)
	w.newPrice = beverage.Cost() + w.basePrice
	return w
}

func NewWhip() Condiment {
	return &Whip{
		basePrice:   7.99,
		description: "Whipped, not whooped.",
	}
}

func main() {
	// trying wrapping / decorating
	dr := NewDarkRoast()
	fmt.Printf("Base Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding Mocha
	m := NewMocha()
	dr = m.AddToBeverage(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding whip
	w := NewWhip()
	dr = w.AddToBeverage(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())
}
