package models

import (
	"errors"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (u *User) Create() error {
	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	err := DB.QueryRow(
		"INSERT INTO users(username) VALUES($1) RETURNING id",
		u.Username,
	).Scan(&u.ID)

	return err
}

func GetUser(id int64) (*User, error) {
	u := &User{}
	err := DB.QueryRow(
		"SELECT id, username FROM users WHERE id = $1",
		id,
	).Scan(&u.ID, &u.Username)
	
	return u, err
}