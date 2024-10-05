package item

import (
	"context"
	"net/http"

	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

type Service interface {
	GetItemByID(ctx context.Context, itemID string) (domainitem.Item, error)
	GetItems(ctx context.Context) ([]domainitem.Item, error)
}

type Implementation struct {
	service Service
}

func NewItemServerImplementation(service Service) *Implementation {
	return &Implementation{
		service: service,
	}
}

func RegisterRoutes(mux *http.ServeMux, i *Implementation) {
	mux.HandleFunc("GET /items", i.GetItems)
	mux.HandleFunc("GET /items/{item_id}", i.GetItemByID)
}
