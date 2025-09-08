package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your_secret_key")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user := findUser(creds.Email, creds.Password)
	if user == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func userInfo(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("GoToken")
	if tokenStr == "" {
		http.Error(w, "Missing GoToken header", http.StatusUnauthorized)
		return
	}
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	user := findUser(claims.Email, "")
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
