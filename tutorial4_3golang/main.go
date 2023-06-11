package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, HEAD, OPTIONS")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, XMLHttpRequest")
	w.WriteHeader(http.StatusOK)
	encode := map[string]string{
		"message": "Hello World",
	}
	json.NewEncoder(w).Encode(encode)
}

func main() {
	http.HandleFunc("/", home)
	log.Println("Server connected at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error in starting the server")
	}
}
