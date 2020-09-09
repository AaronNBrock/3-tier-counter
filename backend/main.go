package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	retry "github.com/avast/retry-go"
	_ "github.com/lib/pq"
)

type server struct {
	db *sql.DB
}

type countResponse struct {
	Count int `json:"count"`
}

func (s *server) countHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Todo: Set to proper domain

	if r.Method == http.MethodPost {
		s.db.Exec("UPDATE counters SET count=count+1 WHERE id=1;")
	}

	var count int
	row := s.db.QueryRow("SELECT count FROM counters WHERE id=1;")
	err := row.Scan(&count)
	checkError(err)

	w.Header().Set("Content-Type", "application/json")
	countResponse := countResponse{
		Count: count,
	}
	json.NewEncoder(w).Encode(countResponse)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func main() {

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		getEnv("DB_USERNAME", "postgres"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_NAME", "postgres"),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := server{db: db}

	err = retry.Do(
		func() error {
			_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS counters (
				id    int  PRIMARY KEY,
				count int
			);
			INSERT INTO counters (id, count) VALUES (1, 0) ON CONFLICT DO NOTHING;
			`)
			return err
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/api/count", s.countHandler)
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
