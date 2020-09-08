package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/count", countHandler)

	// fs := http.FileServer(http.Dir("./frontend/dist"))
	// http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

var count = 0

type countResponse struct {
	Count int `json:"count"`
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Todo: Set to proper domain

	if r.Method == http.MethodPost {
		count++
	}

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
