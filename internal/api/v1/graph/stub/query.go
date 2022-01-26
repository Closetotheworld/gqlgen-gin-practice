package stub

import (
	"context"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/generated"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
)

type stubQuery struct {
	*Stub
}

func (s *Stub) Query() generated.QueryResolver {
	return &stubQuery{Stub: s}
}

func (s stubQuery) Todos(ctx context.Context, input *model.PaginationInput) (*model.TodoConnection, error) {
	return s.QueryResolver.Todos(ctx, input)
}

func (s stubQuery) Todo(ctx context.Context, id int) (*model.Todo, error) {
	return s.QueryResolver.Todo(ctx, id)
}
