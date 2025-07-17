package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"app/internal"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

// NewInvoicesDefault returns a new InvoicesDefault
func NewInvoicesDefault(sv internal.ServiceInvoice) *InvoicesDefault {
	return &InvoicesDefault{sv: sv}
}

// InvoicesDefault is a struct that returns the invoice handlers
type InvoicesDefault struct {
	// sv is the invoice's service
	sv internal.ServiceInvoice
}

// InvoiceJSON is a struct that represents a invoice in JSON format
type InvoiceJSON struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	Total      float64 `json:"total"`
	CustomerId int     `json:"customer_id"`
}

// GetAll returns all invoices
func (h *InvoicesDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		i, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting invoices")
			return
		}

		// response
		// - serialize
		ivJSON := make([]InvoiceJSON, len(i))
		for ix, v := range i {
			ivJSON[ix] = InvoiceJSON{
				Id:         v.Id,
				Datetime:   v.Datetime,
				Total:      v.Total,
				CustomerId: v.CustomerId,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "invoices found",
			"data":    ivJSON,
		})
	}
}

// RequestBodyInvoice is a struct that represents the request body for a invoice
type RequestBodyInvoice struct {
	Datetime   string  `json:"datetime"`
	Total      float64 `json:"total"`
	CustomerId int     `json:"customer_id"`
}

// Create creates a new invoice
func (h *InvoicesDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyInvoice
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing request body")
			return
		}

		// process
		// - deserialize
		i := internal.Invoice{
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   reqBody.Datetime,
				Total:      reqBody.Total,
				CustomerId: reqBody.CustomerId,
			},
		}
		// - save
		err = h.sv.Save(&i)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving invoice")
			return
		}

		// response
		// - serialize
		iv := InvoiceJSON{
			Id:         i.Id,
			Datetime:   i.Datetime,
			Total:      i.Total,
			CustomerId: i.CustomerId,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "invoice created",
			"data":    iv,
		})
	}
}

func (h *InvoicesDefault) Populate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.sv.Populate()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error: "+err.Error())
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"data": result,
		})
	}
}

func (h *InvoicesDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing id")
			return
		}

		var reqBody RequestBodyInvoice
		err = json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing request body")
			return
		}

		i := internal.Invoice{
			Id: idInt,
			InvoiceAttributes: internal.InvoiceAttributes{
				Datetime:   reqBody.Datetime,
				Total:      reqBody.Total,
				CustomerId: reqBody.CustomerId,
			},
		}

		result, err := h.sv.Update(&i)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error: "+err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": result,
		})
	}
}
