package repository

import (
	"app/internal"
	"app/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []internal.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv internal.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *internal.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
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
	(*i).Id = int(id)

	return
}

func (r *InvoicesMySQL) Populate() (i []internal.Invoice, err error) {
	file, err := utils.OpenFile("invoices.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var invoice []internal.Invoice
	err = json.NewDecoder(file).Decode(&invoice)
	if err != nil {
		log.Fatal(err)
	}

	var listInvoiceResponse []internal.Invoice
	for _, inv := range invoice {
		if inv.CustomerId == 0 {
			byte, err := json.Marshal(inv)
			if err != nil {
			}
			return nil, fmt.Errorf(string(byte))
		}

		res, err := r.db.Exec(
			"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
			(inv).Datetime, (inv).Total, (inv).CustomerId,
		)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		inv.Id = int(id)
		listInvoiceResponse = append(listInvoiceResponse, inv)
	}

	return listInvoiceResponse, nil
}

func (r *InvoicesMySQL) Update(i *internal.Invoice) (internal.Invoice, error) {
	query := `UPDATE invoices SET datetime=?, total=?, customer_id=? WHERE id=?`
	row, err := r.db.Exec(query, (*i).Datetime, (*i).Total, (*i).CustomerId, (*i).Id)
	if err != nil {
		return internal.Invoice{}, err
	}

	rows, err := row.RowsAffected()
	if err != nil {
		return internal.Invoice{}, err
	}

	if rows == 0 {
		var exists bool
		err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM customers WHERE id=?)", (*i).Id).Scan(&exists)
		if err != nil {
			return internal.Invoice{}, err
		}
		if !exists {
			return internal.Invoice{}, errors.New("customer does not exist")
		}
	}

	return *i, nil
}
