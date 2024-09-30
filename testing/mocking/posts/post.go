package posts

import (
	"time"

	"github.com/meetmorrowsolonmars/go-lessons/testing/mocking/users"
)

type Post struct {
	ID        string
	Title     string
	Body      string
	Author    users.User
	CreatedAt time.Time
}
