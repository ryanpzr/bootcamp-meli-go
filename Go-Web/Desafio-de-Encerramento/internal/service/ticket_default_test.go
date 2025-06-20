package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "Indonesia",
				Hour:    "10:00",
				Price:   100,
			},
		}
		rp := repository.NewRepositoryTicketMock(db, "", 0)
		sv := service.NewServiceTicketDefault(rp)

		r := chi.NewRouter()
		r.Get("/tickets/getTotalByCountry/{dest}", sv.GetTotalTicketsByDestCountry)
		req := httptest.NewRequest("GET", "/tickets/getTotalByCountry/Indonesia", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedTotal := "1"
		require.Equal(t, expectedTotal, res.Body.String())
	})

	t.Run("success to get tickets", func(t *testing.T) {
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "Indonesia",
				Hour:    "10:00",
				Price:   100,
			},
		}
		rp := repository.NewRepositoryTicketMock(db, "", 0)
		sv := service.NewServiceTicketDefault(rp)

		r := chi.NewRouter()
		r.Get("/tickets/getByCountry/{dest}", sv.GetTicketsByDestCountry)
		req := httptest.NewRequest("GET", "/tickets/getByCountry/Indonesia", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedTotal := `{"1":{"name":"John","email":"johndoe@gmail.com","country":"Indonesia","hour":"10:00","price":100}}
`
		require.Equal(t, expectedTotal, res.Body.String())
	})

	t.Run("success to get average tickets", func(t *testing.T) {
		db := map[int]internal.TicketAttributes{
			1: {
				Name:    "John",
				Email:   "johndoe@gmail.com",
				Country: "Indonesia",
				Hour:    "10:00",
				Price:   100,
			},
			2: {
				Name:    "Alex",
				Email:   "alex@gmail.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100,
			},
			3: {
				Name:    "Mary",
				Email:   "mary@gmail.com",
				Country: "Indonesia",
				Hour:    "10:00",
				Price:   100,
			},
			4: {
				Name:    "Jane",
				Email:   "jane@gmail.com",
				Country: "Argentina",
				Hour:    "10:00",
				Price:   100,
			},
		}
		rp := repository.NewRepositoryTicketMock(db, "", 0)
		sv := service.NewServiceTicketDefault(rp)

		r := chi.NewRouter()
		r.Get("/tickets/getAverage/{dest}", sv.GetAverageTicketsByDest)
		req := httptest.NewRequest("GET", "/tickets/getAverage/Indonesia", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		expectedTotal := `50
`
		require.Equal(t, expectedTotal, res.Body.String())
	})
}
