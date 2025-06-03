package main

import (
	"errors"
	"fmt"
)

type Product struct {
	id          int64
	name        string
	price       float64
	description string
	category    string
}

func (p Product) GetPrice() float64 {
	return p.price
}

type ProductsRepository struct {
	products []Product
}

func (p *ProductsRepository) Save(product Product) {
	p.products = append(p.products, product)
}

func (p *ProductsRepository) GetAll() {
	for _, product := range p.products {
		fmt.Println(product)
	}
}

func (p *ProductsRepository) GetById(id int64) (Product, error) {
	for _, product := range p.products {
		if product.id == id {
			return product, nil
		}
	}

	return Product{}, errors.New("Produto não encontrado na lista.")
}

func main() {
	p := Product{
		id:          1,
		name:        "Macbook",
		price:       36.240,
		description: "Macbook Apple 9geração",
		category:    "Notebook",
	}
	p2 := Product{
		id:          2,
		name:        "Monitor",
		price:       6.240,
		description: "Monitor",
		category:    "Monitor",
	}

	productsRepository := ProductsRepository{}
	productsRepository.Save(p)
	productsRepository.Save(p2)

	pr, err := productsRepository.GetById(1)
	if err != nil {
		return
	}
	fmt.Println("Produto capturado através do Id: ", pr.GetPrice())

	productsRepository.GetAll()
}
