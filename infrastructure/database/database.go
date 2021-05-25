package database

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=root dbname=menu_project password=menu_project sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}
}
