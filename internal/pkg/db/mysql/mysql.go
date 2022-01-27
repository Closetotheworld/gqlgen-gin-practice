package database

import (
	// external packages
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"strings"
	// project packages
)

var Db *gorm.DB

func InitDB() {
	// environment에 ROOT_PATH 추가
	sqlitePath := strings.Join([]string{os.Getenv("ROOT_PATH"), "sqllite.db"}, "/")
	//db, err := gorm.Open(sqlite.Open("sqllite.db"), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	if err = sqlDb.Ping(); err != nil {
		log.Fatal(err)
	}

	Db = db
}
