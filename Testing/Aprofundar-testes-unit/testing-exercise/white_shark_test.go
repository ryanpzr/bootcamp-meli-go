package hunt_test

import (
	"github.com/stretchr/testify/require"
	hunt "testdoubles"
	"testing"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 300)
		tuna := hunt.NewTuna("Claudio", 200)
		err := shark.Hunt(tuna)

		require.Nil(t, err)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		shark := hunt.NewWhiteShark(false, false, 300)
		tuna := hunt.NewTuna("Claudio", 200)
		err := shark.Hunt(tuna)

		require.Equal(t, err, hunt.ErrSharkIsNotHungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, true, 300)
		tuna := hunt.NewTuna("Claudio", 200)
		err := shark.Hunt(tuna)

		require.Equal(t, err, hunt.ErrSharkIsTired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 100)
		tuna := hunt.NewTuna("Claudio", 200)
		err := shark.Hunt(tuna)

		require.Equal(t, err, hunt.ErrSharkIsSlower)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		shark := hunt.NewWhiteShark(true, false, 300)
		err := shark.Hunt(nil)

		require.Equal(t, err, hunt.ErrTunaIsNil)
	})
}
