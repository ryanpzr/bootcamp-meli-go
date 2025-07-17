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

func TestProductsMySQL_GetRankingProds(t *testing.T) {
	db, err := sql.Open("txdb", "fantasy_products_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rp := NewProductsMySQL(db)
	sv := service.NewProductsDefault(rp)
	hd := handler.NewProductsDefault(sv)

	req := &http.Request{
		Method: "GET",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
	}
	res := httptest.NewRecorder()

	hd.GetRankingProd().ServeHTTP(res, req)

	expected := `{"data":[{"description":"Truffle Cups - Red","total":575.4},{"description":"Sword Pick Asst","total":271.92}]}`

	require.Equal(t, http.StatusOK, res.Code)
	require.Equal(t, http.MethodGet, req.Method)
	require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	require.Equal(t, expected, res.Body.String())
}
