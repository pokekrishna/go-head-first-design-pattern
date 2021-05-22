package main

// Q: solving HF DP, SimUDuck problem

// Duck interface enforces to have Display() in all the
// implementations of Duck.
//
// Display is intended to be implemented by the dynamic type.
//
// Swim (,unlike Fly()) is intended to be present in all duck types however in order
// to aid code reuse it can simply just call the Action which it is
// passed.
type Duck interface {
	Display()
	Swim(SwimBehaviour)
	//Swim(action SwimBehaviour)
}

// RubberDuck is intended to satisfy the Duck interface
type RubberDuck struct {
	QuackBehaviour
}

func (rd *RubberDuck) Display()               {}
func (rd *RubberDuck) Swim(behaviour SwimBehaviour){
	// Promoting code reuse
	behaviour.Swim()
}

// RedHeadDuck is intended to satisfy the Duck interface
type RedHeadDuck struct {
	FlyBehaviour
	QuackBehaviour
}

func (rd *RedHeadDuck) Display()               {}
func (rd *RedHeadDuck) Swim(behaviour SwimBehaviour){
	// Promoting code reuse
	behaviour.Swim()
}

// testing
func main() {
	// creating ducks
	var d1 Duck = &RubberDuck{QuackBehaviour: &MuteQuack{}}
	d1.Display()
	d1.Swim(&DefaultSwimBehaviour{Speed: 100})
	d1.(*RubberDuck).Quack()

	d2fly, err := NewRocketPoweredFlying(1000, 100, 999, "ethanol")
	if err != nil {
		panic(err)
	}
	var d2 Duck = &RedHeadDuck{
		FlyBehaviour:   d2fly,
		QuackBehaviour: &DefaultQuack{},
	}
	d2.Display()
	d2.Swim(&UnderwaterSwimBehaviour{
		Depth: 100,
		Speed: 90,
	})
	d2.(*RedHeadDuck).Quack()
	d2.(*RedHeadDuck).Fly()
}
