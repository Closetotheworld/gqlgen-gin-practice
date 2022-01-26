package repository

import "github.com/closetotheworld/gqlgen-gin-practice/internal/models"

type Repository interface {
	Create(data *models.Todo) error
	GetByID(id uint) (models.Todo, error)
	GetByIDwithOffset(id uint, offset uint) ([]*models.Todo, error)
	CheckPage(startCursor int, endCursor int) (bool, bool, error)
}
