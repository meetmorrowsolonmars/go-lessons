package card

import (
	"context"
	"net/http"

	domaincard "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/card"
	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

type Service interface {
	Create(ctx context.Context, userID string) error
	GetByUserID(ctx context.Context, userID string) (domaincard.Card, error)
	AddItem(ctx context.Context, userID string, item domainitem.Item) (domaincard.Card, error)
	RemoveItem(ctx context.Context, userID string, itemID string) (domaincard.Card, error)
}

type ItemService interface {
	GetItemByID(ctx context.Context, itemID string) (domainitem.Item, error)
}

type Implementation struct {
	cardService Service
	itemService ItemService
}

func NewCardServerImplementation(cardService Service, itemService ItemService) *Implementation {
	return &Implementation{
		cardService: cardService,
		itemService: itemService,
	}
}

func RegisterRoutes(mux *http.ServeMux, i *Implementation) {
	mux.HandleFunc("GET /cards/{user_id}", i.GetByUserID)
	mux.HandleFunc("POST /cards/{user_id}", i.Create)
	mux.HandleFunc("POST /cards/{user_id}/items", i.AddItem)
	mux.HandleFunc("DELETE /cards/{user_id}/items/{item_id}", i.RemoveItem)
}
