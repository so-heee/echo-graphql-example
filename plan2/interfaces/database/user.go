package database

import (
	"context"

	"github.com/so-heee/graphql-example/plan2/models"
)

func (r *Repository) Users(ctx context.Context) ([]*models.User, error) {
	return models.Users().All(ctx, r.db)
}

func (r *Repository) UsersCount(ctx context.Context) (int64, error) {
	return models.Users().Count(ctx, r.db)
}
