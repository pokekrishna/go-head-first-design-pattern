package main

import "fmt"

// Q: solving HF DP, SimUDuck problem

// Duck interface enforces to have Display() in all the
// implementations of Duck.
//
// Display is intended to be implemented by the dynamic type.
//
// Swim is intended to be present in all duck types however in order
// to aid code reuse it can simply just call the Action which it is
// passed.
type Duck interface {
	Display()
	Swim(action SwimAction)
}

// TODO: FLAW with design. each action is a function. which means you cannot work
// on 'state' other than the ones passed to the functions as parameter. Had it
// been an object, you could have object variables which could hold the
// aforementioned state information. like rocket flying behaviour object could be
// initialised with speed & fuel type along with maxHeight, whereas, default
// flying could take no inputs, apart from the interface's maxHeight.
//
// Also, after reading the solution from the book, the nice thing with the
// FlyBehaviour and QuackBehaviour is that finally you can have just the fly()
// method with no params, and offload all the state setting, and prerequisites
// using constructor NewXXX() and then finally call the fly() with no params.

// All types of Actions which a duck can perform
// TODO: find a way to group them better without impacting the ease of use.
type FlyAction func(maxHeight int)
type QuackAction func(noise string)
type SwimAction func()

var defaultSwim SwimAction = func() {
	fmt.Println("splash, splash!")
}

// RubberDuck is intended to satisfy the Duck interface
type RubberDuck struct {
	QuackBehaviour
}

func (rd *RubberDuck) Swim(action SwimAction) { action() }
func (rd *RubberDuck) Display()               {}

// RedHeadDuck is intended to satisfy the Duck interface
type RedHeadDuck struct {
	FlyBehaviour
	QuackBehaviour
}

func (rd *RedHeadDuck) Swim(action SwimAction) { action() }
func (rd *RedHeadDuck) Display()               {}

// testing
func main() {
	// creating ducks
	var d1 Duck = &RubberDuck{QuackBehaviour: &MuteQuack{}}
	d1.Display()
	d1.Swim(defaultSwim)
	d1.(*RubberDuck).quack()

	d2fly, err := NewRocketPoweredFlying(1000, 100, 999, "ethaneol")
	if err != nil {
		panic(err)
	}
	var d2 Duck = &RedHeadDuck{
		FlyBehaviour:   d2fly,
		QuackBehaviour: &DefaultQuack{},
	}
	d2.Display()
	d2.Swim(defaultSwim)
	d2.(*RedHeadDuck).quack()
	d2.(*RedHeadDuck).fly()
}
