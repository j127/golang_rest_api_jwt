package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

// User represents a user
type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// JWT represents a JWT
type JWT struct {
	Token string `json:"token"`
}

// Error represents an error
type Error struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDatabase := os.Getenv("POSTGRES_DB")
	postgresURL := "postgresql://" + postgresUser + ":" + postgresPassword + "@localhost:5432/" + postgresDatabase

	pgURL, err := pq.ParseURL(postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pgURL)

	db, err = sql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("listening on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("signup"))
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("login"))
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("protectedEndpoint")
}

// TokenVerifyMiddleWare function
func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return nil
}
