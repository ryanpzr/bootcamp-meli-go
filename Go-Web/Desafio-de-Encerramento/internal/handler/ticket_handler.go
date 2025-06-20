package handler

import (
	"app/internal/service"
	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	service service.ServiceTicketDefault
	router  *chi.Mux
}

func NewHandler(sv *service.ServiceTicketDefault, rt *chi.Mux) *TicketHandler {
	return &TicketHandler{service: *sv, router: rt}
}

func (h *TicketHandler) NewRouters() {
	h.router.Route("/tickets", func(rt chi.Router) {
		rt.Get("/getTotalByCountry/{dest}", h.service.GetTotalTicketsByDestCountry)
		rt.Get("/getByCountry/{dest}", h.service.GetTicketsByDestCountry)
		rt.Get("/getAverage/{dest}", h.service.GetAverageTicketsByDest)
	})
}
