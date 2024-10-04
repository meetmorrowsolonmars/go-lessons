package card

import (
	"github.com/shopspring/decimal"

	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

type Card struct {
	UserID     string
	Items      []Item
	TotalPrice decimal.Decimal
}

type Item struct {
	domainitem.Item
	Quantity int64
}
