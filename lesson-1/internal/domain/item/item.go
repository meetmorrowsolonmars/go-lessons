package item

import "github.com/shopspring/decimal"

type Item struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price"`
}
