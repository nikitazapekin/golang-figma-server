package models

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
}
 
func CreateUser(db *sql.DB, username, email, password string) error {
	tx, err := db.Begin()  
	if err != nil {
		return err
	}
 
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return err
	}

	 
	_, err = tx.Exec(`
		INSERT INTO users (username, email, password_hash) 
		VALUES ($1, $2, $3)`, username, email, string(passwordHash))
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
 
func AuthenticateUser(db *sql.DB, email, password string) (*User, error) {
	var user User
	err := db.QueryRow(`
		SELECT id, username, email, password_hash 
		FROM users WHERE email = $1`, email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}
 
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
