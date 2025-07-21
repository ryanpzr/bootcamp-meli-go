package simulator

import (
	"github.com/stretchr/testify/require"
	"testdoubles/positioner"
	"testing"
)

func TestCatchSimulatorDefault_CanCatch(t *testing.T) {
	t.Run("the subject reaches the other", func(t *testing.T) {
		hunter := &Subject{
			Position: &positioner.Position{X: 31, Y: 11, Z: 55},
			Speed:    300,
		}
		prey := &Subject{
			Position: &positioner.Position{X: 14, Y: 4, Z: 8},
			Speed:    200,
		}
		ps := positioner.NewPositionerDefault()

		catchSimulator := NewCatchSimulatorDefault(100, ps)
		ok := catchSimulator.CanCatch(hunter, prey)

		require.True(t, ok)
	})

	t.Run("the subject not reaches the other", func(t *testing.T) {
		hunter := &Subject{
			Position: &positioner.Position{X: 31, Y: 11, Z: 55},
			Speed:    150,
		}

		prey := &Subject{
			Position: &positioner.Position{X: 14, Y: 4, Z: 8},
			Speed:    200,
		}

		ps := positioner.NewPositionerDefault()

		catchSimulator := NewCatchSimulatorDefault(100, ps)
		ok := catchSimulator.CanCatch(hunter, prey)

		require.False(t, ok)
	})

	t.Run("the subject cannot reach the other, even though he is faster, because he does not have time", func(t *testing.T) {
		hunter := &Subject{
			Position: &positioner.Position{X: 452, Y: 320, Z: 780},
			Speed:    300,
		}

		prey := &Subject{
			Position: &positioner.Position{X: 550, Y: 230, Z: 402},
			Speed:    200,
		}

		ps := positioner.NewPositionerDefault()

		catchSimulator := NewCatchSimulatorDefault(2, ps)
		ok := catchSimulator.CanCatch(hunter, prey)

		require.False(t, ok)
	})

	t.Run("simulates the mock implementation of the CanCatch method with the possibility of checking whether it was called", func(t *testing.T) {
		simulatorMock := NewCatchSimulatorMock()

		hunter := &Subject{
			Position: &positioner.Position{X: 452, Y: 320, Z: 780},
			Speed:    300,
		}

		prey := &Subject{
			Position: &positioner.Position{X: 550, Y: 230, Z: 402},
			Speed:    200,
		}

		simulatorMock.On("CanCatch", hunter, prey).Return(false)

		ok := simulatorMock.CanCatch(hunter, prey)

		require.False(t, ok)
		simulatorMock.AssertExpectations(t)
	})
}
