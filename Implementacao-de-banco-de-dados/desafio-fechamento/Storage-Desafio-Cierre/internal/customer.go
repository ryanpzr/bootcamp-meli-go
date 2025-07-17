package internal

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string `json:"first_name"`
	// LastName is the last name of the customer.
	LastName string `json:"last_name"`
	// Condition is the condition of the customer.
	Condition int `json:"condition"`
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int `json:"id"`
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

type ConditionCustomer struct {
	Name         string  `json:"name"`
	InvoiceTotal float64 `json:"invoice_total"`
}

type CustomerRanking struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Total     float64 `json:"total"`
}
