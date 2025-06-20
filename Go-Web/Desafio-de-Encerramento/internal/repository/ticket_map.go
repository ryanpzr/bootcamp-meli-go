package repository

import (
	"app/internal"
	"context"
	"errors"
)

func NewRepositoryTicketMap(db map[int]internal.TicketAttributes, dbFile string, lastId int) TicketInterface {
	return &RepositoryTicketMap{
		db:     db,
		dbFile: dbFile,
		lastId: lastId,
	}
}

type TicketInterface interface {
	GetTotalTicketsByDestCountry(dest string) (int, error)
	GetTicketsByDestCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error)
	GetAverageTicketsByDest(dest string) (float64, error)
}

type RepositoryTicketMap struct {
	db     map[int]internal.TicketAttributes
	lastId int
	dbFile string
}

func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

func (r *RepositoryTicketMap) GetTotalTicketsByDestCountry(dest string) (int, error) {
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

func (r *RepositoryTicketMap) GetTicketsByDestCountry(ctx context.Context, country string) (map[int]internal.TicketAttributes, error) {
	t := make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return t, nil
}

func (r *RepositoryTicketMap) GetAverageTicketsByDest(dest string) (float64, error) {
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
