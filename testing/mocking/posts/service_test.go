package posts

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meetmorrowsolonmars/go-lessons/testing/mocking/errors"
	"github.com/meetmorrowsolonmars/go-lessons/testing/mocking/users"
)

func TestService_Create(t *testing.T) {
	type args struct {
		title  string
		body   string
		author users.User
	}

	userID := uuid.NewString()
	createdAt := time.Date(2024, 9, 30, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		ctx     context.Context
		argv    args
		service func(t testing.TB) *Service
		want    Post
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "created successfully",
			ctx:  context.Background(),
			argv: args{
				title: "Title 1",
				body:  "Body 1",
				author: users.User{
					ID:   userID,
					Name: "General",
					Role: users.RoleUser,
				},
			},
			service: func(t testing.TB) *Service {
				store := &StoreMock{}

				store.SaveFunc = func(ctx context.Context, post Post) error {
					assert.NotEmpty(t, post.ID)
					assert.NotEmpty(t, post.Title)
					assert.NotEmpty(t, post.Body)
					assert.Equal(t, userID, post.Author.ID)

					return nil
				}

				t.Cleanup(func() {
					require.Len(t, store.SaveCalls(), 1)
				})

				return &Service{
					store: store,
					now:   func() time.Time { return createdAt },
				}
			},
			want: Post{
				Title: "Title 1",
				Body:  "Body 1",
				Author: users.User{
					ID:   userID,
					Name: "General",
					Role: users.RoleUser,
				},
				CreatedAt: createdAt,
			},
			wantErr: assert.NoError,
		},
		{
			name: "access denied error",
			ctx:  context.Background(),
			argv: args{
				title: "Title 1",
				body:  "Body 1",
				author: users.User{
					ID:   userID,
					Name: "Guest",
					Role: users.RoleGuest,
				},
			},
			service: func(t testing.TB) *Service {
				store := &StoreMock{}

				store.SaveFunc = func(ctx context.Context, post Post) error {
					return nil
				}

				t.Cleanup(func() {
					require.Len(t, store.SaveCalls(), 0)
				})

				return &Service{
					store: store,
					now:   func() time.Time { return createdAt },
				}
			},
			want: Post{},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, errors.AccessDenied, msgAndArgs...)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := tt.service(t)

			result, err := service.Create(tt.ctx, tt.argv.title, tt.argv.body, tt.argv.author)

			tt.wantErr(t, err)

			tt.want.ID = result.ID
			assert.Equal(t, tt.want, result)
		})
	}
}
