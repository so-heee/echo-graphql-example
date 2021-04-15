package database

import (
	"context"

	"github.com/so-heee/graphql-example/plan2/models"
	"github.com/so-heee/graphql-example/plan2/pagination"
)

func (r *Repository) Todos(ctx context.Context, paginator *pagination.Paginator) ([]*models.Todo, error) {
	if paginator != nil {
		return models.Todos(paginator.Queries()...).All(ctx, r.db)
	}
	return models.Todos().All(ctx, r.db)
}

func (r *Repository) TodosCount(ctx context.Context) (int64, error) {
	return models.Todos().Count(ctx, r.db)
}
