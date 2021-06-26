package beverage

type Dollar float64

type Beverage interface {
	Description() string
	Cost() Dollar
}

const (
	_               = iota
	LevelLightRoast = iota * (100/3)
	LevelMediumRoast
	LevelHeavyRoast

	LevelDefaultRoast = LevelMediumRoast
)

// DarkRoast implements Beverage
type DarkRoast struct {
	baseCost    Dollar
	description string
	roastPercentage int
}

// NewDarkRoast returns pointer reference to DarkRoast and not the interface type
// so that the caller can access the fields of the concrete type and
// methods not defined on the interface can be accessed, if required.
// Later the caller can convert it to interface type.
//
// Example Usage:
//	dr := NewDarkRoast()
//	level := dr.roastPercentage
//  var Beverage b = dr
//	b = AddMocha(b)
func NewDarkRoast() *DarkRoast {
	return &DarkRoast{
		baseCost:        99.99,
		description:     "The Most Excellent Dark Roast on the Planet!",
		roastPercentage: LevelDefaultRoast,
	}
}

func (d *DarkRoast) Description() string {
	return d.description
}

func (d *DarkRoast) Cost() Dollar {
	return d.baseCost
}

func (d *DarkRoast) SetRoastLevel(roastLevel int) {
	d.roastPercentage = roastLevel
}
