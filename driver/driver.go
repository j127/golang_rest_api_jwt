package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

var db *sql.DB

// ConnectDB exports the database connection
func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	// TODO: sslmode is disabled here just because this is a simple learning
	// experiment, but be sure to read up about it for production projects.
	// - https://errorsingo.com/github.com-lib-pq-err-ssl-not-supported/
	// - https://jdbc.postgresql.org/documentation/head/connect.html
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DB")
	postgresURL := "postgresql://" + postgresUser + ":" + postgresPassword + "@localhost:5432/" + postgresDatabase + "?sslmode=disable"

	pgURL, err := pq.ParseURL(postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
