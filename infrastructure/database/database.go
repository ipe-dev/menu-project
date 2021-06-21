package database

import (
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	var err error
	dsn := os.Getenv("DSN")
	Db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
}
