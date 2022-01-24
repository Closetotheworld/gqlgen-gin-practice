package stub

import (
	"context"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/generated"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
)

type stubMutation struct {
	*Stub
}

func (s *Stub) Mutation() generated.MutationResolver {
	return &stubMutation{Stub: s}
}

func (s *stubMutation) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic("implement me")
}
