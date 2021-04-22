package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/so-heee/graphql-example/plan2/domain/model"
	"github.com/so-heee/graphql-example/plan2/graph/generated"
)

func (r *queryResolver) Todo(ctx context.Context, id *int) (*model.Todo, error) {
	todo, err := r.repo.TodoByID(ctx, id)
	if err != nil {
		return nil, err
	}
	node := &model.Todo{
		ID:        strconv.Itoa(todo.ID),
		Text:      &todo.Text.String,
		CreatedAt: &todo.CreatedAt,
		UpdatedAt: &todo.UpdatedAt,
	}
	return node, nil
}

func (r *queryResolver) Todos(ctx context.Context, after *string, before *string, first *int, last *int, orderBy []*model.TodoOrder) (*model.TodoConnection, error) {
	paginator := model.NewPaginator(
		after,
		before,
		first,
		last,
		model.TodoOrderToPaginationOrders(orderBy),
	)

	todos, err := r.repo.Todos(ctx, paginator)
	if err != nil {
		return nil, err
	}

	conn := &model.TodoConnection{}

	if len(todos) == 0 {
		return conn, nil
	}

	limit := len(todos)
	if limit > paginator.Limit() {
		limit = paginator.Limit()
	}

	conn.Edges = make([]*model.TodoEdge, limit)
	conn.Nodes = make([]*model.Todo, limit)

	for i, t := range todos[:limit] {
		cursor, _ := paginator.CreateEncodedCursor(t)

		node := &model.Todo{
			ID:        strconv.Itoa(t.ID),
			Text:      &t.Text.String,
			CreatedAt: &t.CreatedAt,
			UpdatedAt: &t.UpdatedAt,
		}

		pos := i
		if !paginator.IsAfter() {
			pos = len(conn.Edges) - i - 1
		}

		conn.Edges[pos] = &model.TodoEdge{Cursor: cursor, Node: node}
		conn.Nodes[pos] = node
	}

	conn.PageInfo = &model.PageInfo{
		StartCursor: &conn.Edges[0].Cursor,
		EndCursor:   &conn.Edges[len(conn.Edges)-1].Cursor,
	}

	if len(todos) > limit {
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

func (r *queryResolver) Users(ctx context.Context, after *string, before *string, first *int, last *int, orderBy []*model.UserOrder) (*model.UserConnection, error) {
	paginator := model.NewPaginator(
		after,
		before,
		first,
		last,
		model.UserOrderToPaginationOrders(orderBy),
	)

	users, err := r.repo.Users(ctx, paginator)
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
