package card

import (
	"context"
	"slices"

	"github.com/shopspring/decimal"

	domainitem "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/item"
)

type Repository interface {
	Save(ctx context.Context, card Card) error
	GetByUserID(ctx context.Context, userID string) (Card, error)
}

type Service struct {
	repo Repository
}

func NewCardService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(ctx context.Context, userID string) error {
	err := s.repo.Save(ctx, Card{
		UserID:     userID,
		Items:      []Item{},
		TotalPrice: decimal.Zero,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetByUserID(ctx context.Context, userID string) (Card, error) {
	card, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

func (s *Service) AddItem(ctx context.Context, userID string, item domainitem.Item) (Card, error) {
	card, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return Card{}, err
	}

	found := false

	// Пробуем увеличить количество, если товар уже в корзине.
	for i, itm := range card.Items {
		if itm.ID == item.ID {
			found = true
			// Прояснить почему так.
			card.Items[i].Quantity++

			// Гонка данных на примере AddItem, RemoveItem и slice. Даже если len у slice не меняется, то память под
			// элементы slice общая.
			// Тут мы хотим добавить товар, а в RemoveItem в это же время удаляем одну единицу.
		}
	}

	// Если товара нет в корзине, то добавляем новый товар в корзину.
	if !found {
		card.Items = append(card.Items, Item{
			Item:     item,
			Quantity: 1,
		})
	}

	card.TotalPrice = card.TotalPrice.Add(item.Price)

	err = s.repo.Save(ctx, card)
	if err != nil {
		return Card{}, err
	}

	return card, nil
}

func (s *Service) RemoveItem(ctx context.Context, userID string, itemID string) (Card, error) {
	card, err := s.repo.GetByUserID(ctx, userID)
	if err != nil {
		return Card{}, err
	}

	price := decimal.Zero

	// Ищем товар в корзине.
	for i, itm := range card.Items {
		if itm.ID == itemID {
			if itm.Quantity <= 1 {
				// Если товар найден и он в единственном экземпляре, то удаляем его.
				card.Items = slices.Delete(card.Items, i, i+1)

				// Можно сделать вот так.
				// card.Items = append(card.Items[:i], card.Items[i+1:]...)

				// И вот так.
				// card.Items[i] = card.Items[len(card.Items)-1]
				// card.Items = card.Items[:len(card.Items)-1]
			} else {
				// Если товар найден и его больше чем 1, то удаляем одну штуку.
				card.Items[i].Quantity--
			}

			price = itm.Price

			break
		}
	}

	card.TotalPrice = card.TotalPrice.Sub(price)

	err = s.repo.Save(ctx, card)
	if err != nil {
		return Card{}, err
	}

	return card, nil

}
