package db

import (
	"database/sql"
	"fmt"
	"os"
	//"log"

	"github.com/cenkalti/backoff/v4"
	_ "github.com/lib/pq"

	//"github.com/golang-migrate/migrate/v4"
	//_ "github.com/golang-migrate/migrate/v4/source/file"
	//_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func InitStore() (*sql.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
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

// func RunMigrations() {
// 	m, err := migrate.New(
// 		"file://db/migrations",
// 		"cockroachdb://ymagin@db:26257/cookhubdb?sslmode=disable")
// 	if err != nil {
// 		log.Fatalf("failed to migrate: %s", err)
// 	}
// 	if err := m.Up(); err != nil {
// 		log.Fatalf("failed to up: %s", err)	
// 	}
// }

