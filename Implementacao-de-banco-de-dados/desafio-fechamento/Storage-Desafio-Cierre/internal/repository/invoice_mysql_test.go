package repository

import (
	"app/internal/handler"
	"app/internal/service"
	"bytes"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestInvoicesMySQL_Update(t *testing.T) {
	db, err := sql.Open("txdb", "fantasy_products_test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rp := NewInvoicesMySQL(db)
	sv := service.NewInvoicesDefault(rp)
	hd := handler.NewInvoicesDefault(sv)

	r := chi.NewRouter()
	r.Put("/invoices/{id}", hd.Update())

	body := []byte(`{"datetime":"2008-11-24 00:00:00","total":35.0,"customer_id":2}`)
	req := &http.Request{
		Method: "PUT",
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		URL:  &url.URL{Path: `/invoices/2`},
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	expected := `{"data":{"id":2,"datetime":"2008-11-24 00:00:00","total":35,"customer_id":2}}`

	require.Equal(t, res.Code, http.StatusOK)
	require.Equal(t, http.MethodPut, req.Method)
	require.Equal(t, "application/json", res.Header().Get("Content-Type"))
	require.Equal(t, expected, res.Body.String())
}
