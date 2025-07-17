package service

import "app/internal"

// NewProductsDefault creates new default service for product entity.
func NewProductsDefault(rp internal.RepositoryProduct) *ProductsDefault {
	return &ProductsDefault{rp}
}

// ProductsDefault is the default service implementation for product entity.
type ProductsDefault struct {
	// rp is the repository for product entity.
	rp internal.RepositoryProduct
}

// FindAll returns all products.
func (s *ProductsDefault) FindAll() (p []internal.Product, err error) {
	p, err = s.rp.FindAll()
	return
}

// Save saves the product.
func (s *ProductsDefault) Save(p *internal.Product) (err error) {
	err = s.rp.Save(p)
	return
}

func (s *ProductsDefault) Populate() (p []internal.Product, err error) {
	p, err = s.rp.Populate()
	return
}

func (s *ProductsDefault) GetRankingProd() (p []internal.ProductRanking, err error) {
	p, err = s.rp.GetRankingProd()
	return
}
