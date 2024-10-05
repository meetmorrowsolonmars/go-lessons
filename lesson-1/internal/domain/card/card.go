package card

import (
	"github.com/shopspring/decimal"

	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

type Card struct {
	UserID     string          `json:"user_id"`
	Items      []Item          `json:"items"`
	TotalPrice decimal.Decimal `json:"total_price"`
}

type Item struct {
	domainitem.Item
	Quantity int64 `json:"quantity"`
}
