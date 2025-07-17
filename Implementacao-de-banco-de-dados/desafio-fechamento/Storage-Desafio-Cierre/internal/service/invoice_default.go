package service

import (
	"app/internal"
	"errors"
)

// NewInvoicesDefault creates new default service for invoice entity.
func NewInvoicesDefault(rp internal.RepositoryInvoice) *InvoicesDefault {
	return &InvoicesDefault{rp}
}

// InvoicesDefault is the default service implementation for invoice entity.
type InvoicesDefault struct {
	// rp is the repository for invoice entity.
	rp internal.RepositoryInvoice
}

// FindAll returns all invoices.
func (s *InvoicesDefault) FindAll() (i []internal.Invoice, err error) {
	i, err = s.rp.FindAll()
	return
}

// Save saves the invoice.
func (s *InvoicesDefault) Save(i *internal.Invoice) (err error) {
	err = s.rp.Save(i)
	return
}

func (s *InvoicesDefault) Populate() (i []internal.Invoice, err error) {
	i, err = s.rp.Populate()
	return
}

func (s *InvoicesDefault) Update(invoice *internal.Invoice) (i internal.Invoice, err error) {
	if invoice.CustomerId == 0 {
		return i, errors.New("customer id can not be zero or nil")
	}

	i, err = s.rp.Update(invoice)
	return
}
