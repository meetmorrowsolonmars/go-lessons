package posts

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/meetmorrowsolonmars/go-lessons/testing/mocking/errors"
	"github.com/meetmorrowsolonmars/go-lessons/testing/mocking/users"
)

//go:generate go run -mod=mod github.com/matryer/moq@v0.5.0 -rm -out store_mock.go . Store
type Store interface {
	Save(ctx context.Context, post Post) error
	DeleteByID(ctx context.Context, ID string) error
	GetByID(ctx context.Context, ID string) (Post, error)
	GetList(ctx context.Context, limit int, offset int) ([]Post, error)
}

type Service struct {
	store Store
	now   func() time.Time
}

func NewService(store Store) *Service {
	return &Service{
		store: store,
		now:   time.Now,
	}
}

func (s *Service) Create(ctx context.Context, title string, body string, author users.User) (Post, error) {
	if !author.Role.CanWrite() {
		return Post{}, fmt.Errorf("user does not have permission to create post: %w", errors.AccessDenied)
	}

	post := Post{
		ID:        uuid.NewString(),
		Title:     title,
		Body:      body,
		Author:    author,
		CreatedAt: s.now(),
	}

	if err := s.store.Save(ctx, post); err != nil {
		return Post{}, fmt.Errorf("failed to save post: %w", err)
	}

	return post, nil
}

func (s *Service) GetAll(ctx context.Context, user users.User, limit int, offset int) ([]Post, error) {
	if !user.Role.CanRead() {
		return nil, fmt.Errorf("user does not have permission to get posts: %w", errors.AccessDenied)
	}

	posts, err := s.store.GetList(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}

	return posts, nil
}
