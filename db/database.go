package db

import (
	"database/sql"
	"fmt"
	"os"
	"log"

	"github.com/cenkalti/backoff/v4"
	_ "github.com/cockroachdb/cockroach-go/v2/crdb"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
)

func InitStore() (*sql.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	var (
		db  *sql.DB
		err error
	)
	
	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		return err
	}

	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigrations() {
	m, err := migrate.New(
		"file://db/migrations",
		"cockroachdb://ymagin@db:26257/cookhubdb?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to migrate: %s", err)
	}
	if err := m.Up(); err != nil {
		log.Fatalf("failed to up: %s", err)	
	}
}

