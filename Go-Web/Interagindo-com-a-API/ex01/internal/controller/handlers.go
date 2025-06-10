package controller

import (
	"bootcamp-meli-go/Go-Web/Interagindo-com-a-API/ex01/internal/products/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service service.Product
}

func NewHandler(service service.Product) *Handler {
	return &Handler{service: service}
}

func (h *Handler) NewRouters(router *chi.Mux) {
	router.Get("/ping", h.service.GetPing)
	router.Get("/products", h.service.GetProducts)
	router.Get("/products/id", h.service.GetProductsById)
	router.Get("/products/search", h.service.GetProductsSearch)
	router.Post("/products", h.service.PostProduct)
	router.Put("/products", h.service.PutProduct)
	router.Patch("/products", h.service.PatchProduct)
	router.Delete("/products/{id}", h.service.DeleteProduct)
}
