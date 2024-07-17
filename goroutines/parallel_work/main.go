package main

import (
	"context"
	"database/sql"
	"fmt"
)

func main() {
}

type Person struct {
	ID   int
	Name string
}

type Repository struct {
	shards map[int]*sql.DB
}

func (r *Repository) FindByName(ctx *context.Context, name string) ([]*Person, error) {
	// Our database has many shards. We use Person.ID as the sharding key.
	// Now we want to find all records containing this name.
	// Implement this function.

	query := fmt.Sprintf("SELECT id, name FROM person WHERE name = %s", name)
	_ = query

	return make([]*Person, 0), nil
}
