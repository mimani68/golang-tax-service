package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	Database(dbName string) *gorm.DB
}

type Table interface {
	First(interface{}, ...interface{}) Table
	Create(interface{}) Table
	Delete(interface{}, ...interface{}) Table
	Find(interface{}, ...interface{}) Table
	Updates(interface{}) Table
}

type Client interface {
	Database(string) Database
	Connect() error
	Disconnect() error
}

type postgresClient struct {
	db *gorm.DB
}

type postgresDatabase struct {
	db *gorm.DB
}

type postgresTable struct {
	table *gorm.DB
}

func NewClient(connection string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func (pc *postgresClient) Connect() error {
	sqlDB, err := pc.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func (pc *postgresClient) Disconnect() error {
	sqlDB, err := pc.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (pc *postgresClient) Database(dbName string) *gorm.DB {
	db := pc.db.Session(&gorm.Session{})
	return db
}

func (pd *postgresDatabase) Table(tableName string) Table {
	table := pd.db.Table(tableName)
	return &postgresTable{table: table}
}

func (pt *postgresTable) First(dest interface{}, conds ...interface{}) Table {
	return &postgresTable{table: pt.table.First(dest, conds...)}
}

func (pt *postgresTable) Create(value interface{}) Table {
	return &postgresTable{table: pt.table.Create(value)}
}

func (pt *postgresTable) Delete(value interface{}, conds ...interface{}) Table {
	return &postgresTable{table: pt.table.Delete(value, conds...)}
}

func (pt *postgresTable) Find(dest interface{}, conds ...interface{}) Table {
	return &postgresTable{table: pt.table.Find(dest, conds...)}
}

func (pt *postgresTable) Updates(values interface{}) Table {
	return &postgresTable{table: pt.table.Updates(values)}
}
