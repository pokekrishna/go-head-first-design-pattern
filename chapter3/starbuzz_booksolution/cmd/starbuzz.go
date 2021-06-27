package main

import (
	"fmt"
	"github.com/pokekrishna/starbuzz_booksolution/internal/beverage"
	"github.com/pokekrishna/starbuzz_booksolution/internal/condiment"
)

func main() {
	var espB beverage.Beverage = new(beverage.Espresso)
	fmt.Printf("%s\t$%0.2f\n", espB.Description(), espB.Cost())

	var darkB beverage.Beverage = new(beverage.DarkRoast)
	darkB = condiment.NewMocha(darkB)
	darkB = condiment.NewMocha(darkB)
	darkB = condiment.NewWhip(darkB)
	fmt.Printf("%s\t$%0.2f\n", darkB.Description(), darkB.Cost())
}
