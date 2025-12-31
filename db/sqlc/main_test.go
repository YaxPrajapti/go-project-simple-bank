package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbConnString = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
)

var testQueris *Queries

func TestMain(m *testing.M) {
	ctx := context.Background()
	conn, err := pgxpool.New(ctx, dbConnString)

	if err != nil {
		log.Fatal("Connection to db failed due to error: ", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("Ping to db failed due to error: ", err)
	}

	testQueris = New(conn)

	code := m.Run()
	conn.Close()

	os.Exit(code)
}
