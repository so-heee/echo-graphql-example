package database

import (
	"context"

	"github.com/so-heee/graphql-example/plan2/graph/pagination"
	"github.com/so-heee/graphql-example/plan2/models"
)

func (r *Repository) TodoByID(ctx context.Context, id int) (*models.Todo, error) {
	return models.FindTodo(ctx, r.db, id)
}

func (r *Repository) Todos(ctx context.Context, paginator *pagination.Paginator) ([]*models.Todo, error) {
	if paginator != nil {
		return models.Todos(paginator.Queries()...).All(ctx, r.db)
	}
	return models.Todos().All(ctx, r.db)
}

func (r *Repository) TodosCount(ctx context.Context) (int64, error) {
	return models.Todos().Count(ctx, r.db)
}
