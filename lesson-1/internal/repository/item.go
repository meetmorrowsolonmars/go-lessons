package repositories

import (
	"context"
	"errors"

	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

type ItemRepository struct {
	items []domainitem.Item
}

func NewItemRepository(items []domainitem.Item) ItemRepository {
	return ItemRepository{
		items: items,
	}
}

func (r *ItemRepository) GetItemByID(ctx context.Context, itemID string) (domainitem.Item, error) {
	for _, item := range r.items {
		if item.ID == itemID {
			select {
			case <-ctx.Done():
				return domainitem.Item{}, ctx.Err()
			default:
			}

			return item, nil
		}
	}

	return domainitem.Item{}, errors.New("item not found")
}

func (r *ItemRepository) GetItems(ctx context.Context) ([]domainitem.Item, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	return r.items, nil
}
