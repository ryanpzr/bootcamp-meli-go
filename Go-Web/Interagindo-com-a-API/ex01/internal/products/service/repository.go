package service

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/domain"
	"encoding/json"
	"errors"
	"os"
)

type repository struct {
	ProductList []domain.Product
}

func NewRepository() *repository {
	return &repository{}
}

func (r *repository) GetListProductsJson() ([]domain.Product, error) {
	if len(r.ProductList) > 0 {
		return r.ProductList, nil
	}

	fileJson, err := os.Open("../products.json")
	if err != nil {
		return nil, err
	}
	defer fileJson.Close()

	if err := json.NewDecoder(fileJson).Decode(&r.ProductList); err != nil {
		return nil, err
	}

	return r.ProductList, nil
}

func (r *repository) SaveProduct(product domain.Product) (domain.Product, error) {
	list := append(r.ProductList, product)
	r.ProductList = list
	return product, nil
}

func (r *repository) PutProduct(product domain.Product) (domain.Product, error) {
	for i, p := range r.ProductList {
		if product.Code_value == p.Code_value {
			if product.Id != p.Id {
				return domain.Product{}, errors.New("Id não deve ser diferente do já cadastrado")
			}
			r.ProductList[i] = product
			return r.ProductList[i], nil
		}
	}

	return domain.Product{}, errors.New("Não foi possivel achar um produto conforme o code enviado")
}

func (r *repository) PatchProduct(product domain.Product) (domain.Product, error) {
	for i, p := range r.ProductList {
		if product.Code_value == p.Code_value {
			switch {
			case product.Experation != "" && product.Experation != p.Experation:
				r.ProductList[i].Experation = product.Experation
			case product.Name != "" && product.Name != p.Name:
				r.ProductList[i].Name = product.Name
			case product.Price != 0 && product.Price != p.Price:
				r.ProductList[i].Price = product.Price
			case product.Quantity != 0 && product.Quantity != p.Quantity:
				r.ProductList[i].Quantity = product.Quantity
			default:
				break
			}

			return r.ProductList[i], nil
		}
	}

	return domain.Product{}, errors.New("Não foi possivel achar um produto conforme o code enviado")
}

func (r *repository) DeleteProduct(code_value string) (domain.Product, error) {
	for i, p := range r.ProductList {
		if p.Code_value == code_value {
			r.ProductList = append(r.ProductList[:i], r.ProductList[i+1:]...)
			return r.ProductList[i], nil
		}
	}

	return domain.Product{}, errors.New("Não foi possivel achar um produto conforme o code enviado")
}
