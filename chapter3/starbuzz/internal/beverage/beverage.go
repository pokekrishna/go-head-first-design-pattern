package beverage

type Dollar float64

type Beverage interface {
	Description() string
	Cost() Dollar
}

// DarkRoast implements Beverage
type DarkRoast struct {
	baseCost    Dollar
	description string
}

func NewDarkRoast() Beverage {
	return &DarkRoast{
		baseCost:    99.99,
		description: "The Most Excellent Dark Roast on the Planet!",
	}
}

func (d *DarkRoast) Description() string {
	return d.description
}

func (d *DarkRoast) Cost() Dollar {
	return d.baseCost
}

