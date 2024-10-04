package repositories

import (
	"context"
	"errors"

	domaincard "github.com/meetmorrowsolonmars/go-lessons/lesson-1/internal/domain/card"
)

type CardRepository struct {
	cards map[string]domaincard.Card
}

func NewCardRepository() *CardRepository {
	return &CardRepository{
		cards: make(map[string]domaincard.Card, 32),
	}
}

func (c *CardRepository) Save(ctx context.Context, card domaincard.Card) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	c.cards[card.UserID] = card

	return nil
}

func (c *CardRepository) GetByUserID(ctx context.Context, userID string) (domaincard.Card, error) {
	select {
	case <-ctx.Done():
		return domaincard.Card{}, ctx.Err()
	default:
	}

	// Правильная ли это логика?
	// Можно если корзины нет, то создать ее.
	card, exists := c.cards[userID]
	if !exists {
		return domaincard.Card{}, errors.New("card not found")
	}

	return card, nil
}
