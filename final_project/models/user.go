package models

import (
	"errors"

	"example.com/final_project/db"
	"example.com/final_project/utils"
)

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Email     string `json:"email" binding:"required"`
	CreatedAt string `json:"created_at"`
}

func (u User) Save() error {
	query := `INSERT INTO users (name, password, email, created_at)
	VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	cryptedPw, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// Exec: Query for modify data
	result, err := stmt.Exec(u.Name, cryptedPw, u.Email, u.CreatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)
	var retrivePassword string
	err := row.Scan(&retrivePassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivePassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
