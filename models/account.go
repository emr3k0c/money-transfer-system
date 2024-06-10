package models

type Account struct {
	ID         int     `json:"id"`
	CustomerID int     `json:"customer_id"`
	Balance    float64 `json:"balance"`
}
