package main

import (
	"database/sql"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Values   int    `json:"values"`
}

const goBackendTag = "go"

var fallbackUsers = []User{
	{Email: "user10@test.com", Password: "test", Values: 10},
	{Email: "user25@test.com", Password: "test", Values: 25},
	{Email: "user50@test.com", Password: "test", Values: 50},
	{Email: "user100@test.com", Password: "test", Values: 100},
	{Email: "user200@test.com", Password: "test", Values: 200},
}

func findUser(email, password string) *User {
	if appDB == nil {
		for _, user := range fallbackUsers {
			if user.Email != email {
				continue
			}
			if password != "" && user.Password != password {
				continue
			}
			userCopy := user
			return &userCopy
		}
		return nil
	}

	query := "SELECT email, password, values_count FROM benchmark_users WHERE email = $1 AND backend_tag = $2"
	args := []any{email, goBackendTag}
	if password != "" {
		query += " AND password = $3"
		args = append(args, password)
	}

	row := appDB.QueryRow(query, args...)
	var user User
	err := row.Scan(&user.Email, &user.Password, &user.Values)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return nil
	}

	return &user
}
