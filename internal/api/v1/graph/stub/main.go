package stub

import (
	"context"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
)

type Stub struct {
	QueryResolver struct {
		Todos func(ctx context.Context) ([]*model.Todo, error)
	}
	MutationResolver struct {
		CreateTodo func(ctx context.Context, input model.NewTodo) (*model.Todo, error)
	}
}
