package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/so-heee/graphql-example/plan2/graph/generated"
	"github.com/so-heee/graphql-example/plan2/graph/model"
)

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.repo.Users(ctx)
	if err != nil {
		return nil, err
	}
	results := []*model.User{}
	for _, user := range users {
		results = append(results, &model.User{
			ID:        strconv.Itoa(user.ID),
			Name:      &user.Name.String,
			Email:     &user.Email.String,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		})
	}
	return results, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
