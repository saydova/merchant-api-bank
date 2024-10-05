package models

type Customer struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Balance  float64 `json:"balance"`
}

type Merchant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type History struct {
	CustomerID int     `json:"customer_id"`
	Action     string  `json:"action"`
	Amount     float64 `json:"amount,omitempty"`
}
