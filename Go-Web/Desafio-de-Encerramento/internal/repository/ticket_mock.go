package repository

import (
	"app/internal"
	"context"
	"errors"
)

func NewRepositoryTicketMock(db map[int]internal.TicketAttributes, dbFile string, lastId int) TicketMockInterface {
	return &RepositoryTicketMock{
		db:     db,
		dbFile: dbFile,
		lastId: lastId,
	}
}

type TicketMockInterface interface {
	GetTotalTicketsByDestCountry(dest string) (int, error)
	GetTicketsByDestCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error)
	GetAverageTicketsByDest(dest string) (float64, error)
}

type RepositoryTicketMock struct {
	db     map[int]internal.TicketAttributes
	lastId int
	dbFile string
}

func (r *RepositoryTicketMock) GetTotalTicketsByDestCountry(dest string) (int, error) {
	if r.db == nil {
		return 0, errors.New("Database is null")
	}

	var quantityTickets int
	for _, ti := range r.db {
		if ti.Country == dest {
			quantityTickets++
		}
	}

	if quantityTickets == 0 {
		return 0, errors.New("No tickets found")
	}

	return quantityTickets, nil
}

func (r *RepositoryTicketMock) GetTicketsByDestCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error) {
	t := make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return t, nil
}

func (r *RepositoryTicketMock) GetAverageTicketsByDest(dest string) (float64, error) {
	if r.db == nil {
		return 0, errors.New("Database is null")
	}

	var sumOfTickets float64 = 0
	for _, ti := range r.db {
		if ti.Country == dest {
			sumOfTickets++
		}
	}

	if sumOfTickets == 0 {
		return 0, errors.New("No tickets found")
	}

	average := sumOfTickets / float64(len(r.db)) * 100
	return average, nil
}
