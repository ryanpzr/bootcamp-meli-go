package repository

import (
	"app/utils"
	"database/sql"
	"encoding/json"
	"log"

	"app/internal"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

// ProductsMySQL is the MySQL repository implementation for product entity.
type ProductsMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all products from the database.
func (r *ProductsMySQL) FindAll() (p []internal.Product, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var pr internal.Product
		// scan the row into the product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		// append the product to the slice
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the product into the database.
func (r *ProductsMySQL) Save(p *internal.Product) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
		(*p).Description, (*p).Price,
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
	(*p).Id = int(id)

	return
}

func (r *ProductsMySQL) Populate() (p []internal.Product, err error) {
	file, err := utils.OpenFile("products.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var product []internal.Product
	err = json.NewDecoder(file).Decode(&product)
	if err != nil {
		log.Fatal(err)
	}

	var listProductResponse []internal.Product
	for _, p := range product {
		res, err := r.db.Exec(
			"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
			(p).Description, (p).Price,
		)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		p.Id = int(id)
		listProductResponse = append(listProductResponse, p)
	}

	return listProductResponse, nil
}

func (r *ProductsMySQL) GetRankingProd() (p []internal.ProductRanking, err error) {
	query := `SELECT p.description, ROUND(SUM(p.price * s.quantity), 2) as Total
	FROM products p JOIN sales s ON p.id = s.product_id
	GROUP BY p.description
	ORDER BY Total DESC
	LIMIT 5`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ranking []internal.ProductRanking
	for rows.Next() {
		var p internal.ProductRanking
		err = rows.Scan(&p.Description, &p.Total)
		if err != nil {
			return nil, err
		}

		ranking = append(ranking, p)
	}

	return ranking, nil
}
