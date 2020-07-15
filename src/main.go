package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

func main() {
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
