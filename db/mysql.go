package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client interface {
	Database(string) *gorm.DB
	Connect() error
	Disconnect() error
}

type Database interface {
}

type gormClient struct {
	db *gorm.DB
}

func NewMySqlClient(connection string) (Client, error) {
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	return &gormClient{db: db}, err
}

func (gc *gormClient) Connect() error {
	db, err := gc.db.DB()
	if err != nil {
		return err
	}
	return db.Ping()
}

func (gc *gormClient) Disconnect() error {
	db, err := gc.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}

func (gc *gormClient) Database(str string) *gorm.DB {
	// db, _ := gc.db
	return gc.db
}
