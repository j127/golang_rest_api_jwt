package userRepository

import (
	"database/sql"
	"log"

	"github.com/j127/golang_rest_api_jwt/models"
)

type UserRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Signup creates a user in the database
func (u UserRepository) Signup(db *sql.DB, user models.User) models.User {

	stmt := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id;"
	err := db.QueryRow(stmt, user.Email, user.Password).Scan(&user.ID)
	logFatal(err)

	user.Password = ""
	return user
}

// Login sees if a user is in the database
func (u UserRepository) Login(db *sql.DB, user models.User) (models.User, error) {
	row := db.QueryRow("SELECT * FROM users WHERE email = $1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}
