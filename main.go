package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/song0180/simple-bank/api"
	db "github.com/song0180/simple-bank/db/sqlc"
	"github.com/song0180/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Unable to read configs:", err)
	}

	testConnPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Unable to connect to the DB:", err)
	}

	store := db.NewStore(testConnPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Unable to start the server:", err)
	}
}
