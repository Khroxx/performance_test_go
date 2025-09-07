package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	_ "github.com/lib/pq"
)

type User struct {
	Email    string
	Password string
	Values   map[string]string
}

var users = []User{
	{Email: "user10@test.com", Password: "test", Values: generateValues(10)},
	{Email: "user25@test.com", Password: "test", Values: generateValues(25)},
	{Email: "user50@test.com", Password: "test", Values: generateValues(50)},
	{Email: "user100@test.com", Password: "test", Values: generateValues(100)},
	{Email: "user200@test.com", Password: "test", Values: generateValues(200)},
}

func generateValues(n int) map[string]string {
	result := make(map[string]string)
	for i := 1; i <= n; i++ {
		key := "key" + strconv.Itoa(i)
		value := randomString(8)
		result[key] = value
	}
	return result
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func findUser(email, password string) *User {
	for _, u := range users {
		if u.Email == email && u.Password == password {
			return &u
		}
	}
	return nil
}

func insertUserValues(db *sql.DB) error {
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var email string
		if err := rows.Scan(&id, &email); err != nil {
			return err
		}

		// Extract number from email, e.g. user10@test.com -> 10
		var n int
		_, err := fmt.Sscanf(email, "user%d@test.com", &n)
		if err != nil {
			continue // skip if not matching pattern
		}

		values := generateValues(n)
		for k, v := range values {
			_, err := db.Exec("INSERT INTO user_values (user_id, key, value) VALUES ($1, $2, $3)", id, k, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
