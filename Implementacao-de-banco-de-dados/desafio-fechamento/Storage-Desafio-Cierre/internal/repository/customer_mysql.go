package repository

import (
	"app/internal"
	"app/utils"
	"database/sql"
	"encoding/json"
	"log"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []internal.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs internal.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *internal.Customer) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
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
	(*c).Id = int(id)

	return
}

func (r *CustomersMySQL) Populate() (c []internal.Customer, err error) {
	file, err := utils.OpenFile("customers.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var costumer []internal.Customer
	err = json.NewDecoder(file).Decode(&costumer)
	if err != nil {
		log.Fatal(err)
	}

	var listCostumerResponse []internal.Customer
	for _, cost := range costumer {
		res, err := r.db.Exec(
			"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
			(cost).FirstName, (cost).LastName, (cost).Condition,
		)
		if err != nil {
			return nil, err
		}

		id, err := res.LastInsertId()
		if err != nil {
			return nil, err
		}

		cost.Id = int(id)
		listCostumerResponse = append(listCostumerResponse, cost)
	}

	return listCostumerResponse, nil
}

func (r *CustomersMySQL) GetCondition() ([]internal.ConditionCustomer, error) {
	query := `SELECT IF(c.condition = 1, 'activo', 'inactivo') AS Cond, ` +
		`ROUND(SUM(i.total), 2) AS Total FROM customers c JOIN invoices i on c.id = i.customer_id GROUP BY c.condition`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var conditionList []internal.ConditionCustomer
	for rows.Next() {
		var condition internal.ConditionCustomer
		err = rows.Scan(&condition.Name, &condition.InvoiceTotal)
		if err != nil {
			return nil, err
		}

		conditionList = append(conditionList, condition)
	}

	return conditionList, nil
}

func (r *CustomersMySQL) GetRankingCus() ([]internal.CustomerRanking, error) {
	query := `SELECT c.first_name, c.last_name, SUM(i.total) AS amount ` +
		`FROM customers c JOIN invoices i on c.id = i.customer_id GROUP BY c.first_name, c.last_name
		ORDER BY amount DESC LIMIT 5;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rankingList []internal.CustomerRanking
	for rows.Next() {
		var ranking internal.CustomerRanking
		err = rows.Scan(&ranking.FirstName, &ranking.LastName, &ranking.Total)
		if err != nil {
			return nil, err
		}

		rankingList = append(rankingList, ranking)
	}

	return rankingList, nil
}
