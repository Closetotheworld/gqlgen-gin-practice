package usecase

import (
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/models"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/repository"
)

type TodoUsecase interface {
	GetID(id uint) (*model.Todo, error)
	Create(todo model.NewTodo) (*model.Todo, error)
	GetTodos(id uint, offset uint) (*model.TodoConnection, error)
}

type Usecase struct {
	repo repository.Repository
}

func NewTodoUsecase(repo repository.Repository) TodoUsecase {
	return &Usecase{repo: repo}
}

func (u Usecase) GetID(id uint) (*model.Todo, error) {
	data, err := u.repo.GetByID(id)
	respModel := model.Todo{
		int(data.ID),
		data.Title,
		data.Text,
		data.Done,
	}
	return &respModel, err
}

func (u Usecase) Create(todo model.NewTodo) (*model.Todo, error) {
	modelData := models.Todo{
		Title: todo.Title,
		Text:  todo.Text,
	}
	err := u.repo.Create(&modelData)
	if err != nil {
		return &model.Todo{}, err
	}

	respModel := model.Todo{
		int(modelData.ID),
		modelData.Title,
		modelData.Text,
		modelData.Done,
	}

	return &respModel, err
}

func (u Usecase) GetTodos(id uint, offset uint) (*model.TodoConnection, error) {
	modelData, err := u.repo.GetByIDwithOffset(id, offset)
	if err != nil {
		return nil, err
	}
	if len(modelData) == 0 {
		return &model.TodoConnection{
			PageInfo: &model.PageInfo{
				StartCursor:     nil,
				EndCursor:       nil,
				HasNextPage:     false,
				HasPreviousPage: false,
			},
			Edges: []*model.TodoEdge{},
		}, nil
	}

	edges := make([]*model.TodoEdge, len(modelData))
	for i, todo := range modelData {
		e := &model.TodoEdge{
			Cursor: int(todo.ID),
			Node: &model.Todo{
				ID:    int(todo.ID),
				Title: todo.Title,
				Text:  todo.Text,
				Done:  todo.Done,
			},
		}
		edges[i] = e
	}
	startCursor := edges[0].Cursor
	endCursor := edges[len(edges)-1].Cursor
	hasPreviousPage, hasNextPage, err := u.repo.CheckPage(startCursor, endCursor)

	return &model.TodoConnection{
		PageInfo: &model.PageInfo{
			StartCursor:     &startCursor,
			EndCursor:       &endCursor,
			HasNextPage:     hasNextPage,
			HasPreviousPage: hasPreviousPage,
		},
		Edges: edges,
	}, err

}
