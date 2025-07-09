package models

import (
	"inventory-app/config"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

func GetUserByUsername(username string) (User, error) {
	row := config.DB.QueryRow("SELECT id, username, password, role FROM users WHERE username = ?", username)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)
	return user, err
}

func CreateUser(username, password, role string) error {
	_, err := config.DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", username, password, role)
	return err
}
