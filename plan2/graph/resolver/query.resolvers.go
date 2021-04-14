package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/so-heee/graphql-example/plan2/graph/generated"
	"github.com/so-heee/graphql-example/plan2/graph/model"
	"github.com/so-heee/graphql-example/plan2/pagination"
)

func (r *queryResolver) Users(ctx context.Context, after *string, before *string, first *int, last *int) (*model.UserConnection, error) {

	paginator := pagination.NewPaginator(
		after,
		before,
		first,
		last,
	)

	users, err := r.repo.Users(ctx)
	if err != nil {
		return nil, err
	}

	conn := &model.UserConnection{}

	if len(users) == 0 {
		return conn, nil
	}

	limit := len(users)
	if limit > paginator.Limit() {
		limit = paginator.Limit()
	}

	conn.Edges = make([]*model.UserEdge, limit)
	conn.Nodes = make([]*model.User, limit)

	for i, u := range users[:limit] {
		cursor, _ := paginator.CreateEncodedCursor(u)

		node := &model.User{
			ID:        strconv.Itoa(u.ID),
			Name:      &u.Name.String,
			Email:     &u.Email.String,
			CreatedAt: &u.CreatedAt,
			UpdatedAt: &u.UpdatedAt,
		}

		pos := i
		if !paginator.IsAfter() {
			pos = len(conn.Edges) - i - 1
		}

		conn.Edges[pos] = &model.UserEdge{Cursor: cursor, Node: node}
		conn.Nodes[pos] = node
	}

	conn.PageInfo = &model.PageInfo{
		StartCursor: &conn.Edges[0].Cursor,
		EndCursor:   &conn.Edges[len(conn.Edges)-1].Cursor,
	}

	if len(users) > limit {
		if paginator.IsAfter() {
			conn.PageInfo.HasNextPage = true
		} else {
			conn.PageInfo.HasPreviousPage = true
		}
	}

	totalCount, err := r.repo.UsersCount(ctx)

	if err != nil {
		return conn, err
	}

	conn.TotalCount = int(totalCount)

	return conn, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
