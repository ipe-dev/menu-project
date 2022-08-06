package database

import (
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	var err error
	dsn := os.Getenv("MySqlDsn")
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
