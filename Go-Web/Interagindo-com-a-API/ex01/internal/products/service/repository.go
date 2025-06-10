package service

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/domain"
	"encoding/json"
	"os"
)

type repository struct{}

var productList []domain.Product

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetListProductsJson() ([]domain.Product, error) {

	fileJson, err := os.Open("products.json")
	if err != nil {
		return nil, err
	}
	defer fileJson.Close()

	if err := json.NewDecoder(fileJson).Decode(&productList); err != nil {
		return nil, err
	}

	return productList, nil
}

func (r *repository) SaveProduct(product domain.Product) ([]domain.Product, error) {
	productList = append(productList, product)
	return productList, nil
}
