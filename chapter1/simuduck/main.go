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

// All types of Actions which a duck can perform
// TODO: find a way to group them better without impacting the ease of use.
type FlyAction func(maxHeight int)
type QuackAction func(noise string)
type SwimAction func()

// Default implementations of actions, for code reuse
// TODO: looks like this could be some design pattern, maybe Factory?
var defaultFly FlyAction = func(maxHeight int) {
	fmt.Println("flying at height", maxHeight)
}

var rocketPoweredFlying = func(maxHeight int) {
	fmt.Println("rocket powered flying at height, ")
}

var defaultQuack QuackAction = func(noise string) {
	fmt.Println("making noise", noise)
}

var defaultSwim SwimAction = func() {
	fmt.Println("splash, splash!")
}

// RubberDuck is intended to satisfy the Duck interface
type RubberDuck struct {
	quack QuackAction
}
func (rd *RubberDuck) Swim(action SwimAction){action()}
func (rd *RubberDuck) Display(){}

// RedHeadDuck is intended to satisfy the Duck interface
type RedHeadDuck struct {
	fly FlyAction
	quack QuackAction
}
func (rd *RedHeadDuck) Swim(action SwimAction){action()}
func (rd *RedHeadDuck) Display(){}

// creating ducks
var d1 Duck = &RubberDuck{quack: defaultQuack}
var d2 Duck = &RedHeadDuck{
	fly:   defaultFly,
	quack: defaultQuack,
}

// testing
func main() {
	d1.Display()
	d1.Swim(defaultSwim)
	d1.(*RubberDuck).quack("squeak")

	d2.Display()
	d2.Swim(defaultSwim)
	d2.(*RedHeadDuck).quack("quack")
	d2.(*RedHeadDuck).fly(1)
}
