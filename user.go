package main

import (
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/lib/pq"
)

type User struct {
	Email    string
	Password string
	Keys     []string
}

var users = []User{
	{Email: "user10@test.com", Password: "test", Keys: generateKeys(10)},
	{Email: "user25@test.com", Password: "test", Keys: generateKeys(25)},
	{Email: "user50@test.com", Password: "test", Keys: generateKeys(50)},
	{Email: "user100@test.com", Password: "test", Keys: generateKeys(100)},
	{Email: "user200@test.com", Password: "test", Keys: generateKeys(200)},
}

func generateKeys(n int) []string {
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		keys[i] = randomString(8)
	}
	return keys
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

		values := generateKeys(n)
		for k, v := range values {
			_, err := db.Exec("INSERT INTO user_values (user_id, key, value) VALUES ($1, $2, $3)", id, k, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
