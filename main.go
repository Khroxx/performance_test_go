package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Performance Test Go Backend läuft!"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}
