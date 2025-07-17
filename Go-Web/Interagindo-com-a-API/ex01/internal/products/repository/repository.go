package repository

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/domain"
	"errors"
)

type Repository struct {
	Product     domain.Product
	ProductList []domain.Product
}

func NewRepository(product domain.ProductSt) *Repository {
	return &Repository{Product: product.Product, ProductList: product.ListProduct}
}

func (r *Repository) GetProductsById(id int) (domain.Product, error) {
	for _, product := range r.ProductList {
		if product.Id == id {
			return product, nil
		}
	}

	return domain.Product{}, errors.New("Não foi encontrado registro que se aplica ao filtro enviado")
}

func (r *Repository) GetProductsSearch(priceInteger float64) ([]domain.Product, error) {
	var productSearchList []domain.Product
	for _, product := range r.ProductList {
		if product.Price >= priceInteger {
			productSearchList = append(productSearchList, product)
		}
	}

	if len(productSearchList) > 0 {
		return productSearchList, nil
	} else {
		return nil, errors.New("Não foi encontrado registros que se aplicam ao filtro enviado")
	}
}

func (r *Repository) SaveProduct(product domain.Product) (domain.Product, error) {
	list := append(r.ProductList, product)
	r.ProductList = list
	return product, nil
}

func (r *Repository) PutProduct(product domain.Product) (domain.Product, error) {
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

func (r *Repository) PatchProduct(product domain.Product) (domain.Product, error) {
	for i, p := range r.ProductList {
		if product.Code_value == p.Code_value {
			switch {
			case product.Expiration != "" && product.Expiration != p.Expiration:
				r.ProductList[i].Expiration = product.Expiration
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

func (r *Repository) DeleteProduct(code_value string) (domain.Product, error) {
	for i, p := range r.ProductList {
		if p.Code_value == code_value {
			result := r.ProductList[i]
			r.ProductList = append(r.ProductList[:i], r.ProductList[i+1:]...)
			return result, nil
		}
	}

	return domain.Product{}, errors.New("Não foi possivel achar um produto conforme o code enviado")
}
