package database

import (
	// external packages
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// project packages
)

var Db *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("sqllite.db"), &gorm.Config{})
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
