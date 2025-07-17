package repository

import (
	"app/internal/handler"
	"app/internal/service"
	"database/sql"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomersMySQL_GetCondition(t *testing.T) {
	db, err := sql.Open("txdb", "fantasy_products_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rp := NewCustomersMySQL(db)
	sv := service.NewCustomersDefault(rp)
	hd := handler.NewCustomersDefault(sv)

	req := &http.Request{
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
	res := httptest.NewRecorder()

	hd.GetCondition().ServeHTTP(res, req)

	expected := `{"data":[{"name":"activo","invoice_total":23},{"name":"inactivo","invoice_total":34.5}]}`

	require.Equal(t, http.StatusOK, res.Code)
	require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	require.Equal(t, expected, res.Body.String())
}

func TestCustomersMySQL_GetRankingCus(t *testing.T) {
	db, err := sql.Open("txdb", "fantasy_products_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rp := NewCustomersMySQL(db)
	sv := service.NewCustomersDefault(rp)
	hd := handler.NewCustomersDefault(sv)

	req := &http.Request{
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
	res := httptest.NewRecorder()

	hd.GetRankingCus().ServeHTTP(res, req)

	expected := `{"data":[{"first_name":"Marcos","last_name":"Frota","total":34.5},{"first_name":"Ryan","last_name":"Pereira","total":23}]}`

	require.Equal(t, http.StatusOK, res.Code)
	require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	require.Equal(t, expected, res.Body.String())
}
