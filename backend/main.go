package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=password dbname=testdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	fmt.Println("Connected to DB!")

	http.HandleFunc("/api/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong from backend"))
	})

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
