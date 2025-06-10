package main

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/controller"
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	repository := service.NewRepository()
	service := service.NewService(*repository)
	handler := controller.NewHandler(*service)
	handler.NewRouters(router)
	http.ListenAndServe(":8080", router)
}
