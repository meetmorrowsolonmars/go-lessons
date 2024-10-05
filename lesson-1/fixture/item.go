package fixture

import (
	"github.com/shopspring/decimal"

	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

var (
	Items = []domainitem.Item{
		{
			ID:    "item-1",
			Name:  "Milk",
			Price: decimal.NewFromInt32(99),
		},
		{
			ID:    "item-2",
			Name:  "iPhone",
			Price: decimal.NewFromInt32(1199),
		},
	}
)
