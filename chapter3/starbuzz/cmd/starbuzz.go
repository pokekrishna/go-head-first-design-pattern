package main

import (
	"fmt"
	bvg "github.com/pokekrishna/starbuzz/internal/beverage"
)

func main() {
	dr := bvg.NewDarkRoast()
	dr.SetRoastLevel(bvg.LevelMediumRoast)

	var drb bvg.Beverage
	drb = dr // Extract out the interface compliant part of 'dr'
	fmt.Printf("Base Cost of your order:\n%f : %s\n", drb.Cost(), drb.Description())

	// adding Mocha
	drb = bvg.AddMocha(drb)
	fmt.Printf("New Cost of your order:\n%f : %s\n", drb.Cost(), drb.Description())

	// adding another Mocha
	drb = bvg.AddMocha(drb)
	fmt.Printf("New Cost of your order:\n%f : %s\n", drb.Cost(), drb.Description())

	// adding whip
	drb = bvg.AddWhip(drb)
	fmt.Printf("New Cost of your order:\n%f : %s\n", drb.Cost(), drb.Description())
}
