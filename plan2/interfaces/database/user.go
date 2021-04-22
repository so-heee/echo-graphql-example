package database

import (
	"context"

	"github.com/so-heee/graphql-example/plan2/domain/model"
	"github.com/so-heee/graphql-example/plan2/infrastructure/database/models"
)

func (r *Repository) Users(ctx context.Context, paginator *model.Paginator) ([]*models.User, error) {
	if paginator != nil {
		return models.Users(paginator.Queries()...).All(ctx, r.db)
	}
	return models.Users().All(ctx, r.db)
}

func (r *Repository) UsersCount(ctx context.Context) (int64, error) {
	return models.Users().Count(ctx, r.db)
}
