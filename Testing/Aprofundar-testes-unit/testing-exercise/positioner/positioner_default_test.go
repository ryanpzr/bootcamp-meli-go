package positioner

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPositionerDefault_GetLinearDistance(t *testing.T) {
	t.Run("coordinates is negativa", func(t *testing.T) {
		from := &Position{-10, -7, -5}
		to := &Position{-14, -4, -8}

		p := NewPositionerDefault()
		linearDistance := p.GetLinearDistance(from, to)

		require.Equal(t, 5.830951894845301, linearDistance)
	})

	t.Run("coordinates is postive", func(t *testing.T) {
		from := &Position{10, 7, 5}
		to := &Position{14, 4, 8}

		p := NewPositionerDefault()
		linearDistance := p.GetLinearDistance(from, to)

		require.Equal(t, 5.830951894845301, linearDistance)
	})

	t.Run("coordinates not return decimail", func(t *testing.T) {
		from := &Position{1, 1, 1}
		to := &Position{3, 4, 7}

		p := NewPositionerDefault()
		linearDistance := p.GetLinearDistance(from, to)

		require.Equal(t, float64(7), linearDistance)
	})
}
