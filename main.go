package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Performance Test Go Backend läuft auf Port 8081!"))
}

func waitForDB(connStr string, maxAttempts int, delay time.Duration) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for i := 0; i < maxAttempts; i++ {
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
	connStr := "host=db port=5432 user=testuser password=testpassword dbname=testdb sslmode=disable"
	db, err := waitForDB(connStr, 20, 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := insertUserValues(db); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8081", nil)
}
