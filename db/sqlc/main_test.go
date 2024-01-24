package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/song0180/simple-bank/util"
)

var testQueries *Queries
var testConnPool *pgxpool.Pool

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("Unable to read config for test", err)
	}

	testConnPool, err = pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Unable to connect to the DB:", err)
	}

	testQueries = New(testConnPool)

	os.Exit(m.Run())
}
