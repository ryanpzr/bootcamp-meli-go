package repository

import (
	"app/internal"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductsMap_SearchProducts(t *testing.T) {
	t.Run("case success", func(t *testing.T) {
		rpMock := NewProductsMapMock()

		rpMock.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(map[int]internal.Product{
			1: internal.Product{
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       1.0,
					SellerId:    1,
				},
			},
		}, nil)

		p, err := rpMock.SearchProducts(internal.ProductQuery{Id: 1})
		require.NoError(t, err)
		pJson, err := json.Marshal(p)
		require.NoError(t, err)

		expected := "{\"1\":{\"Id\":1,\"Description\":\"Product 1\",\"Price\":1,\"SellerId\":1}}"

		require.Equal(t, expected, string(pJson))
	})

	t.Run("case fail", func(t *testing.T) {
		rpMock := NewProductsMapMock()

		rpMock.On("SearchProducts", internal.ProductQuery{Id: 2}).Return(map[int]internal.Product{}, errors.New("error"))

		p, err := rpMock.SearchProducts(internal.ProductQuery{Id: 2})

		require.Error(t, err)
		require.Equal(t, 0, len(p))
	})
}
