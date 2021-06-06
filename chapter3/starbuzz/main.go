package main

import "fmt"

type Dollar float64

type Beverage interface {
	Description() string
	Cost() Dollar
}

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

// Mocha is a Condiment and also a Beverage
type Mocha struct {
	basePrice   Dollar
	description string

	newPrice Dollar
}

// TODO: does not solve for 'with double mocha' as 'One' is hardcoded...
// TODO: ... in cost and description.
func (m *Mocha) AddToBeverage(beverage Beverage) Beverage {
	// change the price
	// change description
	m.basePrice = 7.99
	m.description = "Charged with One Mocha."
	m.newPrice = beverage.Cost() + m.basePrice
	return m
}

func (m *Mocha) Description() string {
	return m.description
}

func (m *Mocha) Cost() Dollar {
	return m.newPrice
}

func main() {
	// trying wrapping / decorating
	dr := NewDarkRoast()
	fmt.Printf("Base Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())
	// adding Mocha
	m := Mocha{}
	dr = m.AddToBeverage(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())
}
