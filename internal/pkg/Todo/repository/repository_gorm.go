package repository

import (
	"github.com/closetotheworld/gqlgen-gin-practice/api/v1/graph/model"
	"github.com/closetotheworld/gqlgen-gin-practice/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) Repository {
	return &TodoRepository{db}
}

func (t TodoRepository) Create(data *models.Todo) error {
	return t.db.Create(&data).Error
}

func (t TodoRepository) GetByID(id uint) (models.Todo, error) {
	todoModel := models.Todo{}
	err := t.db.First(&todoModel, id).Error
	return todoModel, err
}

func (t TodoRepository) GetByIDwithOffset(id uint, offset uint) ([]*models.Todo, error) {
	var todoModels []*models.Todo
	db := t.db
	if id > 0 {
		db = db.Where("id > ?", id)
	}
	if err := db.Limit(int(offset)).Find(&todoModels).Error; err != nil {
		return nil, err
	}

	return todoModels, nil
}

func (t TodoRepository) CheckPage(startCursor int, endCursor int) (bool, bool, error) {
	var dummy model.Todo
	var dummy2 model.Todo
	hasPreviousPage := false
	hasNextPage := false

	if err := t.db.Where("id < ?", startCursor).First(&dummy).Error; err == nil {
		hasPreviousPage = true
	}
	if err := t.db.Where("id > ?", endCursor).First(&dummy2).Error; err == nil {
		hasNextPage = true
	}
	return hasPreviousPage, hasNextPage, nil
}
