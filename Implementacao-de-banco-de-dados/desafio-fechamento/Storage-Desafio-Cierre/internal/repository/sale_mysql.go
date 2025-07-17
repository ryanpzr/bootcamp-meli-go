package repository

import (
	"app/utils"
	"database/sql"
	"encoding/json"
	"log"

	"app/internal"
)

// NewSalesMySQL creates new mysql repository for sale entity.
func NewSalesMySQL(db *sql.DB) *SalesMySQL {
	return &SalesMySQL{db}
}

// SalesMySQL is the MySQL repository implementation for sale entity.
type SalesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all sales from the database.
func (r *SalesMySQL) FindAll() (s []internal.Sale, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `quantity`, `product_id`, `invoice_id` FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var sa internal.Sale
		// scan the row into the sale
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		// append the sale to the slice
		s = append(s, sa)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the sale into the database.
func (r *SalesMySQL) Save(s *internal.Sale) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)",
		(*s).Quantity, (*s).ProductId, (*s).InvoiceId,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*s).Id = int(id)

	return
}

func (r *SalesMySQL) Populate() (s []internal.Sale, err error) {
	file, err := utils.OpenFile("sales.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var sale []internal.Sale
	err = json.NewDecoder(file).Decode(&sale)
	if err != nil {
		log.Fatal(err)
	}

	var listSaleResponse []internal.Sale
	for _, sa := range sale {
		res, err := r.db.Exec(
			"INSERT INTO sales (`product_id`, `invoice_id`, `quantity`) VALUES (?, ?, ?)",
			(sa).ProductId, (sa).InvoiceId, (sa).Quantity,
		)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		sa.Id = int(id)
		listSaleResponse = append(listSaleResponse, sa)
	}

	return listSaleResponse, nil
}
