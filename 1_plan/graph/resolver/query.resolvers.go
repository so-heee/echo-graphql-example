package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/so-heee/graphql-example/plan01/graph/database"
	"github.com/so-heee/graphql-example/plan01/graph/generated"
	"github.com/so-heee/graphql-example/plan01/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := database.NewUserDao(r.DB).FindAll(ctx)
	if err != nil {
		return nil, err
	}
	results := []*model.User{}
	for _, user := range users {
		results = append(results, &model.User{
			ID:   user.ID,
			Name: user.Name,
		})
	}
	return results, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	user, err := database.NewUserDao(r.DB).FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("not found")
	}
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
