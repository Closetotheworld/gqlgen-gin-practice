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

func (s stubQuery) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic("implement me")
}
