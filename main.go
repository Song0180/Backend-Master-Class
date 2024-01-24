package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/song0180/simple-bank/api"
	db "github.com/song0180/simple-bank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable"
	serverAddress = "127.0.0.1:8080"
)

func main() {
	var err error
	testConnPool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Unable to connect to the DB", err)
	}

	store := db.NewStore(testConnPool)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Unable to start the server:", err)
	}
}
