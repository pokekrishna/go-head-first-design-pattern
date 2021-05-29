// This package is an attempt to replicate the solution from the book
package main

import "fmt"

// Duck is an attempt to port from the Duck class mentioned in the book
//
// Problem: Why does all Duck need to have FlyBehavior? even if it never going to
// fly. Same with QuackBehavior.
type Duck struct {
	fb FlyBehaviour
	qb QuackBehaviour
	display func()
}

func (d *Duck) performFly(){
	d.fb.Fly()
}

func (d *Duck) performQuack(){
	d.fb.Fly()
}

func (d *Duck) swim(){
	fmt.Println("All ducks float, even decoys!")
}


type MallardDuck struct {
	Duck
}

// I don't think these setters are needed here, since instantiating
// the struct gives you the option to set all the fields.
func (md *MallardDuck) setFlyBehavior(fb FlyBehaviour){
	md.fb = fb
}

func (md *MallardDuck) setQuackBehavior(qb QuackBehaviour){
	md.qb = qb
}


func main() {

	var md *MallardDuck = &MallardDuck{Duck{
		fb:                     &DefaultFly{100, 99},
		qb:                     &DefaultQuack{},
		display: func() {
			fmt.Println("I am a mallard duck")
		},
	}}
	md.display()

	md.performFly()

	mfly, err := NewRocketPoweredFlying(1000, 100, 999, "ethanol")
	if err != nil {
		panic(err)
	}
	// Demonstrating runtime behavior change leveraging Polymorphism
	md.setFlyBehavior(mfly)
	md.performFly()
}
