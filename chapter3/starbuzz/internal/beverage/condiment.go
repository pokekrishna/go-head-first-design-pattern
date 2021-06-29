package beverage

import (
	"fmt"
	"strings"
)

// Mocha is a decorator. Implements Beverage interface for Mirroring.
type Mocha struct {
	baseCost    Dollar
	description string
	beverage    Beverage
}

func (m *Mocha) Description() string {
	if !strings.Contains(m.beverage.Description(), m.description){
		return fmt.Sprintf("%s %s", m.beverage.Description(), m.description)
	}
	return m.beverage.Description()
}

func (m *Mocha) Cost() Dollar {
	return m.beverage.Cost() + m.baseCost
}

func AddMocha(b Beverage) *Mocha {
	return &Mocha{
		baseCost:    7.99,
		description: "Charged with Mocha.",
		beverage:    b,
	}
}

// Whip is a decorator. Implements Beverage interface for Mirroring.
type Whip struct {
	baseCost    Dollar
	description string
	beverage    Beverage
}

func (w *Whip) Description() string {
	if !strings.Contains(w.beverage.Description(), w.description) {
		return fmt.Sprintf("%s %s", w.beverage.Description(), w.description)
	}
	return w.beverage.Description()
}

func (w *Whip) Cost() Dollar {
	return w.beverage.Cost() + w.baseCost
}

func AddWhip(b Beverage) *Whip {
	return &Whip{
		baseCost:    7.99,
		description: "Whipped, not whooped.",
		beverage:    b,
	}
}
