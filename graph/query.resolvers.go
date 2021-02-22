package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/so-heee/echo-graphql-example/graph/generated"
	"github.com/so-heee/echo-graphql-example/graph/model"
)

func (r *queryResolver) Categoryes(ctx context.Context) ([]*model.CategorySmall, error) {
	categoryes := []*model.CategorySmall{}

	r.DB.Table("category_small").Find(&categoryes)

	return categoryes, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
