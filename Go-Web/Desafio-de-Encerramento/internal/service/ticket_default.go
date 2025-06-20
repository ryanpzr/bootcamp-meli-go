package service

import (
	"app/internal"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

// ServiceTicketDefault represents the default service of the tickets
type ServiceTicketDefault struct {
	// rp represents the repository of the tickets
	rp ServiceTicketInterface
}

type ServiceTicketInterface interface {
	GetTotalTicketsByDestCountry(dest string) (int, error)
	GetTicketsByDestCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error)
	GetAverageTicketsByDest(dest string) (float64, error)
}

// NewServiceTicketDefault creates a new default service of the tickets
func NewServiceTicketDefault(rp ServiceTicketInterface) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

// GetTotalTickets returns the total number of tickets
func (s *ServiceTicketDefault) GetTotalTicketsByDestCountry(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método inserido não permitido", http.StatusMethodNotAllowed)
		return
	}

	dest := chi.URLParam(r, "dest")

	totalTickets, err := s.rp.GetTotalTicketsByDestCountry(dest)
	if err != nil {
		http.Error(w, "Erro ao recuperar o total de tickets: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(strconv.Itoa(totalTickets)))
}

func (s *ServiceTicketDefault) GetTicketsByDestCountry(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método inserido não permitido", http.StatusMethodNotAllowed)
		return
	}

	dest := chi.URLParam(r, "dest")

	ctx := context.Background()
	totalTickets, err := s.rp.GetTicketsByDestCountry(ctx, dest)
	if err != nil {
		http.Error(w, "Erro ao recuperar média de tickets: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(totalTickets)
}

func (s *ServiceTicketDefault) GetAverageTicketsByDest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Método inserido não permitido", http.StatusMethodNotAllowed)
		return
	}

	dest := chi.URLParam(r, "dest")

	totalAvTickets, err := s.rp.GetAverageTicketsByDest(dest)
	if err != nil {
		http.Error(w, "Erro ao recuperar média de tickets: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(totalAvTickets)
}
