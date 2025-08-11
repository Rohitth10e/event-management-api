package models

import (
	"errors"
	"event-management-api/db"
	"event-management-api/utils"
)

type Users struct {
	ID       int64  `json:"id"`
	EMAIL    string `json:"email" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

var users []Users

func (user *Users) SAVE() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?);"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if user.PASSWORD == "" {
		return errors.New("User Password cannot be empty")
	}
	hashedPassword, err := utils.HashPassword(user.PASSWORD)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.EMAIL, hashedPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err == nil {
		user.ID = userId
	}
	return err
}

func GetUsers() ([]Users, error) {
	query := "SELECT id,email FROM users"
	res, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	var users []Users
	for res.Next() {
		var user Users
		err := res.Scan(&user.ID, &user.EMAIL)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := res.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Users) ValidateUserLogin() error {
	query := "SELECT id,email, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.EMAIL)
	var retrievedUser, hashedPassword string
	err := row.Scan(&u.ID, &retrievedUser, &hashedPassword)
	if err != nil {
		return err
	}

	match := utils.CheckHashedPassword(u.PASSWORD, hashedPassword)
	if !match {
		return errors.New("User login failed, Invalid credentials")
	}
	return nil
}
