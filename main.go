package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// "github.com/gorilla/mux"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var appDB *sql.DB

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Performance Test Go Backend läuft auf Port 8081!"))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, GoToken, SpringToken, DjangoToken")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func waitForDB(connStr string, maxAttempts int, delay time.Duration) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for range maxAttempts {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		time.Sleep(delay)
	}
	return nil, err
}

func main() {
	connStr := buildConnectionString()
	db, err := waitForDB(connStr, 20, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	appDB = db
	defer appDB.Close()

	// Mit CORS
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/login", login).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/userinfo", userInfo).Methods("GET")
	http.ListenAndServe(":8081", corsMiddleware(router))

	// Ohne  CORS
	// http.HandleFunc("/", home)
	// http.HandleFunc("/login", login)
	// http.ListenAndServe(":8081", nil)
}

func buildConnectionString() string {
	host := getenvDefault("DB_HOST", "localhost")
	port := getenvDefault("DB_PORT", "5432")
	user := getenvDefault("DB_USER", "testuser")
	password := getenvDefault("DB_PASSWORD", "testpassword")
	database := getenvDefault("DB_NAME", "testdb")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
}

func getenvDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
