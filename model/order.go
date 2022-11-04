package model

type Order struct {
	OrderID    string
	VAT        float64
	TotalPrice float64
	Products   []Product
}
