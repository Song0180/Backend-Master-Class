package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable"
)

var testQueries *Queries
var testConnPool *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	testConnPool, err = pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("Unable to connect to the DB", err)
	}

	testQueries = New(testConnPool)

	os.Exit(m.Run())
}
