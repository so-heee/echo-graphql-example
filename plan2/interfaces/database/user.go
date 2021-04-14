package database

import (
	"context"

	"github.com/so-heee/graphql-example/plan2/models"
)

func (r *Repository) Users(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.db)
}
