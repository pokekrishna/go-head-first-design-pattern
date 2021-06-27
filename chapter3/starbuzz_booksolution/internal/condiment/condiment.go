package condiment

import (
	"fmt"
	"github.com/pokekrishna/starbuzz_booksolution/internal/beverage"
	i18n "github.com/pokekrishna/starbuzz_booksolution/internal/internationalization"
)

type Decorator interface {
	beverage.Beverage
	Description() string
}
type Mocha struct {
	B beverage.Beverage
}

func NewMocha(b beverage.Beverage) *Mocha {
	return &Mocha{B: b}
}

func (m *Mocha) Description() string {
	return fmt.Sprintf("%s, Mocha", m.B.Description())
}

func (m *Mocha) Cost() i18n.Dollar {
	return m.B.Cost() + 0.20
}

type Whip struct {
	B beverage.Beverage
}

func NewWhip(b beverage.Beverage) *Whip {
	return &Whip{B: b}
}

func (w *Whip) Description() string {
	return fmt.Sprintf("%s, Whip", w.B.Description())
}

func (w *Whip) Cost() i18n.Dollar {
	return w.B.Cost() + 0.39
}
