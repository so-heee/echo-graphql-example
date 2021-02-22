package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/so-heee/echo-graphql-example/graph/generated"
	"github.com/so-heee/echo-graphql-example/graph/model"
)

func (r *mutationResolver) CreateCategoryLarge(ctx context.Context, input model.NewCategoryLarge) (*model.CategoryLarge, error) {
	categoryLarge := model.CategoryLarge{
		CategoryName: input.CategoryName,
	}
	r.DB.Table("category_large").Create(&categoryLarge)

	return &categoryLarge, nil
}

func (r *queryResolver) Categoryes(ctx context.Context) ([]*model.CategoryLarge, error) {
	categoryes := []*model.CategoryLarge{}

	r.DB.Table("category_large").Find(&categoryes)

	return categoryes, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
