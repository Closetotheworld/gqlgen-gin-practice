package models

import database "github.com/closetotheworld/gqlgen-gin-practice/internal/pkg/db/mysql"

func Migrate() {
	database.Db.AutoMigrate(&Todo{})
}
