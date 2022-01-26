package stub

import (
	"context"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
)

type Stub struct {
	QueryResolver struct {
		Todos func(ctx context.Context, input *model.PaginationInput) (*model.TodoConnection, error)
		Todo  func(ctx context.Context, id int) (*model.Todo, error)
	}
	MutationResolver struct {
		CreateTodo func(ctx context.Context, input model.NewTodo) (*model.Todo, error)
		UpdateTodo func(ctx context.Context, input model.UpdateTodo) (*model.Todo, error)
	}
}
