package delivery

import (
	"context"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/models"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/repository"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/usecase"
	database "github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/db/mysql"
	"testing"
)

func TestTodoHandler_Todo(t *testing.T) {
	database.InitDB()
	models.Migrate()
	repo := repository.NewTodoRepository(database.Db)
	uc := usecase.NewTodoUsecase(repo)
	handler := NewTodoHandler(uc)

	todo, err := handler.Todo(context.Background(), 1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(todo)
}
