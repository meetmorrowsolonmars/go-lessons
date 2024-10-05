package item

import domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"

type GetItemByIDResponse struct {
	domainitem.Item
}

type GetItemsResponse struct {
	Items []domainitem.Item `json:"items"`
}
