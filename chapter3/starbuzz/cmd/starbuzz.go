package main

import (
	"fmt"
	"github.com/pokekrishna/starbuzz/internal/beverage"
)

func main() {
	// trying wrapping / decorating
	dr := beverage.NewDarkRoast()
	fmt.Printf("Base Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding Mocha
	dr = beverage.AddMocha(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding another Mocha
	dr = beverage.AddMocha(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())

	// adding whip
	dr = beverage.AddWhip(dr)
	fmt.Printf("New Cost of your order:\n%f : %s\n", dr.Cost(), dr.Description())
}
