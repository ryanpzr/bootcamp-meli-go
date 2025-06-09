package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Product struct {
	Id           int
	Name         string
	Quantity     int
	Code_value   string
	Is_published bool
	Experation   string
	Price        float64
}

var productList []Product

func main() {
	fileJson, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileJson.Close()

	if err := json.NewDecoder(fileJson).Decode(&productList); err != nil {
		fmt.Println(err)
		return
	}

	for _, prod := range productList {
		fmt.Printf("Nome: %s, Preço: %.2f\n", prod.Name, prod.Price)
	}

	router := http.NewServeMux()

	router.HandleFunc("/ping", getPing)
	router.HandleFunc("/products", getProducts)
	router.HandleFunc("/products/id", getProductsById)
	router.HandleFunc("/products/search", getProductsSearch)

	http.ListenAndServe(":8080", router)
}

func getPing(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productList)
}

func getProductsById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idInteger, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	for _, product := range productList {
		if product.Id == idInteger {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
}

func getProductsSearch(w http.ResponseWriter, r *http.Request) {
	price := r.URL.Query().Get("price")
	priceInteger, err := strconv.ParseFloat(price, 64)
	if err != nil {
		panic(err)
	}

	var productSearchList []Product

	for _, product := range productList {
		if product.Price >= priceInteger {
			productSearchList = append(productSearchList, product)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productSearchList)
}
