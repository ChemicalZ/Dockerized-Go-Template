package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewSQLStorage() *sqlx.DB {
	dbURL := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sqlx.Connect("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	return db
}

func NewTestSQLStorage() *sqlx.DB {
	db, err := sqlx.Connect("postgres", os.Getenv("TEST_DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	return db
}
