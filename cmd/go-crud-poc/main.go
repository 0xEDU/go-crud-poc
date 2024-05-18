package main

import (
	"log"
	"net/http"

	"github.com/0xEDU/go-crud-poc/cmd/go-crud-poc/handler"
)

func main() {
	log.Println("Starting server on port 8080")
	http.HandleFunc("/items", handler.ItemsHandler)
	http.ListenAndServe(":8080", nil)
}
