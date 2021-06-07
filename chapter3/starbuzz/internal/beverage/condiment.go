package beverage

import "fmt"

// Mocha is a decorator. Satisfies Beverage interface for Mirroring.
type Mocha struct {
	baseCost    Dollar
	description string
	beverage    Beverage
}

func (m *Mocha) Description() string {
	return fmt.Sprintf("%s %s", m.beverage.Description(), m.description)
}

func (m *Mocha) Cost() Dollar {
	return m.beverage.Cost() + m.baseCost
}

// TODO: does not solve for 'with double mocha' as 'One' is hardcoded...
// TODO: ... in cost and description.
func AddMocha(b Beverage) Beverage {
	return &Mocha{
		baseCost:    7.99,
		description: "Charged with One Mocha.",
		beverage:    b,
	}
}

// Whip is a decorator. Satisfies Beverage interface for Mirroring.
type Whip struct {
	baseCost    Dollar
	description string
	beverage    Beverage
}

func (w *Whip) Description() string {
	return fmt.Sprintf("%s %s", w.beverage.Description(), w.description)
}

func (w *Whip) Cost() Dollar {
	return w.beverage.Cost() + w.baseCost
}

func AddWhip(b Beverage) Beverage {
	return &Whip{
		baseCost:    7.99,
		description: "Whipped, not whooped.",
		beverage:    b,
	}
}
