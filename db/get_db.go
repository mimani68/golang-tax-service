package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
