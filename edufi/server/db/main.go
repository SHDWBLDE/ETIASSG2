package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Postgres struct{}

var Db *sql.DB

func StartDB() {

	dbUser := os.Getenv("EDUFI_DB_USER")
	dbPwd := os.Getenv("EDUFI_DB_PWD")
	dbName := os.Getenv("EDUFI_DB_NAME")
	dbPort := os.Getenv("EDUFI_DB_PORT")
	dbHost := os.Getenv("EDUFI_DB_HOST")

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbHost, dbUser, dbPwd, dbPort, dbName)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Db = db

	// Confirm a successful connection.
	if err := Db.Ping(); err != nil {
		log.Fatalf("failed to establish database connection: %v", err)
	}

}
