package db

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database interface {
	Table(string) Table
	Client() Client
}

type Table interface {
	First(interface{}, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Delete(interface{}, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	Count(interface{}, ...interface{}) *gorm.DB
	Updates(interface{}, ...interface{}) *gorm.DB
}

type Client interface {
	Database(string) Database
	Connect() error
	Disconnect() error
}

type gormClient struct {
	db *gorm.DB
}

type gormDatabase struct {
	db *gorm.DB
}

type gormTable struct {
	table *gorm.DB
}

func NewClient(connection string) (Client, error) {
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

func (gc *gormClient) Database(dbName string) gorm.DB {
	return &gormDatabase{db: gc.db.Session(&gorm.Session{Context: context.Background()}).Model(&gorm.Model{})}
}

// func (gd *gormDatabase) Table(tableName string) Table {
// 	return &gormTable{table: gd.db.Table(tableName)}
// }

func (gd *gormDatabase) Client() gorm.DB {
	return &gormClient{db: gd.db}
}

func (gt *gormTable) First(dest interface{}, conds ...interface{}) *gorm.DB {
	return gt.table.First(dest, conds...)
}

func (gt *gormTable) Create(value interface{}) *gorm.DB {
	return gt.table.Create(value)
}

func (gt *gormTable) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	return gt.table.Delete(value, conds...)
}

func (gt *gormTable) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	return gt.table.Find(dest, conds...)
}

// func (gt *gormTable) Count(dest interface{}, conds ...interface{}) *gorm.DB {
// 	return gt.table.Count(dest, conds...)
// }

// func (gt *gormTable) Updates(value interface{}, conds ...interface{}) *gorm.DB {
// 	return gt.table.Updates(value, conds...)
// }
