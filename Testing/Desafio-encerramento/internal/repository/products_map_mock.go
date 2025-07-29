package repository

import (
	"app/internal"
	"github.com/stretchr/testify/mock"
)

func NewProductsMapMock() *ProductsMapMock {
	return &ProductsMapMock{}
}

type ProductsMapMock struct {
	mock.Mock
}

func (r *ProductsMapMock) SearchProducts(query internal.ProductQuery) (p map[int]internal.Product, err error) {
	args := r.Called(query)
	return args.Get(0).(map[int]internal.Product), args.Error(1)
}
