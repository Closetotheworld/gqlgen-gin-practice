package models

type Todo struct {
	ID        uint `gorm:"primaryKey,autoIncrement" json:"id"`
	CreatedAt uint `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt uint `gorm:"autoUpdateTime" json:"updated_at"`
	Title     string
	Text      string
	Done      bool `gorm:"default:false"`
}
