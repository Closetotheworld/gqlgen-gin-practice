package repository

import (
	"github.com/closetotheworld/gqlgen-gin-practice/internal/models"
	database "github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/db/mysql"
	"testing"
)

func TestTodoRepository_Create(t *testing.T) {
	database.InitDB()
	models.Migrate()

	todo := models.Todo{
		Title: "test2",
		Text:  "content-2",
	}
	repo := NewTodoRepository(database.Db)
	if err := repo.Create(&todo); err != nil {
		t.Fatal(err)
	}

	t.Log(todo)
}
