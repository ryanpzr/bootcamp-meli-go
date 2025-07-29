package handler

import (
	"app/internal"
	"app/internal/repository"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductsDefault_Get(t *testing.T) {
	t.Run("case success", func(t *testing.T) {
		mockRepo := repository.NewProductsMapMock()
		hd := NewProductsDefault(mockRepo)

		mockRepo.On("SearchProducts", internal.ProductQuery{Id: 1}).Return(map[int]internal.Product{
			1: internal.Product{
				Id: 1,
				ProductAttributes: internal.ProductAttributes{
					Description: "Product 1",
					Price:       1.0,
					SellerId:    1,
				},
			},
		}, nil)

		req := httptest.NewRequest(http.MethodGet, "/product?id=1", nil)
		res := httptest.NewRecorder()

		hd.Get().ServeHTTP(res, req)

		expected := "{\"data\":{\"1\":{\"id\":1,\"description\":\"Product 1\",\"price\":1,\"seller_id\":1}},\"message\":\"success\"}"

		require.Equal(t, http.StatusOK, res.Code)
		require.Equal(t, "application/json", res.Header().Get("Content-Type"))
		require.JSONEq(t, expected, res.Body.String())

		mockRepo.AssertExpectations(t)
	})

	t.Run("Product not exist - case fail", func(t *testing.T) {
		mockRepo := repository.NewProductsMapMock()
		hd := NewProductsDefault(mockRepo)

		mockRepo.On("SearchProducts", internal.ProductQuery{Id: 2}).Return(map[int]internal.Product{}, errors.New("some error"))

		req := httptest.NewRequest(http.MethodGet, "/product?id=2", nil)
		res := httptest.NewRecorder()

		hd.Get().ServeHTTP(res, req)

		require.Equal(t, http.StatusInternalServerError, res.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("insert a different type to Id - case fail", func(t *testing.T) {
		mockRepo := repository.NewProductsMapMock()
		hd := NewProductsDefault(mockRepo)

		mockRepo.On("SearchProducts", internal.ProductQuery{Id: 2}).Return(map[int]internal.Product{}, errors.New("some error"))

		req := httptest.NewRequest(http.MethodGet, "/product?id=dois", nil)
		res := httptest.NewRecorder()

		hd.Get().ServeHTTP(res, req)

		require.Equal(t, http.StatusBadRequest, res.Code)
		require.Equal(t, `{"status":"Bad Request","message":"invalid id"}`, res.Body.String())
	})
}
