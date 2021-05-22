// Implementing the solution from the book of course in idiomatic go
package main

import "fmt"

type FlyBehaviour interface {
	fly()
}

type DefaultFly struct {
	Height int
	Speed  int
}

func (df *DefaultFly) fly() {}

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

func (rpf *RocketPoweredFlying) fly() {}

type QuackBehaviour interface {
	quack()
}

type DefaultQuack struct {
	Frequency int
}

func (dq *DefaultQuack) quack() {}

type MuteQuack struct{}

func (mq *MuteQuack) quack() {}
