package bootstrap

import (
	"interview/db"
	"log"
)

func NewDatabase(db_connection string) db.Client {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	client, err := db.NewMySqlClient(db_connection)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseDBConnection(client db.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to DB closed.")
}
