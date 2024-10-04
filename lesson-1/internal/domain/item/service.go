package item

import "context"

type Repository interface {
	GetItemByID(ctx context.Context, itemID string) (Item, error)
	GetItems(ctx context.Context) ([]Item, error)
}

type Service struct {
	repo Repository
}

func NewItemService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetItemByID(ctx context.Context, itemID string) (Item, error) {
	item, err := s.repo.GetItemByID(ctx, itemID)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func (s *Service) GetItems(ctx context.Context) ([]Item, error) {
	items, err := s.repo.GetItems(ctx)
	if err != nil {
		return []Item{}, err
	}

	return items, nil
}
