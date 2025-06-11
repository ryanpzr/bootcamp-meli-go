package service

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/domain"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func myHandler() Product {
	product := *domain.NewProduct()
	repository := NewRepository(product)
	service := NewService(*repository)
	service.CreateList("../test/products_test.json")
	return *service
}

func TestProduct(t *testing.T) {
	t.Run("should get products", func(t *testing.T) {
		hd := myHandler()

		req := httptest.NewRequest("GET", "/products", nil)
		res := httptest.NewRecorder()
		hd.GetProducts(res, req)

		expectedCode := http.StatusOK
		expectedBody := `[{"id":1,"name":"Oil - Margarine","quantity":439,"code_value":"S82254D","is_published":true,"expiration":"15/12/2021","price":71.42},
{"id":2,"name":"Pineapple - Canned, Rings","quantity":345,"code_value":"M4637","is_published":true,"expiration":"09/08/2021","price":352.79}]`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should get by id products", func(t *testing.T) {
		hd := myHandler()

		req := httptest.NewRequest("GET", "/products/id?id=1", nil)
		res := httptest.NewRecorder()
		hd.GetProductsById(res, req)

		expectedCode := http.StatusOK
		expectedBody := `{"id":1,"name":"Oil - Margarine","quantity":439,"code_value":"S82254D","is_published":true,"expiration":"15/12/2021","price":71.42}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should post a product", func(t *testing.T) {
		hd := myHandler()

		body := strings.NewReader(`{"id": 0,"name": "Monitor","quantity": 20,"code_value": "T92DAH","is_published": false,"expiration": "07/10/2004","price": 99}`)
		req := httptest.NewRequest("POST", "/products", body)
		res := httptest.NewRecorder()
		hd.PostProduct(res, req)

		expectedCode := http.StatusOK
		expectedBody := `{"id": 3,"name": "Monitor","quantity": 20,"code_value": "T92DAH","is_published": false,"expiration": "07/10/2004","price": 99}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should delete a product", func(t *testing.T) {
		hd := myHandler()

		req := httptest.NewRequest("DELETE", "/products/M4637", nil)
		res := httptest.NewRecorder()
		hd.DeleteProduct(res, req)

		expectedCode := http.StatusOK
		expectedBody := `{"id":2,"name":"Pineapple - Canned, Rings","quantity":345,"code_value":"M4637","is_published":true,"expiration":"09/08/2021","price":352.79}`

		require.Equal(t, expectedCode, res.Code)
		require.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should return a error (Bad Request) when try delete a product with false id", func(t *testing.T) {
		hd := myHandler()

		req := httptest.NewRequest("DELETE", "/products/M4637dsa", nil)
		res := httptest.NewRecorder()
		hd.DeleteProduct(res, req)

		expectedCode := http.StatusBadRequest
		require.Equal(t, expectedCode, res.Code)
	})

	t.Run("should return a error (Not Found) when get a product with false id", func(t *testing.T) {
		hd := myHandler()

		req := httptest.NewRequest("GET", "/products/id?id=10", nil)
		res := httptest.NewRecorder()
		hd.GetProductsById(res, req)

		expectedCode := http.StatusNotFound

		require.Equal(t, expectedCode, res.Code)
	})
}
