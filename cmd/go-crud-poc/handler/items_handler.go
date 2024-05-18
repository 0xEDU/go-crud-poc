package handler

import (
	"log"
	"net/http"
)

func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Println("GET /items")
		getItems(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /items"))
}
