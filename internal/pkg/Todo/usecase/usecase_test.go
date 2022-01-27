package usecase

import (
	"github.com/closetotheworld/gqlgen-gin-practice/internal/models"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/repository"
	database "github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/db/mysql"
	"testing"
)

func TestUsecase_GetID(t *testing.T) {
	database.InitDB()
	models.Migrate()
	repo := repository.NewTodoRepository(database.Db)
	uc := NewTodoUsecase(repo)

	todo, err := uc.GetID(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(todo)
}
