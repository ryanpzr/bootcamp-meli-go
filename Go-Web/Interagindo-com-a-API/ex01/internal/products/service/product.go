package service

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/domain"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	repo repository
}

func NewService(r repository) *Product {
	return &Product{repo: r}
}

func (p *Product) GetPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := p.repo.GetListProductsJson()
	if err != nil {
		http.Error(w, "Erro ao recuperar lista de produtos", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productList)
}

func (p *Product) GetProductsById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	productList, err := p.repo.GetListProductsJson()
	if err != nil {
		http.Error(w, "Erro ao recuperar lista de produtos", http.StatusBadRequest)
	}

	for _, product := range productList {
		if product.Id == idInteger {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func (p *Product) GetProductsSearch(w http.ResponseWriter, r *http.Request) {
	price := r.URL.Query().Get("price")
	priceInteger, err := strconv.ParseFloat(price, 64)
	if err != nil {
		panic(err)
	}

	var productSearchList []domain.Product

	productList, err := p.repo.GetListProductsJson()
	if err != nil {
		http.Error(w, "Erro ao recuperar lista de produtos", http.StatusBadRequest)
	}

	for _, product := range productList {
		if product.Price >= priceInteger {
			productSearchList = append(productSearchList, product)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productSearchList)
}

func (p *Product) PostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "O método da requisição deve ser do tipo POST", http.StatusBadRequest)
		return
	}

	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	productList, err := p.repo.GetListProductsJson()
	if err != nil {
		http.Error(w, "Erro ao recuperar lista de produtos", http.StatusBadRequest)
	}

	product.Id = productList[len(productList)-1].Id + 1

	isEmpty, message := validateAttribs(product, "POST")
	if isEmpty {
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	for _, p := range productList {
		if p.Code_value == product.Code_value {
			http.Error(w, "Código inserido já existe cadastrado no sistema", http.StatusBadRequest)
			return
		}
		if p.Name == product.Name {
			http.Error(w, "Nome inserido já existe cadastrado no sistema", http.StatusBadRequest)
			return
		}
	}

	listProduct, err := p.repo.SaveProduct(product)
	if err != nil {
		http.Error(w, "Erro ao inserir Produto no banco", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&listProduct)
}

func (p *Product) PutProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "O método da requisição deve ser do tipo POST", http.StatusBadRequest)
		return
	}

	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	isEmpty, message := validateAttribs(product, "PUT")
	if isEmpty {
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	newProduct, err := p.repo.PutProduct(product)
	if err != nil {
		http.Error(w, "Erro ao editar produto: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newProduct)
}

func (p *Product) PatchProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PATCH" {
		http.Error(w, "O método da requisição deve ser do tipo PATCH", http.StatusBadRequest)
		return
	}

	var product domain.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		panic(err)
	}

	isEmpty, message := validateAttribs(product, "PATCH")
	if isEmpty {
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	newProduct, err := p.repo.PatchProduct(product)
	if err != nil {
		http.Error(w, "Erro ao editar produto: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&newProduct)
}

func (p *Product) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "O método da requisição deve ser do tipo DELETE", http.StatusBadRequest)
		return
	}

	path := r.URL.Path
	id := strings.Split(path, "/")[2]

	deletedProduct, err := p.repo.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Erro ao deletar produto: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&deletedProduct)
}

func validateAttribs(p domain.Product, typeMethod string) (bool, string) {
	switch {
	case p.Quantity == 0 && typeMethod != "PATCH":
		return true, "Quantidade não pode ser vazio"
	case p.Code_value == "":
		return true, "Code_value não pode ser vazio"
	case p.Experation == "" && typeMethod != "PATCH" || p.Experation != "" && !isDate(p.Experation):
		if p.Experation != "" {
			return true, "Experation deve estar no formato: 00/00/0000"
		}
		return true, "Experation não pode ser vazio"
	case p.Price == 0 && typeMethod != "PATCH":
		return true, "Preço não pode ser vazio"
	case p.Name == "" && typeMethod != "PATCH":
		return true, "Name não pode ser vazio"
	default:
		return false, ""
	}
}

func isDate(date string) bool {
	layout := "01/02/2006"
	_, err := time.Parse(layout, date)
	return err == nil
}
