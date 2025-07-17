package main

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/controller"
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/domain"
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/repository"
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	product := *domain.NewProduct()
	repository := repository.NewRepository(product)
	service := service.NewService(*repository)
	service.CreateList("../products.json")
	handler := controller.NewHandler(*service)
	handler.NewRouters(router)
	http.ListenAndServe(":8080", router)
}
