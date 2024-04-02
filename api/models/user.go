package models

import (
	"errors"

	"suggestions.api/db"
	"suggestions.api/utils"
)

type User struct {
	Id       int64
	Username string `binding:"required" form:"username"`
	Password string `binding:"required" form:"password"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(username, password) VALUES (?, ?)"

	stmt, err := db.Database.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Username, hashPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	u.Id = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE username = ?"

	row := db.Database.QueryRow(query, u.Username)

	var retrivedPass string
	err := row.Scan(&u.Id, &retrivedPass)

	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPass)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
