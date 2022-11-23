package main

import (
	"log"

	"github.com/conghaile/simple-API/api"
	"github.com/conghaile/simple-API/db"
)

func main() {
	store, err := db.NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.InitWarosu(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
