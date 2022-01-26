package Todo

import (
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/delivery"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/repository"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/Todo/usecase"
	"gorm.io/gorm"
)

func InitTodoService(db *gorm.DB) delivery.Handler {
	repo := repository.NewTodoRepository(db)
	uc := usecase.NewTodoUsecase(repo)
	handler := delivery.NewTodoHandler(uc)
	return handler
}
