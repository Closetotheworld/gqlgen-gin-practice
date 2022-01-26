package delivery

import (
	"context"
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/usecase"
)

type Handler interface {
	CreateTodo(context.Context, model.NewTodo) (*model.Todo, error)
	UpdateTodo(context.Context, model.UpdateTodo) (*model.Todo, error)
	Todos(context.Context, *model.PaginationInput) (*model.TodoConnection, error)
	Todo(context.Context, int) (*model.Todo, error)
}

type TodoHandler struct {
	uc usecase.TodoUsecase
}

func NewTodoHandler(uc usecase.TodoUsecase) Handler {
	return &TodoHandler{uc}
}

func (t TodoHandler) CreateTodo(ctx context.Context, todo model.NewTodo) (*model.Todo, error) {
	return t.uc.Create(todo)
}

func (t TodoHandler) UpdateTodo(ctx context.Context, todo model.UpdateTodo) (*model.Todo, error) {
	dummy := model.Todo{}
	return &dummy, nil
}

func (t TodoHandler) Todos(ctx context.Context, input *model.PaginationInput) (*model.TodoConnection, error) {
	var after int
	if input.After == nil {
		after = 0
	} else {
		after = *input.After
	}
	first := *input.First
	return t.uc.GetTodos(uint(after), uint(first))
}

func (t TodoHandler) Todo(ctx context.Context, id int) (*model.Todo, error) {
	return t.uc.GetID(uint(id))
}
