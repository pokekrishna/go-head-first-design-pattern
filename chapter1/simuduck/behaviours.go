// Implementing the solution from the book of course in idiomatic go
package main

import "fmt"

type FlyBehaviour interface {
	Fly()
}

type DefaultFly struct {
	Height int
	Speed  int
}

func (df *DefaultFly) Fly() {
	fmt.Printf("flying normally at speed %v and height %v\n", df.Speed, df.Height)
}

type RocketPoweredFlying struct {
	Height         int
	Speed          int
	RocketCapacity int
	FuelType       string
}

func (rpf *RocketPoweredFlying) validateCapacityAccordingToLaw() error {
	if rpf.RocketCapacity > 1200 {
		return fmt.Errorf("This country does not allow rocket capacity more than 1200cc")
	}
	return nil
}

// NewRocketPoweredFlying() is just an hyped example to demonstrate how the
// solution from the book has capability to much more.
// Doing some work before returning the object satisfying the interface
func NewRocketPoweredFlying(height int, speed int, rocketcap int, fuel string) (FlyBehaviour, error) {
	rpf := &RocketPoweredFlying{
		Height:         height,
		Speed:          speed,
		RocketCapacity: rocketcap,
		FuelType:       fuel,
	}

	if err := rpf.validateCapacityAccordingToLaw(); err != nil {
		return nil, err
	}
	return rpf, nil
}

func (rpf *RocketPoweredFlying) Fly() {
	fmt.Printf("flying like a rocket at speed %v and height %v\n", rpf.Speed, rpf.Height)
}

type QuackBehaviour interface {
	Quack()
}

type DefaultQuack struct {
	Frequency int
}

func (dq *DefaultQuack) Quack() {}

type MuteQuack struct{}

func (mq *MuteQuack) Quack() {}

type SwimBehaviour interface {
	Swim()
}

type DefaultSwimBehaviour struct {
	Speed int
}
func (dsb *DefaultSwimBehaviour) Swim(){}

type UnderwaterSwimBehaviour struct {
	Depth int
	Speed int
}
func (uwsb *UnderwaterSwimBehaviour) Swim(){}