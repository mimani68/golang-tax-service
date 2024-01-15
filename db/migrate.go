package db

import (
	"interview/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MigrateDatabase(dsn string) {
	db := getDatabase(dsn)

	// AutoMigrate will create or update the tables based on the models
	err := db.AutoMigrate(&domain.CartEntity{}, &domain.CartItemEntity{})
	if err != nil {
		panic(err)
	}
}

func getDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
