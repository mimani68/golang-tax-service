package bootstrap

import (
	"interview/db"
	"log"

	"gorm.io/gorm"
)

func NewDatabase(db_connection string) *gorm.DB {
	client, err := db.NewClient(db_connection)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseDBConnection(client *gorm.DB) {
	if client == nil {
		return
	}

	// err := client.DB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Connection to DB closed.")
}
