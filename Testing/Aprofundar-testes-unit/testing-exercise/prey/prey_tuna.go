package prey

import (
	"math/rand"
	"testdoubles/positioner"
)

// CreateTuna creates a new Tuna
func CreateTuna() Prey {
	// default config
	// -> max speed: 252 m/s
	speed := rand.Float64() * 252.0 + 15.0
	// -> position: random
	position := &positioner.Position{
		X: rand.Float64() * 500,
		Y: rand.Float64() * 500,
		Z: rand.Float64() * 500,
	}

	return &Tuna{
		speed: speed,
		position: position,
	}
}

// Tuna is an implementation of the Prey interface
type Tuna struct {
	// speed of the tuna
	speed float64
	// position of the tuna
	position *positioner.Position
}

// GetSpeed returns the speed of the tuna
func (t *Tuna) GetSpeed() (speed float64) {
	// speed is the speed in m/s of the tuna
	speed = t.speed
	return
}

func (t *Tuna) GetPosition() (position *positioner.Position) {
	position = t.position
	return
}