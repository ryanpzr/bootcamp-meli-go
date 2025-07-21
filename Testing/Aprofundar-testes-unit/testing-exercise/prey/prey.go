package prey

import "testdoubles/positioner"

// Prey is an interface that represents a prey
type Prey interface {
	// GetSpeed returns the speed of the prey
	GetSpeed() (speed float64)
	// GetPosition returns the position of the prey
	GetPosition() (position *positioner.Position)
}